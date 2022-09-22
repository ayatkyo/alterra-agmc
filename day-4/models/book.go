package models

type Book struct {
	ID     uint   `json:"id"`
	Title  string `json:"title" form:"title" validate:"required"`
	Author string `json:"author" form:"author" validate:"required"`
	Year   uint   `json:"year" form:"year" validate:"required,numeric"`
}
