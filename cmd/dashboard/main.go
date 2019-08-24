package main

import (
	"flag"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path"
)

var port = flag.String("http", ":8080", "port")

func main() {
	flag.Parse()
	http.HandleFunc("/menu.html", menuHandler)
	//	http.HandleFunc("/default.html", fileHandler)
	http.HandleFunc("/", fileHandler)
	log.Fatal(http.ListenAndServe(*port, nil))
}

type Pair struct {
	Name  string
	Value string
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
	name := path.Base(r.URL.Path)
	file, err := ioutil.ReadFile(name)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		return
	}

	w.Write(file)
}

func menuHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("menu.html")
	data := struct {
		Items []Pair
	}{
		Items: []Pair{
			{"a", "http://178.62.59.88:31364/basic/scatter"},
			{"b", "http://178.62.59.88:31364/basic/bar"}},
	}
	t.Execute(w, &data)
}
