package main

import (
	"fmt"

	"PTH-IT/api_golang/af"
	"PTH-IT/api_golang/utils"
)

func main() {
	fmt.Println("server.Run call")
	utils.UpdateManager("test", "hau", "pth-it-firebase-adminsdk-i11h0-ba3394f404.json")
	utils.DownloadManager("test", "hau", "tesst.json")
	af.Run()
}
