package web

type CatCreateRequest struct {
	Id          int    `json:"id"`
	UserId      int    `json:"user_id"`
	Name        string `json:"name"`
	Race        string `json:"race"`
	Sex         string `json:"sex"`
	AgeInMonth  int    `json:"age_in_month"`
	Description string `json:"description"`
}
