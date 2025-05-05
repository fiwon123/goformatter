package logger

import (
	"log"
	"os"
	"time"
)

var instance Logger

type Logger struct {
	DisableLog bool
	LogName    string
}

func NewLogger(disableLog bool) {
	instance = Logger{
		DisableLog: disableLog,
		LogName:    "./logs/" + time.Now().Format("2006.01.02_1504") + ".log"}
}

func SaveLog(msg string) {
	if instance.DisableLog {
		return
	}

	f, err := os.OpenFile(instance.LogName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)
	log.Println(msg)
}
