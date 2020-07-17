package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"

	url "livro-go/github.com/luizgustavoob/encurtador/url"
)

var (
	porta    *int
	logAtivo *bool
	urlBase  string
)

//Headers definição de tipo customizado
type Headers map[string]string

//Go disponibiliza um mecanismo padrão para inicialização de variáveis em um pacote: a função init()
func init() {
	porta = flag.Int("p", 8888, "porta")
	logAtivo = flag.Bool("l", true, "log ativo/inativo")
	flag.Parse()

	urlBase = fmt.Sprintf("http://localhost:%d", *porta)
}

//Redirecionador tipo
type Redirecionador struct {
	stats chan string
}

func (red *Redirecionador) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	buscarURLEExecutar(w, r, func(u *url.Url) {
		http.Redirect(w, r, u.Destino, http.StatusMovedPermanently)
		red.stats <- u.ID
	})
}

//Encurtador função para encurtar a URL
func Encurtador(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		responderCom(w, http.StatusMethodNotAllowed, Headers{"Allow": "POST"})
		return
	}

	url, created, err := url.BuscarOuCriarNovaURL(extrairURL(r))
	if err != nil {
		responderCom(w, http.StatusBadRequest, nil)
	}

	var status int
	if created {
		status = http.StatusCreated
	} else {
		status = http.StatusOK
	}

	urlCurta := fmt.Sprintf("%s/r/%s", urlBase, url.ID)
	responderCom(w, status, Headers{
		"Location": urlCurta,
		"Link":     fmt.Sprintf("<%s/api/stats/%s>; rel=\"stats\"", urlBase, url.ID),
	})
	logar("URL %s encurtada com sucesso para %s.", url.Destino, urlCurta)
}

//Visualizador visualiza as estatisticas
func Visualizador(w http.ResponseWriter, r *http.Request) {
	buscarURLEExecutar(w, r, func(u *url.Url) {
		json, err := json.Marshal(u.Stats())

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		responderComJSON(w, string(json))
	})
}

func buscarURLEExecutar(w http.ResponseWriter, r *http.Request, executor func(*url.Url)) {
	caminho := strings.Split(r.URL.Path, "/")
	id := caminho[len(caminho)-1]

	if u := url.Buscar(id); u != nil {
		executor(u)
	} else {
		http.NotFound(w, r)
	}
}

func responderCom(w http.ResponseWriter, status int, headers Headers) {
	for i, header := range headers {
		w.Header().Set(i, header)
	}
	w.WriteHeader(status)
}

func responderComJSON(w http.ResponseWriter, json string) {
	responderCom(w, http.StatusOK, Headers{"Content-Type": "application/json"})
	fmt.Fprintf(w, json)
}

//Pega a url que veio no corpo da requisição e transforma numa string
func extrairURL(r *http.Request) string {
	url := make([]byte, r.ContentLength, r.ContentLength)
	r.Body.Read(url)
	return string(url)
}

func registrarEstatisticas(ids <-chan string) {
	for id := range ids {
		url.RegistrarClick(id)
		logar("Click registrado com sucesso para %s.", id)
	}
}

func logar(formato string, valores ...interface{}) {
	if *logAtivo {
		log.Printf(fmt.Sprintf("%s\n", formato), valores...)
	}
}

func main() {
	url.ConfigurarRepositorio(url.NovoRepositorioMemoria())

	stats := make(chan string)
	defer close(stats) // quando main morrer, channel morre tbm
	go registrarEstatisticas(stats)

	//rotas
	http.Handle("/r/", &Redirecionador{stats})
	http.HandleFunc("/api/encurtar", Encurtador)
	http.HandleFunc("/api/stats/", Visualizador)

	logar("Iniciando servidor na porta %d", *porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *porta), nil))
}
