package models

import "go.uber.org/fx"

type Model interface{}

func New() (Model, error) {
	return struct{}{}, nil
}

var Module = fx.Provide(
	New,
)
