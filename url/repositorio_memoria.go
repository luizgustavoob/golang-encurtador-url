package url

type repositorioMemoria struct {
	urls   map[string]*Url
	clicks map[string]int
}

//NovoRepositorioMemoria cria um novo repositório em memória
func NovoRepositorioMemoria() *repositorioMemoria {
	return &repositorioMemoria{
		make(map[string]*Url),
		make(map[string]int),
	}
}

//IDExiste verifica se a url curta existe no repositorio
func (r *repositorioMemoria) IDExiste(id string) bool {
	_, existe := r.urls[id]
	return existe
}

//BuscarPorID retorna o ponteiro pra URL a partir do id curto
func (r *repositorioMemoria) BuscarPorID(id string) *Url {
	return r.urls[id]
}

//BuscarPorURL retorna um ponteiro pra URL a partir da url destino
func (r *repositorioMemoria) BuscarPorURL(destino string) *Url {
	for _, u := range r.urls {
		if u.Destino == destino {
			return u
		}
	}
	return nil
}

//Salvar adiciona nova URL no map do repositorio
func (r *repositorioMemoria) Salvar(url Url) error {
	r.urls[url.ID] = &url
	return nil
}

func (r *repositorioMemoria) RegistrarClick(id string) {
	r.clicks[id]++
}

func (r *repositorioMemoria) BuscarClicks(id string) int {
	return r.clicks[id]
}
