package help

import (
	"app/internal/bootstrap"
	"app/internal/repo/cache"
	redisRepo "app/internal/repo/redis"
	"github.com/bwmarrin/snowflake"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
)

var (
	Dir    = bootstrap.Dir
	Env    = bootstrap.Env
	Config = bootstrap.Config
	Libs   *libraries
	Token  *tokenHelp
	Cache  *redis.Client
	Redis  *redis.Client
)

type libraries struct {
	Snow *snowflake.Node
}

func newLibs() *libraries {
	snowflake.Epoch = 1498612200000 // 2017-06-28 09:10:00
	snowflake.NodeBits = 8
	snowflake.StepBits = 14
	sn, _ := snowflake.NewNode(Config.Snowflake.Node)
	return &libraries{
		Snow: sn,
	}
}

func init() {
	var err error

	Libs = newLibs()
	Token = newTokenHelp()

	host := os.Getenv("CACHE_HOST")
	port := os.Getenv("CACHE_PORT")
	Cache, err = cache.GetHandle(host, port, 0)
	if err != nil {
		log.Fatalln(err.Error())
	}

	host = os.Getenv("REDIS_HOST")
	port = os.Getenv("REDIS_PORT")
	Redis, err = redisRepo.GetHandle(host, port, 0)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
