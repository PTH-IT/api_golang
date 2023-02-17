package main

import (
	"fmt"

	_ "PTH-IT/api_golang/docs"
	InforLog "PTH-IT/api_golang/log/infor"
	af "PTH-IT/api_golang/service"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email haupham404
// @host localhost:1909
func main() {
	InforLog.PrintLog(fmt.Sprintf("af.Run call"))
	// fmt.Println(utils.GetscretManager("test"))
	// utils.UpdateManager("test", "hau", "pth-it-firebase-adminsdk-i11h0-ba3394f404.json")
	// utils.DownloadManager("test", "hau", "tesst.json")
	af.Run()
}
