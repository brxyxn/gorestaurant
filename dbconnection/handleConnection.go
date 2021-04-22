package dbConnection

import (
	"database/sql"
	"fmt"
	"os"

	"gorestaurant.gt/utility"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "gorestaurant_dev"
)

func connDB() int {
	// Connect to DB
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("pgx", psqlInfo)
	// conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		utility.ReturnError("Connection Error: ", err)
		os.Exit(1)
	}
	defer db.Close()

	// Check connection to the DB
	err = db.Ping()
	if err != nil {
		utility.ReturnError("Connection Error: ", err)
	}
	fmt.Println("Connected to the DB Successfully")

	err = getAllItems(db)
	if err != nil {
		utility.ReturnError("Error getting items from DB. ", err)
	}

	return 0
}

func getAllItems(conn *sql.DB) error {
	// Query string
	query := `select id, name, description from public.posts`
	rows, err := conn.Query(query)
	if err != nil {
		utility.ReturnError("Query Error. ", err)
		return err
	}
	var id int
	var name, description string
	fmt.Println(utility.CheckCount(rows))
	for rows.Next() {
		// Mapping results

		err := rows.Scan(&id, &name, &description)
		if err != nil {
			utility.ReturnError("Error mapping rows: ", err)
		}
		fmt.Printf("#%d Title: %s\n Description: %s\n\n", id, name, description)
	}
	defer rows.Close()

	return nil
}
