package InforLog

import (
	"encoding/json"
	"fmt"
	"path"
	"runtime"
	"time"
)

type MyInfo struct {
	Level    string `json:"level"`
	Date     string `json:"date"`
	Location string `json:"location"`
	Info     string `json:"info"`
}

func PrintLog(message string) {
	_, file, lineNo, _ := runtime.Caller(1)
	myinfoMessage := MyInfo{
		Level:    "info",
		Date:     time.Now().String(),
		Location: fmt.Sprintf("%s:%d  ", path.Base(file), lineNo),
		Info:     message,
	}
	loggmessage, err := json.Marshal(myinfoMessage)
	if err != nil {
		panic(err)
	}
	fmt.Printf(string(loggmessage))
}
