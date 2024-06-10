package utils

import (
	"errors"
	"net/http"
	"strings"
)

func HandleHttpError(w http.ResponseWriter, err error) bool {
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return true
	}
	return false
}

func HandleHttpErrors(w http.ResponseWriter, errorsParam ...error) bool {
	errStrings := make([]string, 0)
	for _, err := range errorsParam {
		if err != nil {
			errStrings = append(errStrings, err.Error())
		}
	}

	if len(errStrings) == 0 {
		return false
	}

	joinedErrors := errors.New(strings.Join(errStrings, "\n"))
	return HandleHttpError(w, joinedErrors)
}
