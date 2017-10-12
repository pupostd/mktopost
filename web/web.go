package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/pupostd/mktopost/model"
	"text/template"
	"os"
	"bytes"
	"io"
	"io/ioutil"
)

// Handler http requests.
type Handler struct {
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]

	if strings.HasPrefix(path, "res/") {
		bf, err := ioutil.ReadFile("../resource/" + path)
		if err != nil {
			log.Printf("Problems loading resource %s", path)
		}

		if bf != nil {
			w.Header().Set("Content-Type", "text/css")
			w.Write([]byte(bf))
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(h.wrapData("404 Bro ... " + http.StatusText(http.StatusNotFound))))
		}

	} else {
		log.Println(path)

		s := model.Searcher{}
		data := s.LookFor(h.parsePath(path))

		if data != "" {
			w.Write([]byte(h.wrapData(data)))
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(h.wrapData("404 Bro ... " + http.StatusText(http.StatusNotFound))))
		}
	}
}

func (h *Handler) parsePath(path string) string {
	var name = path
	st := strings.Split(path, ".")
	if len(st) > 1 {
		name = st[0]
	}
	return name
}

func (h *Handler) wrapData(data string) string {
	t, err := template.New("main").Parse(h.loadTemplate())

	if err != nil {
		log.Fatal("Problems wrapping data on main template.")
	}

	bf := bytes.NewBuffer(nil)
	err = t.Execute(bf, data)

	if err != nil {
		log.Fatal("Problems wrapping data on main template.")
	}

	return string(bf.Bytes())
}

func (h *Handler) loadTemplate() string {
	conf := model.Configuration{}
	t, err := os.Open(conf.For("blog.template"))

	if err != nil {
		log.Fatalf("Could not load main template! %T", err)
	}

	buf := bytes.NewBuffer(nil)
	io.Copy(buf, t)

	return string(buf.Bytes())
}

func main() {
	http.Handle("/", new(Handler))
	http.ListenAndServe(":8080", nil)
}
