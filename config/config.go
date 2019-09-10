package config

import (
	"github.com/BurntSushi/toml"
)

var (
	Broker  *broker
	Http    *http
	Service *service
	Db      *database
	cfg     *config
)

//func init() {
//	//if err := Load("./config.toml"); err != nil {
//	//
//	//	fmt.Println(err)
//	//}
//
//}

func Load(path string) error {
	// TODO handle error safely
	if _, err := toml.DecodeFile(path, &cfg); err != nil {
		return err
	}

	Broker = &cfg.Broker
	Http = &cfg.Http
	Service = &cfg.Service
	Db = &cfg.Db

	return nil
}

type config struct {
	Broker  broker
	Http    http
	Service service
	Db      database
}

type broker struct {
	Url string
}

type http struct {
	Port uint16
}

type service struct {
	Name string
}

type database struct {
	Host     string
	Port     string
	Database string
	User     string
	Password string
}
