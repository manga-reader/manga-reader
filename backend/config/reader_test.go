package config

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_LoadConfiguration(t *testing.T) {
	testCfg := &Configuration{
		Connection: Connection{
			ExportHost: "/fake/host",
			ExportPort: "6699",
			Timeout:    Duration{time.Duration(2 * time.Minute)},
			JWTSecret:  "this tis test secret",
		},
		Redis: Redis{
			ServerAddr: "/redis/host:123",
			Password:   DefaultRedisPassword,
			DBIndex:    DefaultRedisDBIndex,
		},
	}

	err := LoadConfiguration("./reader_test.json")
	assert.NoError(t, err)
	assert.Equal(t, testCfg, Cfg)
}
