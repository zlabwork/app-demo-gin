package redis

import (
	"app/internal/help"
	"context"
	"fmt"
	"log"
	"os"
)

func GetSeq(ctx context.Context, key string) int64 {
	cli := help.Redis
	n := cli.Incr(ctx, key).Val()
	if n < 100 {
		log.Printf("Check redis key %s, if the system ID is repeated\n", key)
		log.Println(fmt.Sprintf("For safety, you need set this redis value [key=%s] at least 100 or greater", key))
		os.Exit(0)
	}

	return n
}
