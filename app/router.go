package app

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func NewRouter() *httprouter.Router {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		printText("Test Route")
	})
	router.POST("/v1/user/register", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		printText("Register Route")
	})
	router.POST("/v1/user/login", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		printText("Login Route")
	})
	router.POST("/v1/cat", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		printText("Create Cat Route")
	})
	router.GET("/v1/cat", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		printText("Get Cat Route")
	})
	router.PUT("/v1/cat/:id", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		printText("Update Cat Route")
	})
	router.DELETE("/v1/cat/:id", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		printText("Delete Cat Route")
	})
	router.POST("/v1/cat/match", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		printText("Match Cat Route")
	})
	router.GET("/v1/cat/match", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		printText("Get Cat Match Route")
	})
	router.POST("/v1/cat/match/approve", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		printText("Approve Cat Match Route")
	})
	router.POST("/v1/cat/match/reject", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		printText("Reject Cat Match Route")
	})
	// router.DELETE("/v1/cat/match/:id", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// 	printText("Delete Cat Match Route")
	// })

	return router
}

func printText(text string) {
	fmt.Println(text)
}
