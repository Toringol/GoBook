/*
* Добавьте дополнительные обработчики так, чтобы клиент мог
* создавать, читать, обновлять и удалять записи базы данных.
* Например, запрос вида /update?item=socks&price=6 должен
* обновлять цену товара в базе данных и сообщать об ошибке,
* если товар отсутствует или цена некорректна (предупреждение:
* это изменение вносит в программу параллельное обновление переменных).
 */

package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

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
	data map[string]dollars
	mu   *sync.Mutex
}

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db.data {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db.data[item]; ok {
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
	db.data[item] = dollars(price)
	db.mu.Unlock()
}

func (db database) read(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db.data[item]; ok {
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

	if _, ok := db.data[item]; !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}

	db.mu.Lock()
	db.data[item] = dollars(price)
	db.mu.Unlock()
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if _, ok := db.data[item]; ok {
		delete(db.data, item)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}
