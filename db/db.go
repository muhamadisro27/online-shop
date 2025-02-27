package db

import (
	"database/sql"
	"fmt"
	"online-shop/helper"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func OpenConnection() *sql.DB {

	db, err := sql.Open("pgx", os.Getenv("DB_CONNECTION_URI"))
	if err != nil {
		helper.PanicIfError(err)
	}

	if err = db.Ping(); err != nil {
		helper.PanicIfError(err)
	}

	fmt.Println("success connect DB !")

	return db
}
