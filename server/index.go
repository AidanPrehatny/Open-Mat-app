package main

import (
	"log"
	"net/http"
	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// SELECT * FROM openmatdata;
}

func main() {
	r := chi.NewRouter()
	r.Get("/events", handler)
	log.Fatal(http.ListenAndServe(":8080", r))
}
