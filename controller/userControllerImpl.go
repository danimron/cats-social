package controller

import (
	"cats_social/helper"
	"cats_social/model/web"
	"cats_social/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) Register(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userCreateRequest := web.UserRegisterRequest{}
	helper.ReadFromRequestBody(r, &userCreateRequest)
	userResponse, err := controller.UserService.Register(r.Context(), userCreateRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		webResponse := web.WebResponse{
			Message: err.Error(),
		}
		helper.WriteToResponseBody(w, webResponse)
	} else {
		w.WriteHeader(http.StatusCreated)
		webResponse := web.WebResponse{
			Message: "User registered successfully",
			Data:    userResponse,
		}
		helper.WriteToResponseBody(w, webResponse)
	}
}
