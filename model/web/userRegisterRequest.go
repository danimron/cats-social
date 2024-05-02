package web

type UserRegisterRequest struct {
	Email    string `validate:"email,required"json:"email"`
	Password string `validate:"email,required"json:"password"`
	Name     string `validate:"required"json:"name"`
}
