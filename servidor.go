package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	mylog "github.com/golang-encurtador-url/log"
	url "github.com/golang-encurtador-url/url"
)

var (
	port      *int
	activeLog *bool
)

func init() {
	port = flag.Int("p", 8888, "porta")
	activeLog = flag.Bool("l", true, "log ativo/inativo")
	flag.Parse()
}

func collectStats(ids <-chan string) {
	for id := range ids {
		url.AddLogClick(id)
		mylog.Logar("Click registrado com sucesso para %s.", id)
	}
}

func main() {
	mylog.ConfigureLog(activeLog)
	url.Configure(port, url.NewMemoryRepository())

	stats := make(chan string)
	defer close(stats)
	go collectStats(stats)

	http.Handle("/r/", &url.Redirect{ChanStats: stats})
	http.HandleFunc("/api/encurtar", url.Shorten)
	http.HandleFunc("/api/stats/", url.Visualize)

	mylog.Logar("Iniciando servidor na porta %d", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
