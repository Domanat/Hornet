package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	password, err := ioutil.ReadFile("dbConfig.txt")
	if err != nil {
		panic(err)
	}

	dataSourceName := fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/hornet", password)

	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", logup)
	http.HandleFunc("/thanks", thanks)
	http.HandleFunc("/login", login)

	fmt.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}
