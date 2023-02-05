package usecase

import (
	"PTH-IT/api_golang/domain/model"
	"PTH-IT/api_golang/utils"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo"
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

func (i *Interactor) GetUser(context echo.Context) error {

	authercations := context.Request().Header.Get("token")
	user := utils.ParseToken(authercations)
	userID := user.Claims.(jwt.MapClaims)["userID"]

	return context.JSON(http.StatusOK, userID)

}
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
	return context.JSON(http.StatusOK, token)

}
