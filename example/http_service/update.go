package main

import (
	"github.com/ManiMuridi/goclean/command"
	"github.com/ManiMuridi/goclean/example/http_service/model"
)

type Update struct {
	Request *UpdateByNameRequest
}

type UpdateByNameRequest struct {
	Name string
	User model.User
}

func (u *Update) Execute() *command.Result {
	user, err := Db.FindByName(u.Request.Name)
	user = &u.Request.User

	return &command.Result{Error: err, Data: user}
}
