package database

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type RDConfig struct {
	Addr 		string
	Password 	string
}

var Ctx = context.Background()

func CreateClient(config *RDConfig) (*redis.Client,error) {
	rdb := redis.NewClient(
		&redis.Options{
			Addr: config.Addr,
			Password: config.Password,
			DB: 0,
		},
	)

	err := rdb.Ping(Ctx).Err()

	if err != nil {
		return nil,err
	}

	return rdb, nil
}

