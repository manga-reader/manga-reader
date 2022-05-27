package config

import (
	"encoding/json"
)

func init() {
	Cfg = setDefaultConfig()
}

var Cfg *Configuration

type Configuration struct {
	ExportHost      string   `json:"host,omitempty"`
	ExportPort      string   `json:"port,omitempty"`
	Timeout         Duration `json:"timeout,omitempty"`
	RedisServerAddr string   `json:"redis_server_address,omitempty"`
	RedisPassword   string   `json:"redis_password,omitempty"`
	RedisDBIndex    int      `json:"redis_db_index,omitempty"`
}

func (c *Configuration) String() string {
	out, err := json.Marshal(c)
	if err != nil {
		return ""
	}
	return string(out)
}

func setDefaultConfig() *Configuration {
	config := &Configuration{
		ExportHost:      DefaultExportHost,
		ExportPort:      DefaultExportPort,
		Timeout:         Duration{DefaultTimeout},
		RedisServerAddr: DefaultRedisServerAddr,
		RedisPassword:   DefaultRedisPassword,
		RedisDBIndex:    DefaultRedisDBIndex,
	}

	return config
}
