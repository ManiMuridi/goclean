package model

type User struct {
	Name  string `validate:"required"`
	Email string `validate:"email"`
}
