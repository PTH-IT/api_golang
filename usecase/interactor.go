package usecase

import (
	"net/http"

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

	result, err := i.referrance.GetUser()
	if err != nil {
		return err
	}
	return context.JSON(http.StatusOK, result)

}
