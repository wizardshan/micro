package config

import (
	"github.com/kr/pretty"
	"github.com/spf13/viper"
	"log"
	"path/filepath"
)

var Config *App

type Cache struct {
	Router map[string][]string
}

type App struct {
	Cache *Cache
	Name  string
}

func init() {

	Config = new(App)

	var loader = viper.New()
	path := "./config"
	loader.AddConfigPath(path)
	loader.SetConfigName("app")
	loader.SetConfigType("yaml")
	if err := loader.ReadInConfig(); err != nil {
		log.Println(filepath.Abs(path))
		log.Fatalf("load config file error: %s, path: %s", err.Error(), path)
	}

	if err := loader.Unmarshal(Config); err != nil {
		log.Fatalf("config unmarshal error: %s", err.Error())
	}

	pretty.Println(Config)
}
