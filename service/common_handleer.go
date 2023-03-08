package af

import (
	"PTH-IT/api_golang/usecase"

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
func AppV1RegisterUser(api commonhandler) echo.HandlerFunc {

	return func(context echo.Context) error {
		return api.Interactor.RegisterUser(context)
	}
}

func AppV1SocketMessage(api commonhandler) echo.HandlerFunc {

	return func(context echo.Context) error {
		return api.Interactor.SocketMessage(context)
	}
}

func AppV1GetLogout(api commonhandler) echo.HandlerFunc {

	return func(context echo.Context) error {
		return api.Interactor.GetLogout(context)
	}
}
func AppV1GetMessage(api commonhandler) echo.HandlerFunc {

	return func(context echo.Context) error {
		return api.Interactor.GetMessage(context)
	}
}
func AppV1SaveMessage(api commonhandler) echo.HandlerFunc {

	return func(context echo.Context) error {
		return api.Interactor.SaveMessage(context)
	}
}
