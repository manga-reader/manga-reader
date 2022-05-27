package database

import "github.com/go-redis/redis/v8"

func Connect(serverAddr, password string, dbIdx int) *Database {
	client := redis.NewClient(&redis.Options{
		Addr:     serverAddr,
		Password: password, // no password set
		DB:       dbIdx,    // use default DB
	})

	return &Database{
		client: client,
	}
}
