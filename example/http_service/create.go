package main

import (
	"github.com/ManiMuridi/goclean/command"
	"github.com/ManiMuridi/goclean/example/http_service/model"
	"github.com/ManiMuridi/goclean/validation"
)

type Create struct {
	Name  string `validate:"required"`
	Email string `validate:"email"`
	//Request *CreateRequest
}

type CreateRequest struct {
	User model.User
}

func (c *Create) Validate() error {
	return validation.Validator.Validate(c)
}

func (c *Create) Execute() *command.Result {
	//if err := Db.Store(c.Request.User); err != nil {
	//	return &command.Result{
	//		Error: err,
	//		Data:  nil,
	//	}
	//}

	return &command.Result{Error: nil, Data: model.User{Name: c.Name, Email: c.Email}}
}
