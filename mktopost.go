package main

import (
	"net/http"
	"strings"
	"log"
	"bytes"
	"github.com/pupostd/mktopost/model"
	"os"
	"io"
	"io/ioutil"
	"text/template"
)

func posts(w http.ResponseWriter, r *http.Request)  {
	path := r.URL.Path[1:]
	log.Println(path)

	s := model.Searcher{}
	data := s.LookFor(parsePath(path))

	if data != "" {
		w.Write([]byte(wrapData(data)))
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(wrapData("404 Bro ... " + http.StatusText(http.StatusNotFound))))
	}
}

func resources(w http.ResponseWriter, r *http.Request)  {
	path := r.URL.Path[1:]
	bf, err := ioutil.ReadFile("../resource/" + path)
	if err != nil {
		log.Printf("Problems loading resource %s", path)
	}

	if bf != nil {
		w.Header().Set("Content-Type", "text/css")
		w.Write([]byte(bf))
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(wrapData("404 Bro ... " + http.StatusText(http.StatusNotFound))))
	}
}

func parsePath(path string) string {
	var name = path
	st := strings.Split(path, ".")
	if len(st) > 1 {
		name = st[0]
	}
	return name
}

func wrapData(data string) string {
	t, err := template.New("main").Parse(loadTemplate())

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

func loadTemplate() string {
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
	http.HandleFunc("/", posts)
	http.HandleFunc("/res/", resources)
	http.ListenAndServe(":8181", nil)
}
