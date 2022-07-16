package database

import (
	"context"
	"time"
)

func (d *Database) Set(key, val string) error {
	return d.client.Set(context.Background(), key, val, 10*time.Second).Err()
}

func (d *Database) Get(key string) (string, error) {
	return d.client.Get(context.Background(), key).Result()
}

func (d *Database) Del(key string) error {
	return d.client.Del(context.Background(), key).Err()
}

func (d *Database) ListRangeAll(key string) ([]string, error) {
	return d.client.LRange(context.Background(), key, 0, -1).Result()
}

func (d *Database) ListIndex(key string, index int64) (string, error) {
	return d.client.LIndex(context.Background(), key, index).Result()
}

// push val to the beginning of the list
func (d *Database) ListPush(key string, val []string) error {
	return d.client.LPush(context.Background(), key, val).Err()
}

func (d *Database) ListRemoveElement(key string, val string) error {
	return d.client.LRem(context.Background(), key, 0, val).Err()
}
