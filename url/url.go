package url

import (
	"math/rand"
	"net/url"
	"time"
)

const (
	tamanho  = 5
	simbolos = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_-+"
)

//Repositorio tipo
type Repositorio interface {
	IDExiste(id string) bool
	BuscarPorID(id string) *Url
	BuscarPorURL(url string) *Url
	Salvar(url Url) error
	RegistrarClick(id string)
	BuscarClicks(id string) int
}

//Url tipo
type Url struct {
	ID      string    `json:"id"`
	Criacao time.Time `json:"criacao"`
	Destino string    `json:"destino"`
}

//Stats estrutura das estatísticas
type Stats struct {
	URL    *Url `json:"url"`
	Clicks int  `json:"clicks"`
}

var repo Repositorio

func init() {
	//configurar semente para geração dos numeros aleatórios
	rand.Seed(time.Now().UnixNano())
}

//ConfigurarRepositorio indica o repositorio a utilizar
func ConfigurarRepositorio(r Repositorio) {
	repo = r
}

//Buscar busca no repositorio se a url curta já existe
func Buscar(id string) *Url {
	return repo.BuscarPorID(id)
}

//BuscarOuCriarNovaURL busca ou cria nova url curta
func BuscarOuCriarNovaURL(destino string) (*Url, bool, error) {
	if u := repo.BuscarPorURL(destino); u != nil {
		return u, false, nil
	}

	if _, err := url.ParseRequestURI(destino); err != nil {
		return nil, false, err
	}

	url := Url{gerarID(), time.Now(), destino}
	repo.Salvar(url)
	return &url, true, nil
}

//RegistrarClick colher estatísticas dos clicks
func RegistrarClick(id string) {
	repo.RegistrarClick(id)
}

//Stats retorna as estatisitcas de uma URL
func (url *Url) Stats() *Stats {
	clicks := repo.BuscarClicks(url.ID)
	return &Stats{url, clicks}
}

func gerarID() string {
	novoID := func() string {
		id := make([]byte, tamanho, tamanho)
		for i := range id {
			id[i] = simbolos[rand.Intn(len(simbolos))]
		}
		return string(id)
	}

	for {
		if id := novoID(); !repo.IDExiste(id) {
			return id
		}
	}
}
