package controller

import (
	"cats_social/helper"
	"cats_social/model/web"
	"cats_social/service"
	"net/http"
	"strconv"

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

func (controller *CatControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	CatGetParam := web.CatGetParam{}
	helper.NewGetCatParam(r, &CatGetParam)

	catResponse := controller.CatService.FindAll(r.Context(), &CatGetParam)
	webResponse := web.WebResponse{
		Message: "a",
		Data:    catResponse,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CatControllerImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	catId := p.ByName("cat_id")
	id, err := strconv.Atoi(catId)
	helper.PanicIfError(err)

	catCreateRequest := web.CatCreateRequest{}
	catCreateRequest.Id = id

	helper.ReadFromRequestBody(r, &catCreateRequest)
	controller.CatService.Update(r.Context(), catCreateRequest)
	webResponse := web.WebResponse{
		Message: "a",
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (controller *CatControllerImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	catId := p.ByName("cat_id")
	id, err := strconv.Atoi(catId)
	helper.PanicIfError(err)

	controller.CatService.Delete(r.Context(), id)

	webResponse := web.WebResponse{
		Message: "a",
	}
	helper.WriteToResponseBody(w, webResponse)
}

// book_id := r.URL.Query().Get("book_id")
