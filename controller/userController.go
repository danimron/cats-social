package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserController interface {
	Register(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	Login(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	// Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	// Update(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	// FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	// FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}
