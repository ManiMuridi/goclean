package main

import (
	"path/filepath"

	"github.com/ManiMuridi/goclean/translator"

	"github.com/ManiMuridi/goclean/config"
	"github.com/ManiMuridi/goclean/service"
)

func main() {
	path, _ := filepath.Abs("./config.toml")

	if err := config.Load(path); err != nil {
		panic(err)
	}

	translator.Enable()

	httpSvc := service.NewHttp(&handler{})

	httpSvc.Run()
}
