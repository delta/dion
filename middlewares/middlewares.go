package middlewares

import "go.uber.org/fx"

type Middleware interface{}

func New() (Middleware, error) {
	return struct{}{}, nil
}

var Module = fx.Provide(
	New,
)
