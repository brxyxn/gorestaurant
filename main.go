package main

import (
	"fmt"
	"net/http"

	_ "github.com/jackc/pgx/v4/stdlib" // This library will work along with database/sql
	dbConnection "gorestaurant.gt/dbconnection"
)

// Serve HomePage output
func homePage(w http.ResponseWriter, r *http.Request) {
	//Output as JSON
	fmt.Fprintf(w, "Welcome to homePage!")
	dbConnection.ConnectDB()
}

// Handle Requests
func handleRequest() {
	// localhost/
	http.HandleFunc("/home", homePage)
	http.ListenAndServe(":3005", nil)
}

func main() {
	handleRequest()
}
