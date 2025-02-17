package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	connectURI := os.Getenv("DB_CONNECTION_URI")

	db, err := sql.Open("pgx", connectURI)
	if err != nil {
		fmt.Printf("Failed connect to DB : %v\n", db)
		os.Exit(1)
	}

	defer db.Close()

	err = db.Ping()

	fmt.Println(connectURI)
	if err != nil {
		fmt.Printf("Something went wrong : %v\n", db)
		os.Exit(1)
	}

	fmt.Println("Success connect DB !")
}
