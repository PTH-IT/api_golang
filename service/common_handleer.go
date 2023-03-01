package af

import (
	"github.com/PTH-IT/api_golang/usecase"

	"github.com/labstack/echo/v4"
)

type commonhandler struct {
	Interactor *usecase.Interactor
}

func AppV1GetUsersGormdb(api commonhandler) echo.HandlerFunc {
	return func(context echo.Context) error {
		return api.Interactor.GetUserGormdb(context)
	}

}
func AppV1PostLoginGormdb(api commonhandler) echo.HandlerFunc {
	return func(context echo.Context) error {
		return api.Interactor.LoginUserGormdb(context)
	}

}

func AppV1AddUserGormdb(api commonhandler) echo.HandlerFunc {

	return func(context echo.Context) error {
		return api.Interactor.AddUserGormdb(context)
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

func AppV1AddMovies(api commonhandler) echo.HandlerFunc {

	return func(context echo.Context) error {
		return api.Interactor.PutMovies(context)
	}
}
func AppV1GetMovies(api commonhandler) echo.HandlerFunc {

	return func(context echo.Context) error {
		return api.Interactor.GetMovies(context)
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
