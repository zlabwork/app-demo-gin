package help

import (
	"app/internal/bootstrap"
	"app/internal/bootstrap/core"
	"app/internal/consts"
	"app/pkg"
	"context"
	"github.com/bwmarrin/snowflake"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	Dir    = bootstrap.Dir
	Env    = bootstrap.Env
	Config = bootstrap.Config
	Libs   *libraries
	Cache  *redis.Client
	Redis  *redis.Client
	Db     *gorm.DB
)

var op = pkg.NewOptimus(2123809381, 1885413229, 146808189, 31)

type libraries struct {
	Snow  *snowflake.Node
	Token *tokenHelp
}

func newLibs() *libraries {
	snowflake.Epoch = 1498612200000 // 2017-06-28 09:10:00
	snowflake.NodeBits = 8
	snowflake.StepBits = 14
	sn, _ := snowflake.NewNode(Config.Snowflake.Node)
	return &libraries{
		Snow:  sn,
		Token: newTokenHelp(),
	}
}

func NewSeqId() int {
	id := Redis.Incr(context.Background(), consts.RedisSeqIdKey).Val()
	return op.Encode(int(id))
}

func init() {
	var err error

	Libs = newLibs()

	// cache
	host := os.Getenv("CACHE_HOST")
	port := os.Getenv("CACHE_PORT")
	Cache, err = core.GetRedisHandle(host, port, 0)
	if err != nil {
		log.Fatalln(err.Error())
	}
	// redis
	host = os.Getenv("REDIS_HOST")
	port = os.Getenv("REDIS_PORT")
	Redis, err = core.GetRedisHandle(host, port, 0)
	if err != nil {
		log.Fatalln(err.Error())
	}
	// postgres
	host = os.Getenv("DB_HOST")
	port = os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")
	prefix := os.Getenv("DB_PREFIX")
	Db, err = core.GetDbHandle(host, port, user, pass, name, prefix)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
