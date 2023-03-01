package usecase

import (
	gormdb "PTH-IT/api_golang/database/gormdb"
	"PTH-IT/api_golang/domain/model"
	errormessage "PTH-IT/api_golang/log/error"
	"PTH-IT/api_golang/utils"
	"encoding/json"
	"fmt"
	"net/http"

	InforLog "PTH-IT/api_golang/log/infor"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func NewInteractor(
	// gormDb *gorm.DB,
	referrance Reference,

) Interactor {

	return Interactor{
		// gormDb,
		referrance,
	}
}

type Interactor struct {
	// gormDb     *gorm.DB
	referrance Reference
}

// LoginUser godoc
// @Summary LoginUser
// @Description login username
// @Tags gormDB
// @Accept json
// @Produce json
// @Param user body  model.Login true "model.Login"
// @Success 201 {object} model.Token
// @Failure 400 {object} string
// @Router /gormdb/login [post]
func (i *Interactor) LoginUserGormdb(context echo.Context) error {

	var user model.Login
	err := context.Bind(&user)
	if err != nil {
		return context.String(http.StatusBadRequest, "no user")
	}

	result, err := i.referrance.GetUserGormdb(user.UserID, *utils.CryptPassword(user.Password))
	if err != nil {
		return err
	}
	if result == nil {
		return context.String(http.StatusBadRequest, "user no exist")
	}

	tokenString := utils.GenerateToken(result.UserID)
	token := &model.Token{
		Authorization: tokenString,
		Type:          "bearer",
	}
	err = utils.SetToken(tokenString, user.UserID)
	if err != nil {
		return err
	}
	return context.JSON(http.StatusOK, token)

}

// GetUser godoc
// @Summary GetUser
// @Description get username from token
// @Tags gormDB
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /gormdb/user [get]
func (i *Interactor) GetUserGormdb(context echo.Context) error {

	authercations := context.Request().Header.Get("token")
	user := utils.ParseToken(authercations)
	userID := user.Claims.(jwt.MapClaims)["userID"].(string)
	if !utils.GetToken(authercations, userID) {
		return context.String(http.StatusForbidden, "token awrong")
	}
	return context.JSON(http.StatusOK, userID)

}

// AddUser godoc
// @Summary AddUser
// @Description Add new user to database
// @Tags gormDB
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorizationc"
// @Param token body model.AddUser true "model.AddUser"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /gormdb/adduser [post]
func (i *Interactor) AddUserGormdb(context echo.Context) error {
	authercations := context.Request().Header.Get("Authorization")
	user := utils.ParseToken(authercations)
	userID := user.Claims.(jwt.MapClaims)["userID"].(string)
	if !utils.GetToken(authercations, userID) {
		return context.String(http.StatusForbidden, "Authorization awrong")
	}

	var Adduser model.AddUser
	err := context.Bind(&Adduser)

	if err != nil {
		return context.String(http.StatusBadRequest, "no user")
	}
	if userID == Adduser.UserID {
		return context.String(http.StatusBadRequest, "user exists")
	}
	cryptPassword := utils.CryptPassword(Adduser.Password)
	err = gormdb.Begin().Error
	if err != nil {
		return err
	}
	err = i.referrance.AddtUserGormdb(Adduser.UserID, *cryptPassword)
	if err != nil {
		return err
	}
	err = gormdb.Commit().Error
	if err != nil {
		return err
	}
	return context.String(http.StatusOK, "susscess")
}

// Getfirebase godoc
// @Summary Getfirebase
// @Description getfirebase
// @Tags Firebase
// @Accept json
// @Produce json
// @Success 201 {object} string
// @Failure 400 {object} error
// @Router /firebase/getfirebase [get]
func (i *Interactor) Getfirebase(c echo.Context) error {
	result, err := i.referrance.Getfirebase()
	if err != nil {
		return err
	}
	jsonBody, err := json.Marshal(result)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, string(jsonBody))

}

