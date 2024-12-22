package configs

import (
	"encoding/json"
	"os"
)

type Config struct {
	PathFile string `json:"path_file"`
	Server   struct {
		Port int `json:"port"`
	} `json:"server"`
}

var config *Config

func LoadFile(path string) (*Config, error) {
	fileContent, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	err = json.Unmarshal(fileContent, &cfg)
	if err != nil {
		return nil, err
	}
	config = &cfg
	return config, nil
}

func GetConfig() *Config {
	return config
}
