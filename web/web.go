package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

type Handler struct {
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path[1:]
	log.Println(path)

	data, err := ioutil.ReadFile(string(path))

	if err == nil {
		w.Write(data)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 Bro ... " + http.StatusText(404)))
	}

}

func main() {
	http.Handle("/", new(Handler))
	http.ListenAndServe(":8080", nil)
}
