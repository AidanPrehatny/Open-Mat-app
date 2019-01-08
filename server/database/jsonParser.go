package populate

import (
	"database/sql"
	"fmt"
	"os"
	"encoding/json"
	"io/ioutil"
	"strconv"
	_ "github.com/lib/pq"
	"server/foo"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "openmatws"
)

type GymSummary[]struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	DaysHours string  `json:"days/hours"`
	Street    string  `json:"street"`
	Phone     string  `json:"phone"`
	Lat       float64 `json:"lat"`
	Long      float64 `json:"long"`
}

func connectDatabase() *sql.DB {
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

func insertIntoDB(db *sql.DB, gyms GymSummary) {
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

func main() {
	var db *sql.DB
	db = connectDatabase()

	// defer db.Close()
	// open json file
	jsonFile, err := os.Open("mat-data.json")

	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println("Successfully Opened mat-data.json")

	// defer the closing of our json file so we can parse it later
	defer jsonFile.Close()

	// read json as a byte array
	byteArray, _ := ioutil.ReadAll(jsonFile)

	fmt.Println(byteArray)

	// we initialize are gyms array
	var gyms Gyms

	// we unmarshal our byteArray which contains
	// our jsonFile's content into 'gyms' which is defined above
	json.Unmarshal(byteArray, &gyms);
	
	insertIntoDB(db, gyms);
	webServer(db)
}
