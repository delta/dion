package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/goleak"
)

func TestSanity(t *testing.T) {
	assert.Equal(t, 1, 1)
}

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}
