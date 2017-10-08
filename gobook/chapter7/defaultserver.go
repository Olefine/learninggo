package main

import (
	"fmt"
	"log"
	"net/http"
)

type key string
type value string

type kvstore map[key]value

func (s kvstore) list(w http.ResponseWriter, r *http.Request) {
	for item, value := range s {
		fmt.Fprintf(w, "%s - %s\n", item, value)
	}
}

func (s kvstore) item(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	kvalue, ok := s[key(item)]

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "value associated with key: %s was not found there\n", item)
		return
	}

	fmt.Fprintf(w, "%s\n", kvalue)

}

func main() {
	store := kvstore{"Egor": "Gorodov", "Aleksey": "Chumagin"}
	http.HandleFunc("/list", store.list)
	http.HandleFunc("/item", store.item)
	log.Println(http.ListenAndServe("localhost:9000", nil))
}
