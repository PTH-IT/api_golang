package errormessage

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"time"
)

type Message struct {
	MessageId      string `json:"messageId"`
	MessageCode    string `json:"messageCode"`
	MessageDetails string `json:"messageDetails"`
}
type MyError struct {
	Level        string `json:"level"`
	Date         string `json:"date"`
	Location     string `json:"location"`
	MessageError string `json:"MessageError"`
	MyMessage    string `json:"error"`
}

func PrintError(messageId string, errMessage error) error {

	messageErrorApp := make(map[string]Message)
	err := json.Unmarshal([]byte(messageApp), &messageErrorApp)
	if err != nil {
		panic(err)
	}
	getMessage, check := messageErrorApp[messageId]
	if !check {
		panic(err)
	}
	stringmessage, err := json.Marshal(getMessage)
	if err != nil {
		panic(err)
	}
	_, file, lineNo, _ := runtime.Caller(1)

	myErrorMessage := MyError{
		Level:        "error",
		Date:         time.Now().String(),
		Location:     fmt.Sprintf("%s:%d", path.Base(file), lineNo),
		MessageError: errMessage.Error(),
		MyMessage:    string(stringmessage),
	}
	loggmessage, err := json.Marshal(myErrorMessage)
	if err != nil {
		panic(err)
	}

	var Error = log.New(os.Stdout, "\u001b[31m ", 0)

	Error.Println(string(loggmessage) + "\r")

	return fmt.Errorf("%v", string(stringmessage))

}
