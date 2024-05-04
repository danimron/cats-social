package web

type UserResponse struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Token string `json:"accessToken"`
}
