package main

import (
  "database/sql"
  "fmt"
  "os"
  "encoding/json"
  "io/ioutil"
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

func main() {

  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)

    db, err := sql.Open("postgres", psqlInfo)

    if err != nil {
        panic(err)
    }

    defer db.Close()

    err = db.Ping()

    if err != nil {
      panic(err)
    }

    fmt.Println("Successfully connected!")

    // open json file
    jsonFile, err := os.Open("../mat-data.json")

    if err != nil {
        fmt.Println(err)
    }

    fmt.Println("Successfully Opened mat-data.json")

    // defer the closing of our json file so we can parse it later
    defer jsonFile.Close()

    // read json as a byte array
    byteArray, _ := ioutil.ReadAll(jsonFile)


    type Gyms[]struct {
        ID        string  `json:"id"`
        Name      string  `json:"name"`
        DaysHours string  `json:"days/hours"`
        Street    string  `json:"street"`
        Phone     string  `json:"phone"`
        Lat       float64 `json:"lat"`
        Long      float64 `json:"long"`
    }

    // we initialize are gyms array
    var gyms Gyms

     // we unmarshal our byteArray which contains
    // our jsonFile's content into 'gyms' which is defined above
    json.Unmarshal(byteArray, &gyms)

    for i := 0; i < len(gyms); i++ {
        fmt.Println("ID: " + gyms[i].ID)
        fmt.Println("Name: " + gyms[i].Name)
        fmt.Println("DaysHours: " + gyms[i].DaysHours)
        fmt.Println("Street: " + gyms[i].Street)
        fmt.Println("Phone: " + gyms[i].Phone)
        fmt.Println("Lat: " + strconv.FormatFloat(gyms[i].Lat, 'f', 6, 64))
        fmt.Println("Long: " + strconv.FormatFloat(gyms[i].Long, 'f', 6, 64))
    }

}
