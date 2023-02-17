package usecase

import (
	"PTH-IT/api_golang/domain/model"
	"PTH-IT/api_golang/utils"
	"encoding/json"
	"net/http"

	gormdb "PTH-IT/api_golang/adapter/gormdb"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewInteractor(
	gormDb *gorm.DB,
	referrance Reference,

) Interactor {

	return Interactor{
		gormDb,
		referrance,
	}
}

type Interactor struct {
	gormDb     *gorm.DB
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
// @Router /login [post]
func (i *Interactor) LoginUser(context echo.Context) error {

	var user model.Login
	err := context.Bind(&user)
	if err != nil {
		return context.String(http.StatusBadRequest, "no user")
	}
	result, err := i.referrance.GetUser(user.UserID, user.Password)
	if err != nil {
		return err
	}
	if result == nil {
		return context.String(http.StatusBadRequest, "user no exist")
	}

	tokenString := utils.GenerateToken(result.UserID)
	token := &model.Token{
		Token: tokenString,
		Type:  "bearer",
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
// @Param token header string true "Token"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /user [get]
func (i *Interactor) GetUser(context echo.Context) error {

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
// @Param token header string true "Token"
// @Param token body model.AddUser true "model.AddUser"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /adduser [post]
func (i *Interactor) AddUser(context echo.Context) error {
	authercations := context.Request().Header.Get("token")
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
	if userID == Adduser.UserID {
		return context.String(http.StatusBadRequest, "user exists")
	}
	cryptPassword := utils.CryptPassword(Adduser.Password)
	err = gormdb.Begin().Error
	if err != nil {
		return err
	}
	err = i.referrance.AddtUser(Adduser.UserID, *cryptPassword)
	if err != nil {
		return err
	}
	err = gormdb.Commit().Error
	if err != nil {
		return err
	}
	return context.String(http.StatusOK, "susscess")
}

// Getmongo godoc
// @Summary Getmongo
// @Description login username
// @Tags MonggoDB
// @Accept json
// @Produce json
// @Success 201 {object} model.Movies
// @Failure 400 {object} error
// @Router /getmovies [get]
func (i *Interactor) Getmongo(c echo.Context) error {
	result, err := i.referrance.Getmongo()
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
func (i *Interactor) Putmongo(c echo.Context) error {
	err := i.referrance.Putmongo()
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "susscess")
}

// Getfirebase godoc
// @Summary Getfirebase
// @Description getfirebase
// @Tags Firebase
// @Accept json
// @Produce json
// @Success 201 {object} string
// @Failure 400 {object} error
// @Router /getfirebase [get]
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
// @Router /putfirebase [post]
func (i *Interactor) Putfirebase(c echo.Context) error {
	err := i.referrance.Putfirebase()
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "susscess")

}
