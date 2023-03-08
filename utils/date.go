package utils

import (
	"time"
)

func GettimeNow() string {
	return time.Now().UTC().Local().Format(time.RFC3339Nano)
}

func GettimeNumber() int64 {
	return time.Now().UTC().Unix()
}
