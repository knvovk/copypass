package config

import "github.com/BurntSushi/toml"

type Config struct {
	App app      `toml:"app"`
	DB  database `toml:"database"`
	Log logging  `toml:"logging"`
}

type app struct {
	IsDebug    bool   `toml:"is_debug"`
	Host       string `toml:"host"`
	Port       int    `toml:"port"`
	SigningKey string `toml:"signing_key"`
}

type database struct {
	URL             string `toml:"url"`
	MaxPoolSize     int32  `toml:"max_pool_size"`
	MaxConnIdleTime int    `toml:"max_conn_idle_time"`
}

type logging struct {
	Path           string `toml:"path"`
	DateTimeFormat string `toml:"datetime_format"`
}

func GetConfig(path string) (*Config, error) {
	var config Config
	if _, err := toml.DecodeFile(path, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
