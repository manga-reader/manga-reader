package config

import (
	"encoding/json"
)

func init() {
	Cfg = setDefaultConfig()
}

var Cfg *Configuration

type Configuration struct {
	Connection Connection `json:"connection,omitempty"`
	Redis      Redis      `json:"redis,omitempty"`
}

type Connection struct {
	ExportHost string   `json:"host,omitempty"`
	ExportPort string   `json:"port,omitempty"`
	Timeout    Duration `json:"timeout,omitempty"`
	JWTSecret  string   `json:"jwt_secret,omitempty"`
}

type Redis struct {
	ServerAddr string `json:"server_address,omitempty"`
	Password   string `json:"password,omitempty"`
	DBIndex    int    `json:"db_index,omitempty"`
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
		Connection: Connection{
			ExportHost: DefaultExportHost,
			ExportPort: DefaultExportPort,
			Timeout:    Duration{DefaultTimeout},
			JWTSecret:  DefaultJWTSecret,
		},
		Redis: Redis{
			ServerAddr: DefaultRedisServerAddr,
			Password:   DefaultRedisPassword,
			DBIndex:    DefaultRedisDBIndex,
		},
	}

	return config
}
