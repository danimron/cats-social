package main

import (
	"cats_social/app"
	"fmt"
	"net/http"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load(".env")

	fmt.Println("Hello, World!")

	db, err := app.ConnectToPostgres()
	if err != nil {
		fmt.Println(err)
	}

	db.Close()

	router := app.NewRouter()
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}
	err = server.ListenAndServe()

}
