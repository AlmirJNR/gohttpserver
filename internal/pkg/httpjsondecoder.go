package pkg

import (
	"encoding/json"
	"io"
)

func GetDecoder(r io.Reader) *json.Decoder {
	decoder := json.NewDecoder(r)
	decoder.DisallowUnknownFields()
	return decoder
}
