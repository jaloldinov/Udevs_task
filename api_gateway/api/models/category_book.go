package models

type CreateCategory struct {
	Name string `json:"name" binding:"required"`
}

type Category struct {
	ID           string `json:"id"`
	CategoryName string `json:"category_name" binding:"required"`
}

type GetAllCategoryResponse struct {
	Categories []Category `json:"categories"`
	Count      int32      `json:"count"`
}
