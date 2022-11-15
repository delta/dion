package logging

import (
	"encoding/json"

	"go.uber.org/zap"
)

// Logger with fields.
// Inspired by logrus - https://github.com/sirupsen/logrus
type Fields map[string]interface{}

// Takes in an array of data and returns a named zap logger with the
// given fields
//
// This method can be used to give context a group of logs.
// how to use example,
//
//	l, err := logging.WithFields(logging.Fields{"hello": "world"})
//	l.Sugar().Infof("Hello world")
//
// this will create a named zap logger (https://pkg.go.dev/go.uber.org/zap@v1.19.0#Logger.Named)
// with the name `{"hello": "world"}`
// any log created using l will have the name to it
// example
//
//	2022-11-08T17:37:01.405+0530    INFO    {"hello":"world"}       backend/main.go:16      Hello world
func WithFields(fields Fields) (*zap.Logger, error) {
	data, err := json.Marshal(fields)
	if err != nil {
		return nil, err
	}
	return L.Named(string(data)), nil
}
