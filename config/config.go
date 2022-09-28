package config

import "go.uber.org/fx"

type Config interface{}

type config struct{}

func New() (Config, error) {
	return config{}, nil
}

var Module = fx.Provide(
	New,
)
