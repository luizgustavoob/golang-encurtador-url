package url

import (
	"net/http"
	"time"
)

type Url struct {
	ID          string    `json:"id"`
	CreatedAt   time.Time `json:"criacao"`
	Destination string    `json:"destino"`
}

type Statistics struct {
	URL    *Url `json:"url"`
	Clicks int  `json:"clicks"`
}

type Redirect struct {
	ChanStats chan string
}

func Configure(port *int, repository Repository) {
	ConfigureBaseURL(port)
	ConfigureRepository(repository)
}

func (self *Redirect) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	findURLAndExecute(w, r, func(u *Url) {
		http.Redirect(w, r, u.Destination, http.StatusMovedPermanently)
		self.ChanStats <- u.ID
	})
}
