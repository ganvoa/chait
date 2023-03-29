package internal

import (
	"errors"
	"io"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Chait struct {
		RoleU1     string `yaml:"roleU1"`
		RoleU2     string `yaml:"roleU2"`
		Iterations int    `yaml:"iterations"`
	} `yaml:"chait"`
}

func NewConfig(handle io.Reader) (*Config, error) {
	d := yaml.NewDecoder(handle)

	config := &Config{}

	if err := d.Decode(&config); err != nil {
		return nil, errors.New("error decoding yaml")
	}

	if config.Chait.Iterations <= 0 {
		return nil, errors.New("iterations must be > 0")
	}

	if config.Chait.RoleU1 == "" {
		return nil, errors.New("roleU1 required")
	}

	if config.Chait.RoleU2 == "" {
		return nil, errors.New("roleU2 required")
	}

	return config, nil
}
