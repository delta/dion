package repository

import "go.uber.org/fx"

type Repository interface{}

func New() (Repository, error) {
	return struct{}{}, nil
}

var Module = fx.Provide(
	New,
)
