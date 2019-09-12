package main

import (
	"fmt"
	"path/filepath"

	"github.com/ManiMuridi/goclean/config"

	"github.com/ManiMuridi/goclean/service/brokerservice"
)

func main() {
	path, _ := filepath.Abs("../")
	config.SetPath(path)
	config.SetType("toml")
	config.Load()

	brokerservice.SetContentType("application/json")

	payload := struct {
		Name string
		Age  int
	}{"John", 43}

	for i := 0; i < 10; i++ {
		msg, err := brokerservice.Push("testing-topic-exchange", "testing.created", payload)

		if err != nil {
			panic(err)
		}

		fmt.Println(msg)
	}
}
