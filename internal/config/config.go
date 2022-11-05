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
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Name     string `toml:"name"`
	User     string `toml:"user"`
	Password string `toml:"password"`
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
