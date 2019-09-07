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
	log.Println("Starting Server on:", *port)
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
		BarItems      []Pair
		ScatterItems  []Pair
		BubbleItems   []Pair
		WeakLineItems []Pair
		O2mCircleItems []Pair
		M2mChordItems []Pair
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

	weakLine, err := addWeakLine([]Pair{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	o2mCircle, err := addO2mCircle([]Pair{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	m2mChords , err := addm2mChord([]Pair{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	data.BarItems = bar
	data.ScatterItems = scatter
	data.BubbleItems = bubble
	data.WeakLineItems = weakLine
	data.O2mCircleItems = o2mCircle
	data.M2mChordItems = m2mChords

	t.Execute(w, &data)
}

func addBar(items []Pair) ([]Pair, error) {

	// get all bar matches and add to items
	log.Println("getting bar chart matches from", "http://178.62.59.88:31195/mondial/basic/bar")
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

		link := fmt.Sprintf("http://178.62.59.88:31364/basic2/bar?e=%s&x=%s&label=%s", entity, col, key)
		label := fmt.Sprintf("%s on %s by %s", entity, col, key)

		items = append(items, Pair{label, link})

	}
	return items, nil
}

func addScatter(items []Pair) ([]Pair, error) {

	// get all bar matches and add to items
	log.Println("getting scatter chart matches from", "http://178.62.59.88:31195/mondial/basic/scatter")
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
	log.Println("getting bubble chart matches from", "http://178.62.59.88:31195/mondial/basic/bubble")
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

func addWeakLine(items []Pair) ([]Pair, error) {

	// get all bar matches and add to items
	log.Println("getting weak entity line chart matches from", "http://178.62.59.88:31195/mondial/weak/line")
	resp, err := http.Get("http://178.62.59.88:31195/mondial/weak/line")
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
		strong := path.Base(row[1])
		weak := path.Base(row[2])
		//if strong == "year" { //hack for now.. need a better way of working out which are the strong and weak keys in the pattern match
		//	strong, weak = weak, strong
		//}
		measure := path.Base(row[3])

		link := fmt.Sprintf("http://178.62.59.88:31364/weak/line?e=%s&strong=%s&weak=%s&n=%s", entity, strong, weak, measure)
		label := fmt.Sprintf("%s on %s by %s,%s", entity, measure, strong, weak)

		items = append(items, Pair{label, link})

	}
	return items, nil
}

func addO2mCircle(items []Pair) ([]Pair, error) {

	// http://localhost:8080/o2m/circle?relation=airport&many=province&one=iata_code

	log.Println("getting one 2 many circle packing matches from", "http://178.62.59.88:31195/mondial/o2m/circle")
	resp, err := http.Get("http://178.62.59.88:31195/mondial/o2m/circle")
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
		one := path.Base(row[1])
		many := path.Base(row[2])
		//measure := path.Base(row[3])
		//http://localhost:8080/o2m/circle?relation=airport&many=province&one=iata_code
		link := fmt.Sprintf("http://178.62.59.88:31364/o2m/circle?relation=%s&one=%s&many=%s", entity, one, many)
		label := fmt.Sprintf("%s by %s,%s", entity, one, many)
		items = append(items, Pair{label, link})

	}
	return items, nil
}
//"/mondial/m2m/chord",
func addm2mChord(items []Pair) ([]Pair, error) {
	log.Println("getting many to many chord matches from", "http://178.62.59.88:31195/mondial/m2m/chord")
	resp, err := http.Get("http://178.62.59.88:31195/mondial/m2m/chord")
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
		k1 := path.Base(row[1])
		k2 := path.Base(row[2])
		measure := path.Base(row[3])
		link := fmt.Sprintf("http://178.62.59.88:31364/m2m/chord?e=%s&x=%s&y=%s&n=%s", entity, k1, k2,measure)
		label := fmt.Sprintf("%s on %s by %s,%s", entity, measure, k1, k2)
		items = append(items, Pair{label, link})

	}
	return items, nil
}