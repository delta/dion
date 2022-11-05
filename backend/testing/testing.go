package testing

// Reason behind this: https://brandur.org/fragments/testing-go-project-root
// This problem is occuring when I want to test config merging in config/config_test.go

import (
	"os"
	"path"
	"runtime"
)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}
