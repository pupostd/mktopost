package main

import (
	"github.com/pupostd/mktopost/model"
	"log"
	"net/http"
	"strings"
)

type Handler struct {
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path[1:]
	log.Println(path)

	s := model.Searcher{}

	var name = path
	st := strings.Split(path, ".")
	if len(st) > 1 {
		name = st[0]
	}

	data := s.LookFor(name)

	if data != "" {
		w.Write([]byte(data))
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Bro ... " + http.StatusText(http.StatusNotFound)))
	}

}

func main() {
	http.Handle("/", new(Handler))
	http.ListenAndServe(":8080", nil)
}
