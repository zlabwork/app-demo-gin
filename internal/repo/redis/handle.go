package redis

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
)

var conn *redis.Client

func getHandle() (*redis.Client, error) {

	if conn != nil {
		return conn, nil
	}

	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	name := "1"
	dsn := fmt.Sprintf("redis://%s:%s/%s", host, port, name)
	return connectRedis(dsn)
}

// ConnectRedis
// By default, the pool size is 10 connections per every available CPU as reported by runtime.GOMAXPROCS
// redis://<user>:<pass>@localhost:6379/<db>
// https://redis.uptrace.dev/guide/server.html#connecting-to-redis-server
// https://redis.uptrace.dev/guide/go-redis-debugging.html#connection-pool-size
func connectRedis(dsn string) (*redis.Client, error) {

	opt, err := redis.ParseURL(dsn)
	if err != nil {
		return nil, err
	}
	cli := redis.NewClient(opt)
	return cli, nil
}

// Cluster
// https://redis.uptrace.dev/guide/go-redis-cluster.html
//cli := redis.NewClusterClient(&redis.ClusterOptions{
//	Addrs: []string{":7000", ":7001", ":7002"},
//})

// Cache
// https://redis.uptrace.dev/guide/go-redis-cache.html
