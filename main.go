package main

import (
	"fmt"

	"PTH-IT/api_golang/af"
)

func main() {
	fmt.Println("server.Run call")
	// fmt.Println(utils.GetscretManager("test"))
	// utils.UpdateManager("test", "hau", "pth-it-firebase-adminsdk-i11h0-ba3394f404.json")
	// utils.DownloadManager("test", "hau", "tesst.json")
	af.Run()
}
