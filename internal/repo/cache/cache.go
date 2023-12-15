package cache

import (
	"context"
	"time"
)

func Set(ctx context.Context, key string, value []byte, expiration time.Duration) error {

	cli, err := getHandle()
	if err != nil {
		return err
	}

	cmd := cli.Set(ctx, key, value, expiration)
	return cmd.Err()
}

func Get(ctx context.Context, key string) ([]byte, error) {

	cli, err := getHandle()
	if err != nil {
		return nil, err
	}

	cmd := cli.Get(ctx, key)
	return cmd.Bytes()
}
