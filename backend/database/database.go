package database

import "github.com/go-redis/redis/v8"

type Database struct {
	client *redis.Client
}
