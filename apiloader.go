package tomlutil

import (
	"github.com/BurntSushi/toml"
)

type Service struct {
	KEY string `toml:"KEY"`
}

func TomlKey(filepath string) (*Service, error) {
	var Load Service

	if _, err := toml.DecodeFile(filepath, &Load); err != nil {
		return nil, err
	}

	return &Load, nil
}
