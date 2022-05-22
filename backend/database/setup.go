package database

import "github.com/go-redis/redis/v8"

func Connect() *Database {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return &Database{
		client: client,
	}
}
