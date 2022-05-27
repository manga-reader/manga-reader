package config

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_LoadConfiguration(t *testing.T) {
	testCfg := &Configuration{
		ExportHost: DefaultExportHost,
		ExportPort: DefaultExportPort,
		Timeout:    Duration{time.Duration(1 * time.Minute)},
	}

	err := LoadConfiguration("./reader_test.json")
	assert.NoError(t, err)
	assert.Equal(t, testCfg, Cfg)
}
