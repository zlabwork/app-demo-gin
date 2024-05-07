package cache

import (
	"fmt"
	"github.com/redis/go-redis/v9"
)

var conn *redis.Client

func GetHandle(host, port string, index int) (*redis.Client, error) {

	if conn != nil {
		return conn, nil
	}

	dsn := fmt.Sprintf("redis://%s:%s/%d", host, port, index)
	conn, err := connectRedis(dsn)
	return conn, err
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
