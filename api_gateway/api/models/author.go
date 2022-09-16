package models

type CreateAuthor struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
}

type AuthorId struct {
	ID string `json:"id"`
}

type Author struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
}

type GetAllAuthorResponse struct {
	Authors []Author `json:"authors"`
	Count   int32    `json:"count"`
}
