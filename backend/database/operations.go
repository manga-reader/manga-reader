package database

import (
	"context"
	"time"
)

func (d *Database) Set(key, val string) error {
	return d.client.Set(context.Background(), key, val, 10*time.Second).Err()
}

func (d *Database) Get(key string) (string, error) {
	return d.client.Get(context.TODO(), key).Result()
}
