package main

import "github.com/ManiMuridi/goclean/command"

type GetByName struct {
	Name string
}

func (g *GetByName) Execute() *command.Result {
	user, err := Db.FindByName(g.Name)
	return &command.Result{Error: err, Data: user}
}
