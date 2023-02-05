package af

import (
	"fmt"
	"log"
	"os"

	gormdb "PTH-IT/api_golang/adapter/gormdb"
	usecase "PTH-IT/api_golang/usecase"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Run() {
	e := echo.New()
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	userName := os.Getenv("DB_USER")
	passWord := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	pord := os.Getenv("DB_PORT")
	dataBaseName := os.Getenv("DB_NAME")
	connectString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		userName,
		passWord,
		host,
		pord,
		dataBaseName,
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
	e.Logger.Fatal(e.Start(":1909"))
}
