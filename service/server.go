package af

import (
	"fmt"
	"io"
	"net/http"

	firebasedb "PTH-IT/api_golang/adapter/firebasedb"
	gormdb "PTH-IT/api_golang/adapter/gormdb"
	"PTH-IT/api_golang/adapter/monggodb"
	config "PTH-IT/api_golang/config"
	usecase "PTH-IT/api_golang/usecase"
	"PTH-IT/api_golang/utils"

	InforLog "PTH-IT/api_golang/log/infor"

	"github.com/golang-jwt/jwt/v4"
	echo "github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Run() {
	InforLog.PrintLog(fmt.Sprintf("echo.New call"))
	e := echo.New()

	// connectString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
	// 	config.Getconfig().Mysql.User,
	// 	config.Getconfig().Mysql.Pass,
	// 	config.Getconfig().Mysql.Host,
	// 	config.Getconfig().Mysql.Port,
	// 	config.Getconfig().Mysql.Db,
	// )
	var err error
	// gormDb, err := gorm.Open(mysql.Open(connectString), &gorm.Config{
	// 	Logger: logger.Default.LogMode(logger.Info),
	// })
	if err != nil {
		panic(err)
	}
	// gormdb.Start(gormDb)
	userRepository := gormdb.NewUser()
	mongoRepository := monggodb.NewMongoDriver()
	firebaseRepository := firebasedb.NewFirebaseRepository()
	referrance := usecase.NewReferrance(userRepository, mongoRepository, firebaseRepository)
	// interactor := usecase.NewInteractor(gormDb, referrance)
	interactor := usecase.NewInteractor(referrance)

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

func Checktoken(context echo.Context) error {
	authercations := context.Request().Header.Get("Authorization")
	user := utils.ParseToken(authercations)
	userID := user.Claims.(jwt.MapClaims)["userID"].(string)
	if !utils.GetToken(authercations, userID) {
		return context.String(http.StatusForbidden, "token awrong")
	}
	return nil
}
