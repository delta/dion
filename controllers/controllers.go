package controllers

import "go.uber.org/fx"

type Controller interface{}

func New() (Controller, error) {
	return struct{}{}, nil
}

var Module = fx.Provide(
	New,
)
