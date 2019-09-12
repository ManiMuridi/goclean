package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var (
	configLoaded = false
	confType     = "toml"
	confPath     = "."
)

func Load() {
	configLoaded = true
	viper.SetConfigName("config")
	viper.AddConfigPath(confPath)
	viper.SetConfigType(confType)

	if err := viper.ReadInConfig(); err != nil {
		if err, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(err)
		} else {
			panic(err)
		}
	}
}

func SetPath(path string) {
	checkConfigLoaded()
	confPath = path
}

func SetType(configType string) {
	checkConfigLoaded()
	confType = configType
}

func Get(key string) interface{} {
	checkConfigLoaded()
	return viper.Get(key)
}

func GetString(key string) string {
	checkConfigLoaded()
	return viper.GetString(key)
}

func GetBool(key string) bool {
	checkConfigLoaded()
	return viper.GetBool(key)
}

func GetInt(key string) int {
	checkConfigLoaded()
	return viper.GetInt(key)
}

func GetInt32(key string) int32 {
	checkConfigLoaded()
	return viper.GetInt32(key)
}

func GetInt64(key string) int64 {
	checkConfigLoaded()
	return viper.GetInt64(key)
}

func checkConfigLoaded() {
	if !configLoaded {
		fmt.Println("Config has not been loaded. Did you forget to call config.Load()?")
	}
}
