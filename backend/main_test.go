package main_test

import (
	"testing"

	"delta.nitt.edu/dion"
	"go.uber.org/fx"
	"go.uber.org/goleak"
)

func TestFx(t *testing.T) {
	fx.ValidateApp(
		main.App,
	)
}

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}
