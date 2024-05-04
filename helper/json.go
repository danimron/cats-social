package helper

import (
	"cats_social/model/web"

	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(r *http.Request, result interface{}) {
	decoder := json.NewDecoder(r.Body) // untuk decode json menjadi struct
	err := decoder.Decode(result)
	PanicIfError(err)
}

func WriteToResponseBody(w http.ResponseWriter, response interface{}) {
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w) // untuk encode kembali menjadi json
	err := encoder.Encode(response)
	PanicIfError(err)
}

func NewGetCatParam(r *http.Request, result *web.CatGetParam) {
	result.Id = r.URL.Query().Get("id")
	result.Limit = r.URL.Query().Get("limit")
	result.Offset = r.URL.Query().Get("offset")
	result.Race = r.URL.Query().Get("race")
	result.Sex = r.URL.Query().Get("sex")
	result.HasMatched = r.URL.Query().Get("hasMatched")
	result.AgeInMonth = r.URL.Query().Get("ageInMonth")
	result.Owned = r.URL.Query().Get("owned")
	result.Search = r.URL.Query().Get("search")
}
