package controller

import (
	"cats_social/helper"
	"cats_social/model/web"
	"cats_social/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CatControllerImpl struct {
	CatService service.CatService
}

func NewCatController(catService service.CatService) CatController {
	return &CatControllerImpl{
		CatService: catService,
	}
}

func (controller *CatControllerImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	catCreateRequest := web.CatCreateRequest{}
	helper.ReadFromRequestBody(r, &catCreateRequest)
	catResponse := controller.CatService.Create(r.Context(), catCreateRequest)
	webResponse := web.WebResponse{
		Message: "a",
		Data:    catResponse,
	}
	helper.WriteToResponseBody(w, webResponse)
}
