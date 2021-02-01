package mylog

import (
	"fmt"
	"log"
)

type logStruct struct {
	activeLog *bool
}

var myLog logStruct

func ConfigureLog(paramActveLog *bool) {
	myLog = logStruct{activeLog: paramActveLog}
}

func (self logStruct) logar(formato string, valores ...interface{}) {
	if *self.activeLog {
		log.Printf(fmt.Sprintf("%s\n", formato), valores...)
	}
}

func Logar(formato string, valores ...interface{}) {
	myLog.logar(formato, valores)
}
