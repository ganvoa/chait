package internal

import (
	"errors"
	"io"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Chait struct {
		Rol1    string `yaml:"rol1"`
		Rol2    string `yaml:"rol2"`
		Replies int    `yaml:"replies"`
	} `yaml:"chait"`
}

func NewConfig(handle io.Reader) (*Config, error) {
	d := yaml.NewDecoder(handle)

	config := &Config{}

	if err := d.Decode(&config); err != nil {
		return nil, errors.New("error decoding yaml")
	}

	if config.Chait.Replies <= 0 {
		return nil, errors.New("replies must be > 0")
	}

	if config.Chait.Rol1 == "" {
		return nil, errors.New("rol1 required")
	}

	if config.Chait.Rol2 == "" {
		return nil, errors.New("rol2 required")
	}

	return config, nil
}
