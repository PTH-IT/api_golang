package af

import (
	"PTH-IT/api_golang/usecase"

	"github.com/labstack/echo/v4"
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
		return api.Interactor.Putmongo(context)
	}
}
func AppV1GetMovies(api commonhandler) echo.HandlerFunc {

	return func(context echo.Context) error {
		return api.Interactor.Getmongo(context)
	}
}

func AppV1GetFirebase(api commonhandler) echo.HandlerFunc {

	return func(context echo.Context) error {
		return api.Interactor.Getfirebase(context)
	}
}
func AppV1PutFirebase(api commonhandler) echo.HandlerFunc {

	return func(context echo.Context) error {
		return api.Interactor.Putfirebase(context)
	}
}
