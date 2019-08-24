package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"html/template"
	"io"
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
		BarItems     []Pair
		ScatterItems []Pair
		BubbleItems  []Pair
	}{}

	bar, err := addBar([]Pair{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	scatter, err := addScatter([]Pair{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	bubble, err := addBubble([]Pair{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	data.BarItems = bar
	data.ScatterItems = scatter
	data.BubbleItems = bubble

	t.Execute(w, &data)
}

func addBar(items []Pair) ([]Pair, error) {

	// get all bar matches and add to items
	resp, err := http.Get("http://178.62.59.88:31195/mondial/basic/bar")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	r := csv.NewReader(resp.Body)
	r.Read() // ignore first line
	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		entity := path.Base(row[0])
		key := path.Base(row[1])
		col := path.Base(row[2])

		link := fmt.Sprintf("http://178.62.59.88:31364/basic/bar?e=%s&x=%s&label=%s", entity, col, key)
		label := fmt.Sprintf("%s on %s by %s", entity, col, key)

		items = append(items, Pair{label, link})

	}
	return items, nil
}

func addScatter(items []Pair) ([]Pair, error) {

	// get all bar matches and add to items
	resp, err := http.Get("http://178.62.59.88:31195/mondial/basic/scatter")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	r := csv.NewReader(resp.Body)
	r.Read() // ignore first line
	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		entity := path.Base(row[0])
		key := path.Base(row[1])
		a := path.Base(row[2])
		b := path.Base(row[3])

		//
		// 	entity,key,scalarA,scalarB
		link := fmt.Sprintf("http://178.62.59.88:31364/basic/scatter?e=%s&x=%s&y=%s&c=%s&label=%s", entity, a, b, a, key)
		label := fmt.Sprintf("%s on %s,%s by %s", entity, a, b, key)

		items = append(items, Pair{label, link})

	}
	return items, nil
}

func addBubble(items []Pair) ([]Pair, error) {

	// get all bar matches and add to items
	resp, err := http.Get("http://178.62.59.88:31195/mondial/basic/bubble")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	r := csv.NewReader(resp.Body)
	r.Read() // ignore first line
	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		entity := path.Base(row[0])
		key := path.Base(row[1])
		a := path.Base(row[2])
		b := path.Base(row[3])
		c := path.Base(row[4])

		link := fmt.Sprintf("http://178.62.59.88:31364/basic/bubble?e=%s&x=%s&y=%s&c=%s&s=%s&label=%s", entity, a, b, c, c, key)
		label := fmt.Sprintf("%s on %s,%s,%s by %s", entity, a, b, c, key)

		items = append(items, Pair{label, link})

	}
	return items, nil
}
