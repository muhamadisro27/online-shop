package main

import (
	"fmt"
	"net/http"
	database "online-shop/db"
	"online-shop/helper"
	"os"
)

func main() {
	APP_PORT := os.Getenv("APP_PORT")

	db := database.OpenConnection()

	database.MigrateTable(db)

	defer db.Close()

	server := &http.Server{
		Addr:    ":" + APP_PORT,
		Handler: nil,
	}

	err := server.ListenAndServe()

	helper.PanicIfError(fmt.Sprintf("failed to running server on port %s :", APP_PORT), err)
}
