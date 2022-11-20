package utils

import (
	"os"
	"strconv"
)

func GetInt(key string, defaultVal int) int {
	value, present := os.LookupEnv(key)
	var (
		valueInt int
		err      error
	)
	if !present {
		valueInt = defaultVal
	} else {
		valueInt, err = strconv.Atoi(value)
		if err != nil {
			valueInt = defaultVal 
		}
	}
	return valueInt
}
