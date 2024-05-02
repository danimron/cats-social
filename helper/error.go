package helper

import (
	"cats_social/model/web"
	"net/http"
)

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
func Unauthorized(w http.ResponseWriter) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	webResponse := web.WebResponse{
		Message: "a",
	}
	WriteToResponseBody(w, webResponse)
}
