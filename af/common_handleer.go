package af

import (
	"PTH-IT/api_golang/usecase"

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
func AppV1PostLogin(api commonhandler) echo.HandlerFunc {
	return func(context echo.Context) error {
		return api.Interactor.LoginUser(context)
	}

}

func AppV1AddUser(api commonhandler) echo.HandlerFunc {

	return func(context echo.Context) error {
		return api.Interactor.AddUser(context)
	}
}
