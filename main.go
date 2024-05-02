package main

import (
	"cats_social/app"
	"cats_social/controller"
	"cats_social/repository"
	"cats_social/service"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load(".env")
	validate := validator.New()

	fmt.Println("Hello, World!")

	db, err := app.ConnectToPostgres()
	if err != nil {
		fmt.Println(err)
	}

	//cat
	catRepository := repository.NewCatRepository()
	catService := service.NewCatService(catRepository, db, validate)
	catController := controller.NewCatController(catService)

	router := app.NewRouter(catController)
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}
	err = server.ListenAndServe()

}
