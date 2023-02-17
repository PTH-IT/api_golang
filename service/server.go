package af

import (
	"fmt"
	"io"

	config "PTH-IT/api_golang/config"
	firebasedb "PTH-IT/api_golang/database/firebasedb"
	gormdb "PTH-IT/api_golang/database/gormdb"
	"PTH-IT/api_golang/database/monggodb"
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
	firebaseRepository := firebasedb.NewFirebaseRepository()
	referrance := usecase.NewReferrance(userRepository, mongoRepository, firebaseRepository)
	interactor := usecase.NewInteractor(gormDb, referrance)

	api := commonhandler{
		Interactor: &interactor,
	}
	g := e.Group("/gormdb")

	g.POST("/login", AppV1PostLoginGormdb(api))
	g.GET("/user", AppV1GetUsersGormdb(api))
	g.POST("/adduser", AppV1AddUserGormdb(api))
	f := e.Group("/firebase")

	f.GET("/getfirebase", AppV1GetFirebase(api))
	f.POST("/putfirebase", AppV1PutFirebase(api))

	e.POST("/login", AppV1PostLogin(api))
	e.POST("/adduser", AppV1AddUser(api))
	e.POST("/addmovies", AppV1AddMovies(api))
	e.GET("/getmovies", AppV1GetMovies(api))

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.SetOutput(io.Discard)
	e.Logger.Fatal(e.Start(config.Getconfig().Port))
}
