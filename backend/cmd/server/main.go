package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/AidanPrehatny/Open-Mat-app/backend/database"

	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
)

type gymSummary struct {
	ID        string
	Name      string
	DaysHours string
	Street    string
	Phone     string
	Latitude  float64
	Longitude float64
}

type gyms struct {
	Gyms []gymSummary
}

var db *sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "openmatws"
)

func queryGymz(gymz *gyms) error {
	// run our query of database table column names
	rows, err := db.Query(`
		SELECT
		id,
		name,
		days_hours,
		street,
		phone,
		latitude,
		longitude
		FROM openmatdata`)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		gym := gymSummary{}
		// scan gym rows for errors ?
		err = rows.Scan(
			&gym.ID,
			&gym.Name,
			&gym.DaysHours,
			&gym.Street,
			&gym.Phone,
			&gym.Latitude,
			&gym.Longitude,
		)
		if err != nil {
			return err
		}
		// afters query is done, next we add to gym object
		gymz.Gyms = append(gymz.Gyms, gym)
	}
	err = rows.Err()
	if err != nil {
		return err
	}
	return nil
}

func handler(w http.ResponseWriter, req *http.Request) {
	// w.Write([]byte("welcome"))
	gymz := gyms{}
	err := queryGymz(&gymz) // run gym query, gymz is our object, pass as ref
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	out, err := json.Marshal(gymz) // encode our gym query to json
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Fprintf(w, string(out)) // will print out the json onto the webpage at localhost:8080/api
}

func main() {
	db = database.ConnectDatabase()
	defer db.Close()
	r := chi.NewRouter()
	r.Get("/api", handler)
	fmt.Println("listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
