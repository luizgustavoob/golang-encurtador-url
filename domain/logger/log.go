package logger

import (
	"fmt"
	"log"
)

type logStr struct {
	active bool
}

var mylog *logStr

func init() {
	mylog = &logStr{active: true}
}

func Configure(active *bool) {
	mylog.active = *active
}

func (l *logStr) logar(formato string, valores ...interface{}) {
	if l.active {
		log.Printf(fmt.Sprintf("%s\n", formato), valores...)
	}
}

func Logar(formato string, valores ...interface{}) {
	mylog.logar(formato, valores)
}
