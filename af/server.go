package af

import (
	"fmt"

	"github.com/PTH-IT/api_golang/adapter/api"
	usecase "github.com/PTH-IT/api_golang/usecase"
	"github.com/labstack/echo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Run() {
	e := echo.New()
	userName := ""
	passWord := ""
	host := ""
	port := ""
	dataBaseName := ""
	connectString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		userName,
		passWord,
		host,
		port,
		dataBaseName,
	)
	var err error
	gormDb, err := gorm.Open(mysql.Open(connectString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
	}

	userRepository := api.NewUser()
	referrance := usecase.NewReferrance(userRepository)
	interactor := usecase.NewInteractor(gormDb, referrance)

	api := commonhandler{
		Interactor: &interactor,
	}
	e.GET("/user", AppV1GetUsers(api))
	e.Logger.Fatal(e.Start(":80"))
}
