package main_test

import (
	"testing"

	_ "delta.nitt.edu/dion/testing"
	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}
