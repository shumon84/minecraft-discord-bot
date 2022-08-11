package config

import (
	"encoding/json"
	"io"
)

type Config struct {
	Discord   Discord   `json:"discord"`
	Minecraft Minecraft `json:"minecraft"`
}

func NewConfig(r io.Reader) (Config, error) {
	config := Config{}
	decoder := json.NewDecoder(r)
	if err := decoder.Decode(&config); err != nil {
		return Config{}, err
	}
	return config, nil
}
