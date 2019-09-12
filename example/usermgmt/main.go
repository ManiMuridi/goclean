package main

import (
	"path/filepath"

	"github.com/ManiMuridi/goclean/service/httpservice"

	"github.com/ManiMuridi/goclean/translator"

	"github.com/ManiMuridi/goclean/config"
)

func main() {
	path, _ := filepath.Abs("./config.toml")

	if err := config.Load(path); err != nil {
		panic(err)
	}

	translator.Enable()

	httpSvc := httpservice.NewHttp(&handler{})

	httpSvc.Run()
}
