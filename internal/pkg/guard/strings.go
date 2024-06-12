package guard

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"unicode"
)

// AgainstEmptyString checks if a field of given type T is empty and returns an error in case it is
//   - value, must be a struct
//   - fieldName, must use exactly the same casing as the struct field name
func AgainstEmptyString[T any](value T, fieldName string) error {
	typ := reflect.ValueOf(value)
	field := typ.FieldByName(fieldName)
	if !field.IsValid() {
		log.Panicf(fmt.Sprintf("Invalid field name: %v for type: %T", fieldName, value))
	}

	str := field.String()
	if len(str) != 0 {
		for _, r := range str {
			if !unicode.IsSpace(r) {
				return nil
			}
		}
	}

	return errors.New(fmt.Sprintf("%v: empty string", fieldName))
}
