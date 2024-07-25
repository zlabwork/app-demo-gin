package utils

import (
	"app/internal/repo/redis"
	"app/pkg"
	"context"
	"fmt"
	"math/rand"
	"time"
)

func CreateUID() int {

	opt := pkg.NewOptimus(1580030173, 59260789, 3141592653, 31)
	seq := redis.GetSeq(context.Background(), "_seq_uid")
	n := opt.Encode(int(seq))
	if n < 100000 {
		return CreateUID()
	}
	return n
}

func CreateGroupID() int {
	opt := pkg.NewOptimus(1580030173, 59260789, 3141592, 20)
	seq := redis.GetSeq(context.Background(), "_seq_gid")
	n := opt.Encode(int(seq))
	if n < 100000 {
		return CreateGroupID()
	}
	return n
}

func CreateBusinessID() int {
	opt := pkg.NewOptimus(1580030173, 59260789, 31415926, 24)
	seq := redis.GetSeq(context.Background(), "_seq_bid")
	n := opt.Encode(int(seq))
	if n < 100000 {
		return CreateBusinessID()
	}
	return n
}

func CreateOrderID(seq int) string {

	// random part
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	random := r.Intn(10000)

	// date part
	now := time.Now()
	formattedTime := now.Format("0601021504")

	// number part
	number := seq % 10000

	return formattedTime + fmt.Sprintf("%04d%04d", number, random)
}
