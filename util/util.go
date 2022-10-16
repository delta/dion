package util

import (
	"net/http"
)

func UnwrapResult[T any](result *T, err error) (int, *T){
	if err == nil {
		return http.StatusOK, result
	} else {
		return http.StatusInternalServerError, nil
	}
}
