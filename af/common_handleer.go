package af

import (
	"github.com/PTH-IT/api_golang/usecase"
	"github.com/labstack/echo"
)

type commonhandler struct {
	Interactor *usecase.Interactor
}

func AppV1GetUsers(api commonhandler) echo.HandlerFunc {
	return func(context echo.Context) error {
		return api.Interactor.GetUser(context)
	}

}
