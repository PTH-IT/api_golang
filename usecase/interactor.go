package usecase

import (
	"encoding/json"
	"fmt"
	"net/http"

	gormdb "PTH-IT/api_golang/adapter/gormdb"
	"PTH-IT/api_golang/domain/model"
	errormessage "PTH-IT/api_golang/log/error"
	"PTH-IT/api_golang/utils"

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
	InforLog.PrintLog(fmt.Sprintf("LoginUserGormdb start"))

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
	InforLog.PrintLog(fmt.Sprintf("GetUserGormdb start"))

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
// @Param token body model.RegisterUser true "model.RegisterUser"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /gormdb/adduser [post]
func (i *Interactor) AddUserGormdb(context echo.Context) error {
	InforLog.PrintLog(fmt.Sprintf("AddUserGormdb start"))

	authercations := context.Request().Header.Get("Authorization")
	user := utils.ParseToken(authercations)
	userID := user.Claims.(jwt.MapClaims)["userID"].(string)
	if !utils.GetToken(authercations, userID) {
		return context.String(http.StatusForbidden, "Authorization awrong")
	}

	var Adduser model.RegisterUser
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
	InforLog.PrintLog(fmt.Sprintf("Getfirebase start"))

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
	InforLog.PrintLog(fmt.Sprintf("Putfirebase start"))

	err := i.referrance.Putfirebase()
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "susscess")

}

// LoginUser godoc
// @Summary LoginUser
// @Description login username
// @Tags MonggoDB
// @Accept json
// @Produce json
// @Param Content-Type header string true "application/json" default(application/json)
// @Param Content-Length header string true "1000" default(1000)
// @Param Host header string true "localhost" default(localhost)
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
	InforLog.PrintLog(fmt.Sprintf("reqest user: %v", user))

	result, err := i.referrance.GetUser(user.UserID, *utils.CryptPassword(user.Password))
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	}
	if result == nil {

		return context.JSON(http.StatusForbidden, model.MessageCheckUser{Type: "user", Message: "username or password is not correst"})
	}

	tokenString := utils.GenerateToken(result.UserID)
	token := &model.Token{
		Authorization: tokenString,
		Type:          "bearer",
	}
	InforLog.PrintLog(fmt.Sprintf("response token: %v", token))

	err = utils.SetToken(tokenString, user.UserID)
	if err != nil {
		return context.String(http.StatusBadRequest, errormessage.PrintError("5", err).Error())
	}
	return context.JSON(http.StatusOK, token)

}

