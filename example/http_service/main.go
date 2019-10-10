package main

import (
	"github.com/ManiMuridi/goclean/config"
	"github.com/ManiMuridi/goclean/service/httpservice"
)

func main() {
	//path, _ := filepath.Abs("./")
	//	//config.SetPath(path)
	//	//config.SetType("toml")
	config.Load()

	//util.EnableTranslation()

	httpSvc := httpservice.NewHttp(&handler{})

	httpSvc.Run()
}
