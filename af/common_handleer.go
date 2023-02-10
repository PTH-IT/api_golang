package af

import (
	"PTH-IT/api_golang/usecase"

	"PTH-IT/api_golang/adapter/monggodb"

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

func AppV1AddMovies(api commonhandler) echo.HandlerFunc {

	return func(context echo.Context) error {
		return monggodb.Putmongo(context)
	}
}
func AppV1GetMovies(api commonhandler) echo.HandlerFunc {

	return func(context echo.Context) error {
		return monggodb.Getmongo(context)
	}
}
