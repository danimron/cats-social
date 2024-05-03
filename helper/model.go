package helper

import (
	"cats_social/model/domain"
	"cats_social/model/web"
)

func ToCategoryResponseCat(cat domain.Cat) web.CatCreateResponse {
	return web.CatCreateResponse{
		CatId:     cat.Id,
		CreatedAt: cat.CreatedAt,
	}
}

func ToCategoryResponseUser(user domain.User) web.UserResponse {
	return web.UserResponse{
		Email: user.Email,
		Name:  user.Name,
		// Token: user.Token,
	}
}

// func ToCategoryResponsesUser(students []domain.Student) []web.StudentResponse {
// 	var userResponse []web.StudentResponse
// 	for _, student := range students {
// 		userResponse = append(userResponse, ToCategoryResponseUser(student))
// 	}
// 	return userResponse
// }
// func ToCategoryResponseBook(book domain.Book) web.BookResponse {
// 	return web.BookResponse{
// 		BookId:    book.BookId,
// 		Title:     book.Title,
// 		Available: book.Available,
// 	}
// }
// func ToCategoryResponsesBook(books []domain.Book) []web.BookResponse {
// 	var bookResponse []web.BookResponse
// 	for _, book := range books {
// 		bookResponse = append(bookResponse, ToCategoryResponseBook(book))
// 	}
// 	return bookResponse
// }
