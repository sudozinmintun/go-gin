package dto

// CreateCategoryRequest is the input for POST /categories
type CreateCategoryRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Unit        string `json:"unit"`
}

// CategoryResponse is the output format
type CategoryResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Unit        string `json:"unit"`
}