// AddUser godoc
// @Summary AddUser
// @Description Add new user to database
// @Tags MonggoDB
// @Accept json
// @Produce json
// @Param Content-Type header string true "application/json" default(application/json)
// @Param Content-Length header string true "1000" default(1000)
// @Param Host header string true "localhost" default(localhost)
// @Param token body model.RegisterUser true "model.RegisterUser"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /adduser [post]
func (i *Interactor) RegisterUser(context echo.Context) error {
	InforLog.PrintLog(fmt.Sprintf("RegisterUser start"))

	var Adduser model.RegisterUser
	err := context.Bind(&Adduser)

	if err != nil || Adduser.Email == "" || Adduser.UserID == "" || Adduser.Password == "" {
		errData := map[string]interface{}{
			"message": "request body is invalid",
		}
		return context.JSON(http.StatusBadRequest, errData)
	}
	InforLog.PrintLog(fmt.Sprintf("reqest user: %v", Adduser))

	cryptPassword := utils.CryptPassword(Adduser.Password)
	result, err := i.referrance.CheckUserName(Adduser.UserID, Adduser.Email)
	if err != nil {
		return err
	}

	if result != nil {
		var messageError []model.MessageCheckUser
		for _, r := range result {
			if r.UserID == Adduser.UserID {
				messageError = append(messageError, model.MessageCheckUser{Type: "username", Message: "username is exist "})
			}
			if r.Email == Adduser.Email {
				messageError = append(messageError, model.MessageCheckUser{Type: "email", Message: "email is exist"})
			}

		}
		return context.JSON(http.StatusBadRequest, messageError)
	}
	err = i.referrance.AddUser(Adduser.UserID, *cryptPassword, Adduser.Email)
	if err != nil {
		return err
	}
	return context.NoContent(http.StatusOK)
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
	InforLog.PrintLog(fmt.Sprintf("GetMovies start"))

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
// @Param Authorization header string true "Authorization"
// @Success 201 {object} string
// @Failure 400 {object} error
// @Router /addmovies [post]
func (i *Interactor) PutMovies(context echo.Context) error {
	InforLog.PrintLog(fmt.Sprintf("PutMovies start"))

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

// GetLogout godoc
// @Summary GetLogout
// @Description GetLogout
// @Tags MonggoDB
// @Accept json
// @Produce json
// @Param Content-Type header string true "application/json" default(application/json)
// @Param Content-Length header string true "1000" default(1000)
// @Param Host header string true "localhost" default(localhost)
// @Param Authorization header string true "Authorization"
// @Success 200 {object} string
// @Failure 400 {object} error
// @Router /logout [get]
func (i *Interactor) GetLogout(context echo.Context) error {
	InforLog.PrintLog(fmt.Sprintf("GetLogout start"))

	authercations := context.Request().Header.Get("Authorization")
	InforLog.PrintLog(fmt.Sprintf("authercations = %s", authercations))

	user := utils.ParseToken(authercations)
	InforLog.PrintLog(fmt.Sprintf("user = %v", user))

	userID := user.Claims.(jwt.MapClaims)["userID"].(string)
	InforLog.PrintLog(fmt.Sprintf("userID = %s", userID))
	InforLog.PrintLog(fmt.Sprintf("utils.GetToken"))

	if !utils.GetToken(authercations, userID) {
		return context.String(http.StatusForbidden, "token awrong")
	}
	InforLog.PrintLog(fmt.Sprintf("utils.DeleteToken"))

	if !utils.DeleteToken(authercations, userID) {
		return context.String(http.StatusBadRequest, "Can not delete token")
	}
	InforLog.PrintLog(fmt.Sprintf("StatusOK"))

	return context.String(http.StatusOK, "susscess")
}

// SaveMessage godoc
// @Summary SaveMessage
// @Description SaveMessage
// @Tags MonggoDB
// @Accept json
// @Produce json
// @Param Content-Type header string true "application/json" default(application/json)
// @Param Content-Length header string true "1000" default(1000)
// @Param Host header string true "localhost" default(localhost)
// @Param Authorization header string true "Authorization"
// @Param token body model.Message true "model.Message"
// @Success 200 {object} string
// @Failure 400 {object} error
// @Router /savemessage [post]
func (i *Interactor) SaveMessage(context echo.Context) error {
	InforLog.PrintLog(fmt.Sprintf("SaveMessage start"))

	authercations := context.Request().Header.Get("Authorization")
	user := utils.ParseToken(authercations)
	userID := user.Claims.(jwt.MapClaims)["userID"].(string)
	if !utils.GetToken(authercations, userID) {
		return context.String(http.StatusForbidden, "token awrong")
	}
	var message *model.Message
	err := context.Bind(&message)

	if err != nil || message.Receiver == "" || message.Sender == "" || message.Detail == "" {
		errData := map[string]interface{}{
			"message": "request body is invalid",
		}
		return context.JSON(http.StatusBadRequest, errData)
	}
	message.Time = utils.GettimeNumber()
	err = i.referrance.SaveMessage(message)

	return context.String(http.StatusOK, "susscess")
}

// GetMessage godoc
// @Summary GetMessage
// @Description GetMessage
// @Tags MonggoDB
// @Accept json
// @Produce json
// @Param Content-Type header string true "application/json" default(application/json)
// @Param Content-Length header string true "1000" default(1000)
// @Param Host header string true "localhost" default(localhost)
// @Param message body model.InputGetMessage true "model.InputGetMessage"
// @Success 200 {object} string
// @Failure 400 {object} error
// @Router /message [post]
func (i *Interactor) GetMessage(context echo.Context) error {
	InforLog.PrintLog(fmt.Sprintf("GetMessage start"))

	authercations := context.Request().Header.Get("Authorization")
	user := utils.ParseToken(authercations)
	userID := user.Claims.(jwt.MapClaims)["userID"].(string)
	if !utils.GetToken(authercations, userID) {
		return context.String(http.StatusForbidden, "token awrong")
	}
	var message *model.InputGetMessage
	err := context.Bind(&message)

	if err != nil || message.Receiver == "" || message.Sender == "" {
		errData := map[string]interface{}{
			"message": "request body is invalid",
		}
		return context.JSON(http.StatusBadRequest, errData)
	}

	result, err := i.referrance.Getmessage(message)
	if err != nil {
		return context.String(http.StatusBadRequest, err.Error())
	} else {
		return context.JSON(http.StatusOK, result)
	}

}