// Putfirebase godoc
// @Summary Putfirebase
// @Description putfirebase
// @Tags Firebase
// @Accept json
// @Produce json
// @Success 201 {object} string
// @Failure 400 {object} error
// @Router /firebase/putfirebase [post]
func (i *Interactor) Putfirebase(c echo.Context) error {
	err := i.referrance.Putfirebase()
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "susscess")

}

// LoginUser godoc
// @Summary LoginUser
// @Description login username
// @Tags gormDB
// @Accept json
// @Produce json
// @Param user body  model.Login true "model.Login"
// @Success 201 {object} model.Token
// @Failure 400 {object} string
// @Router /login [post]
func (i *Interactor) LoginUser(context echo.Context) error {
	InforLog.PrintLog(fmt.Sprintf("LoginUser start"))
	var user model.Login
	err := context.Bind(&user)
	if err != nil {
		return context.String(http.StatusBadRequest, errormessage.PrintError("3", err).Error())
	}
	result, err := i.referrance.GetUser(user.UserID, *utils.CryptPassword(user.Password))
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}
	if result == nil {
		return context.JSON(http.StatusBadRequest, errormessage.PrintError("2", err).Error())
	}

	tokenString := utils.GenerateToken(result.UserID)
	token := &model.Token{
		Authorization: tokenString,
		Type:          "bearer",
	}
	err = utils.SetToken(tokenString, user.UserID)
	if err != nil {
		return context.String(http.StatusBadRequest, errormessage.PrintError("5", err).Error())
	}
	return context.JSON(http.StatusOK, token)

}

// AddUser godoc
// @Summary AddUser
// @Description Add new user to database
// @Tags gormDB
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param token body model.AddUser true "model.AddUser"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /adduser [post]
func (i *Interactor) AddUser(context echo.Context) error {
	authercations := context.Request().Header.Get("Authorization")
	user := utils.ParseToken(authercations)
	userID := user.Claims.(jwt.MapClaims)["userID"].(string)
	if !utils.GetToken(authercations, userID) {
		return context.String(http.StatusForbidden, "token awrong")
	}

	var Adduser model.AddUser
	err := context.Bind(&Adduser)

	if err != nil {
		return context.String(http.StatusBadRequest, "no user")
	}

	cryptPassword := utils.CryptPassword(Adduser.Password)
	result, err := i.referrance.GetUser(Adduser.UserID, *cryptPassword)
	if err != nil {
		return err
	}

	if result != nil {
		return context.String(http.StatusBadRequest, "user exist")
	}
	err = i.referrance.AddUser(Adduser.UserID, *cryptPassword)
	if err != nil {
		return err
	}
	return context.String(http.StatusOK, "susscess")
}

// GetMovies godoc
// @Summary GetMovies
// @Description login username
// @Tags MonggoDB
// @Accept json
// @Produce json
// @Success 201 {object} model.Movies
// @Failure 400 {object} error
// @Router /getmovies [get]
func (i *Interactor) GetMovies(c echo.Context) error {
	result, err := i.referrance.GetMovies()
	if err != nil {
		return err
	}
	jsonBody, err := json.Marshal(result)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, string(jsonBody))
}

// Putmongo godoc
// @Summary Putmongo
// @Description login username
// @Tags MonggoDB
// @Accept json
// @Produce json
// @Success 201 {object} string
// @Failure 400 {object} error
// @Router /addmovies [post]
func (i *Interactor) PutMovies(context echo.Context) error {
	authercations := context.Request().Header.Get("Authorization")
	user := utils.ParseToken(authercations)
	userID := user.Claims.(jwt.MapClaims)["userID"].(string)
	if !utils.GetToken(authercations, userID) {
		return context.String(http.StatusForbidden, "token awrong")
	}
	file, err := context.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	location, err := utils.AddManager("test", uuid.New().String()+file.Filename, src)
	if err != nil {
		return err
	}
	err = i.referrance.AddMovies(file.Filename, "test", location)
	if err != nil {
		return err
	}
	return context.String(http.StatusOK, "susscess")
}
