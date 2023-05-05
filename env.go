package tomlutil

import (
	"github.com/BurntSushi/toml"
)

type EnvConfig struct {
	KEY string `toml:"KEY"`
}

func TomlKey(filepath string) (*EnvConfig, error) {
	var config EnvConfig

	if _, err := toml.DecodeFile(filepath, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
