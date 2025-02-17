package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Product struct {
	ID       uint
	Nama     uint
	Kategori string
	Harga    int
}

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

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS produk (
		id SERIAL PRIMARY KEY,
		nama VARCHAR(255),
		kategori VARCHAR(50),
		harga INT);
	`)

	if err != nil {
		fmt.Printf("Failed to create table : %v", err)
		os.Exit(1)
	}

	fmt.Println("success create table")
}
