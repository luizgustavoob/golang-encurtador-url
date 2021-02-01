package response

import (
	"fmt"
	"net/http"
)

type Headers map[string]string

func SendResponseWith(w http.ResponseWriter, status int, headers Headers) {
	for i, header := range headers {
		w.Header().Set(i, header)
	}
	w.WriteHeader(status)
}

func SendResponseWithJSON(w http.ResponseWriter, json string) {
	SendResponseWith(w, http.StatusOK, Headers{"Content-Type": "application/json"})
	fmt.Fprintf(w, json)
}
