package main

import (
	"github.com/ManiMuridi/goclean/command"
	"github.com/ManiMuridi/goclean/example/http_service/model"
)

type Update struct {
	UserName string
	Name     string
	Email    string
	//Request *UpdateByNameRequest
}

//
//type UpdateByNameRequest struct {
//	User model.User
//}

func (u *Update) Execute() *command.Result {
	user, err := Db.FindByName(u.UserName)
	user = &model.User{
		Name:  u.Name,
		Email: u.Email,
	}

	return &command.Result{Error: err, Data: user}
}
