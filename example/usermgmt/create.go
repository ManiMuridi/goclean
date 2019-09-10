package main

import (
	"github.com/ManiMuridi/goclean/command"
	"github.com/ManiMuridi/goclean/example/model"
)

type Create struct {
	Request CreateRequest
}

type CreateRequest struct {
	User model.User
}

func (c *Create) Execute() *command.Result {

	if err := Db.Store(c.Request.User); err != nil {
		return &command.Result{
			Error: err,
			Data:  nil,
		}
	}

	return &command.Result{Error: nil, Data: c.Request.User}
}
