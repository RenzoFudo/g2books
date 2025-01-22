package models

type User struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Pass  string `json:"pass" validate:"required"`
}
type Book struct {
	Lable  string `json:"lable" validate:"required"`
	Author string `json:"author" validate:"required"`
}
