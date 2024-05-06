package help

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

var (
	Dir = &directory{}
	Env = &environment{}
)

type environment struct {
	IsLocal bool
	IsDev   bool
	IsProd  bool
}

type directory struct {
	Root   string
	Config string
	Data   string
	Logs   string
	Public string
}

func init() {

	// 1. setting
	// dir
	Dir.Root = "./"
	Dir.Config = Dir.Root + "config/"
	Dir.Data = Dir.Root + "storage/data/"
	Dir.Logs = Dir.Root + "storage/logs/"
	Dir.Public = Dir.Root + "public/"

	// env
	if os.Getenv("APP_ENV") == "prod" {
		Env.IsProd = true
	} else if os.Getenv("APP_ENV") == "dev" {
		Env.IsDev = true
	} else if os.Getenv("APP_ENV") == "local" {
		Env.IsLocal = true
	} else {
		Env.IsProd = true
	}

	// 2. config
	bs, err := os.ReadFile(Dir.Config + "app.yaml")
	if err != nil {
		log.Fatal(err)
	}
	if yaml.Unmarshal(bs, Config) != nil {
		log.Fatal(err)
	}
}
