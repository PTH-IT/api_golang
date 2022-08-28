package af

import (
	"fmt"

	usecase "github.com/PTH-IT/api_golang/usecase"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Run() {

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
	interactor := usecase.NewInteractor(gormDb)
	err = interactor.Interactor()
	if err != nil {
		fmt.Println(err.Error())

	}
}
