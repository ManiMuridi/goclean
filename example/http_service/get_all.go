package main

import (
	"github.com/ManiMuridi/goclean/command"
)

type GetAll struct{}

func (g *GetAll) Execute() *command.Result {
	cr := &command.Result{
		Error: nil,
		Data:  Db.FindAll(),
	}

	return cr
}
