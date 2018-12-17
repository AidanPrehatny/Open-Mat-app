package main

import ( "fmt"
         "os"
         "encoding/json"
         "io/ioutil"
	 "strconv"
       )

func main() {
    // open json file
    jsonFile, err := os.Open("../mat-data.json")

    if err != nil {
        fmt.Println(err)
    }

    fmt.Println("Successfully Opened mat-data.json")

    // defer the closing of our json file so we can parse it later
    defer jsonFile.Close()

    // read xml as a bye array
    byteValue, _ := ioutil.ReadAll(jsonFile)


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
    json.Unmarshal(byteValue, &gyms)

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
