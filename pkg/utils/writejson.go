package utils

import "net/http"

func WriteJson(w http.ResponseWriter, httpStatusCode int, json []byte) (int, error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	return w.Write(json)
}
