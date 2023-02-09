package af

import (
	"fmt"

	gormdb "PTH-IT/api_golang/adapter/gormdb"
	config "PTH-IT/api_golang/config"
	usecase "PTH-IT/api_golang/usecase"

	"github.com/labstack/echo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Run() {
	e := echo.New()

	connectString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Getconfig().Mysql.User,
		config.Getconfig().Mysql.Pass,
		config.Getconfig().Mysql.Host,
		config.Getconfig().Mysql.Port,
		config.Getconfig().Mysql.Db,
	)
	var err error
	gormDb, err := gorm.Open(mysql.Open(connectString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("Failed to connect database")
	}
	gormdb.Start(gormDb)
	userRepository := gormdb.NewUser()
	referrance := usecase.NewReferrance(userRepository)
	interactor := usecase.NewInteractor(gormDb, referrance)

	api := commonhandler{
		Interactor: &interactor,
	}

	e.GET("/user", AppV1GetUsers(api))
	e.POST("/login", AppV1PostLogin(api))
	e.POST("/adduser", AppV1AddUser(api))
	e.Logger.Fatal(e.Start(config.Getconfig().Port))
}
