package af

import (
	"fmt"
	"io"

	firebasedb "PTH-IT/api_golang/adapter/firebaseDB"
	gormdb "PTH-IT/api_golang/adapter/gormdb"
	"PTH-IT/api_golang/adapter/monggodb"
	config "PTH-IT/api_golang/config"
	usecase "PTH-IT/api_golang/usecase"

	InforLog "PTH-IT/api_golang/log/infor"

	echo "github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Run() {
	InforLog.PrintLog(fmt.Sprintf("echo.New call"))
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
		panic(err)
	}
	gormdb.Start(gormDb)
	userRepository := gormdb.NewUser()
	mongoRepository := monggodb.NewMongoDriver()
	referrance := usecase.NewReferrance(userRepository, mongoRepository)
	interactor := usecase.NewInteractor(gormDb, referrance)

	api := commonhandler{
		Interactor: &interactor,
	}
	e.POST("/login", AppV1PostLogin(api))
	e.GET("/user", AppV1GetUsers(api))
	e.POST("/adduser", AppV1AddUser(api))
	e.POST("/addmovies", AppV1AddMovies(api))
	e.GET("/getmovies", AppV1GetMovies(api))
	e.GET("/getfirebase", firebasedb.Getfirebase)
	e.POST("/putfirebase", firebasedb.Putfirebase)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.SetOutput(io.Discard)
	e.Logger.Fatal(e.Start(config.Getconfig().Port))
}
