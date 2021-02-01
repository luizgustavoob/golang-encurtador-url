package url

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	mylog "github.com/golang-encurtador-url/log"
	"github.com/golang-encurtador-url/response"
)

var baseUrl string

func ConfigureBaseURL(port *int) {
	baseUrl = fmt.Sprintf("http://localhost:%d", *port)
}

func Shorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.SendResponseWith(w, http.StatusMethodNotAllowed, response.Headers{"Allow": http.MethodPost})
		return
	}

	u, created, err := FindOrCreateURL(extractURL(r))
	if err != nil {
		response.SendResponseWith(w, http.StatusBadRequest, nil)
	}

	var status int
	if created {
		status = http.StatusCreated
	} else {
		status = http.StatusOK
	}

	shortURL := fmt.Sprintf("%s/r/%s", baseUrl, u.ID)

	response.SendResponseWith(w, status, response.Headers{
		"Location": shortURL,
		"Link":     fmt.Sprintf("<%s/api/stats/%s>; rel=\"stats\"", baseUrl, u.ID),
	})

	mylog.Logar("URL %s encurtada com sucesso para %s", u.Destination, shortURL)
}

func Visualize(w http.ResponseWriter, r *http.Request) {
	findURLAndExecute(w, r, func(u *Url) {
		js, err := json.Marshal(GetStatistics(u))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		response.SendResponseWithJSON(w, string(js))
	})
}

func findURLAndExecute(w http.ResponseWriter, r *http.Request, executor func(*Url)) {
	path := strings.Split(r.URL.Path, "/")
	id := path[len(path)-1]

	if u := Find(id); u != nil {
		executor(u)
	} else {
		http.NotFound(w, r)
	}
}

func extractURL(r *http.Request) string {
	u := make([]byte, r.ContentLength, r.ContentLength)
	r.Body.Read(u)
	return string(u)
}
