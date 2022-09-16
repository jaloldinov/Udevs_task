package models

type CreateBookModel struct {
	Name       string `json:"name"`
	AuthorId   string `json:"author_id"`
	CategoryId string `json:"category_id"`
}

type BookModel struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	AuthorId   string `json:"author_id"`
	CategoryId string `json:"category_id"`
}

type GetAllBookModel struct {
	Books []BookModel `json:"books"`
	Count int32       `json:"count"`
}
