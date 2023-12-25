package app

import (
	"github.com/bwmarrin/snowflake"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"strconv"
)

var (
	Dir  = &directory{}
	Env  = &environment{}
	Libs *libraries
)

type environment struct {
	IsDev  bool
	IsProd bool
}

type directory struct {
	Root   string
	Config string
	Data   string
}

type libraries struct {
	Snow *snowflake.Node
}

func initLibs() *libraries {
	i, _ := strconv.ParseInt(os.Getenv("APP_NODE"), 10, 64)
	snowflake.Epoch = 1498612200000 // 2017-06-28 09:10:00
	snowflake.NodeBits = 8
	snowflake.StepBits = 14
	sn, _ := snowflake.NewNode(i)
	return &libraries{
		Snow: sn,
	}
}

func Boot() {

	// 1. setting
	Dir.Root = "./"
	Dir.Config = Dir.Root + "config/"
	Dir.Data = Dir.Root + "var/"
	if os.Getenv("APP_ENV") == "dev" {
		Env.IsDev = true
	} else if os.Getenv("APP_ENV") == "prod" {
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

	// 3. libs
	Libs = initLibs()
}
