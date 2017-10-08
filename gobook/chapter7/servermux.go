package main

import (
	"fmt"
	"log"
	"net/http"
)

type dollars float64

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (d database) list(w http.ResponseWriter, r *http.Request) {
	for item, price := range d {
		fmt.Fprintf(w, "%s %s\n", item, price)
	}
}

func (d database) price(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price, ok := d[item]

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}

	fmt.Fprintf(w, "%s\n", price)
}

func main() {
	db := database{"socks": 50, "lens": 50}
	serverMux := http.NewServeMux()
	serverMux.HandleFunc("/list", db.list)
	serverMux.HandleFunc("/price", db.price)

	log.Fatal(http.ListenAndServe("localhost:9000", serverMux))
}
