package sanitizer

import (
	"strings"
	"unicode"
)

func sanitizerBase(sanitizerFunc func(rune) bool, str string) string {
	sanitized := strings.Map(func(r rune) rune {
		if !sanitizerFunc(r) {
			return -1
		}
		return r
	}, str)
	return sanitized
}

func SanitizeName(name string) string {
	return sanitizerBase(unicode.IsLetter, name)
}

func SanitizePhoneNumber(phoneNumber string) string {
	return sanitizerBase(unicode.IsNumber, phoneNumber)
}
