package main

import (
	"errors"
	"strings"

	"github.com/ManiMuridi/goclean/example/http_service/model"
)

var (
	users = []model.User{
		{
			Name:  "Ahmed",
			Email: "ahmed@example.com",
		},
		{
			Name:  "Khalid",
			Email: "khalid@example.com",
		},
		{
			Name:  "Mahmoud",
			Email: "mahmoud@example.com",
		},
	}
	Db = NewPostgresStorage()
)

type Storage interface {
	FindAll() []model.User
	FindByName(name string) (*model.User, error)
	Store(user model.User) error
}

func NewPostgresStorage() Storage {
	return &postgres{}
}

type postgres struct{}

func (p *postgres) FindAll() []model.User {
	return users
}

func (p *postgres) FindByName(name string) (*model.User, error) {
	for i := range users {
		if strings.ToLower(users[i].Name) == strings.ToLower(name) {
			return &users[i], nil
		}
	}

	return nil, errors.New("cannot find a user with that name")
}

func (p *postgres) Store(user model.User) error {
	users = append(users, user)
	return nil
}
