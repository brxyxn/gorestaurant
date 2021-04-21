package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/jackc/pgx/v4"
)

func ConnPG() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var greeting string
	err = conn.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
}

// Serve HomePage output
func homePage(w http.ResponseWriter, r *http.Request) {
	//Output as JSON
	fmt.Fprintf(w, "Welcome to homePage!")
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
