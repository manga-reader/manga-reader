package config

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMarshal(t *testing.T) {
	testJson := `{
		"timeout": "1m"
	}`

	testCfg := Configuration{
		Timeout: Duration{time.Duration(1 * time.Minute)},
	}

	var cfg Configuration
	var err error
	err = json.Unmarshal([]byte(testJson), &cfg)
	assert.NoError(t, err)
	assert.Equal(t, testCfg, cfg)

	_, err = json.Marshal(cfg)
	assert.NoError(t, err)
}
