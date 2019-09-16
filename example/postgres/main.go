package main

import (
	"fmt"

	"github.com/ManiMuridi/goclean/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Employee struct {
	Id       string `db:"id"`
	FullName string `db:"full_name"`
}

var employees []Employee

func main() {
	config.Load()

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.GetString("database.host"),
		config.GetString("database.port"),
		config.GetString("database.username"),
		config.GetString("database.password"),
		"hrm",
	)

	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	err = db.Select(&employees, `SELECT id, full_name FROM "I180814-04-edm".employees`)
	if err != nil {
		panic(err)
	}

	fmt.Println(employees)
}
