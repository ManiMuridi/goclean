package main

import (
	"path/filepath"

	"github.com/ManiMuridi/goclean/config"
	"github.com/ManiMuridi/goclean/service/httpservice"
	"github.com/ManiMuridi/goclean/translator"
)

func main() {
	path, _ := filepath.Abs("./")
	config.SetPath(path)
	config.SetType("toml")
	config.Load()

	translator.Enable()

	httpSvc := httpservice.NewHttp(&handler{})

	httpSvc.Run()
}
