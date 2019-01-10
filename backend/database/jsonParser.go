package database

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "openmatws"
)

type Gyms []struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	DaysHours string  `json:"days/hours"`
	Street    string  `json:"street"`
	Phone     string  `json:"phone"`
	Lat       float64 `json:"lat"`
	Long      float64 `json:"long"`
}

func ConnectDatabase() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	return db
}

func insertIntoDB(db *sql.DB, gyms Gyms) {
	for i := 0; i < len(gyms); i++ {
		fmt.Println("ID: " + gyms[i].ID)
		fmt.Println("Name: " + gyms[i].Name)
		fmt.Println("DaysHours: " + gyms[i].DaysHours)
		fmt.Println("Street: " + gyms[i].Street)
		fmt.Println("Phone: " + gyms[i].Phone)
		fmt.Println("Lat: " + strconv.FormatFloat(gyms[i].Lat, 'f', 6, 64))
		fmt.Println("Long: " + strconv.FormatFloat(gyms[i].Long, 'f', 6, 64))

		sqlStatement := `
			INSERT INTO openmatdata (id, name, days_hours, street, phone, latitude, longitude)
			VALUES($1, $2, $3, $4, $5, $6, $7)
		`
		db.Exec(sqlStatement, gyms[i].ID, gyms[i].Name, gyms[i].DaysHours, gyms[i].Street, gyms[i].Phone, gyms[i].Lat, gyms[i].Long)
	}
}
