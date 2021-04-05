package logger

import (
	"fmt"
	"log"
)

type myLogStruct struct {
	active bool
}

var mylog *myLogStruct

func init() {
	mylog = &myLogStruct{active: true}
}

func (l *myLogStruct) logar(formato string, valores ...interface{}) {
	if l.active {
		log.Printf(fmt.Sprintf("%s\n", formato), valores...)
	}
}

func Configure(active *bool) {
	mylog.active = *active
}

func Logar(formato string, valores ...interface{}) {
	mylog.logar(formato, valores)
}
