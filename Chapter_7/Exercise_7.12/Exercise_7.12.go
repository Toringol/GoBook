/*
* Измените обработчик /list так, чтобы его вывод представлял
* собой таблицу HTML, а не текст. Вам может пригодиться пакет
* html/template (раздел 4.6).
 */

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var itemsTable = template.Must(template.New("itemsTable").Parse(`
<!DOCTYPE html>
<html lang="en">
  <head>
	  <meta charset="utf-8">
		<style media="screen" type="text/css">
		  table {
				border-collapse: collapse;
				border-spacing: 0px;
			}
		  table, th, td {
				padding: 5px;
				border: 1px solid black;
			}
		</style>
	</head>
	<body>
		<h1>Items</h1>
		<table>
			<tbody>
				{{range $key, $value := .Data}}
				<tr>
					<td>{{$key}}</td>
					<td>{{$value}}</td>
				</tr>
				{{end}}
			</tbody>
		</table>
	</body>
</html>
`))

func main() {
	data := map[string]dollars{
		"shoes": 50,
		"socks": 5,
	}
	mu := &sync.Mutex{}

	db := database{data, mu}

	http.HandleFunc("/create", db.create)
	http.HandleFunc("/read", db.read)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database struct {
	Data map[string]dollars
	mu   *sync.Mutex
}

func (db database) list(w http.ResponseWriter, req *http.Request) {
	if err := itemsTable.Execute(w, db); err != nil {
		log.Fatal(err)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db.Data[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	items, ok := req.URL.Query()["item"]
	if !ok || len(items[0]) < 1 {
		fmt.Fprintf(w, "Error Response")
	}
	prices, ok := req.URL.Query()["price"]
	if !ok || len(prices[0]) < 1 {
		fmt.Fprintf(w, "Error Response")
	}

	item := items[0]
	price, err := strconv.ParseFloat(prices[0], 32)
	if err != nil {
		fmt.Fprintf(w, "Error Price")
	}

	db.mu.Lock()
	db.Data[item] = dollars(price)
	db.mu.Unlock()
}

func (db database) read(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db.Data[item]; ok {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	items, ok := req.URL.Query()["item"]
	if !ok || len(items[0]) < 1 {
		fmt.Fprintf(w, "Error Response")
	}
	prices, ok := req.URL.Query()["price"]
	if !ok || len(prices[0]) < 1 {
		fmt.Fprintf(w, "Error Response")
	}

	item := items[0]
	price, err := strconv.ParseFloat(prices[0], 32)
	if err != nil {
		fmt.Fprintf(w, "Error Price")
	}

	if _, ok := db.Data[item]; !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}

	db.mu.Lock()
	db.Data[item] = dollars(price)
	db.mu.Unlock()
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if _, ok := db.Data[item]; ok {
		delete(db.Data, item)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}
