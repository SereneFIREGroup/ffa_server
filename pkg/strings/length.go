package strings

import (
	"github.com/gookit/goutil/strutil"
)

// IsEmpty check string is empty.
func IsEmpty(str string) bool {
	return len(str) == 0
}

// LenUTF8 returns the length of the string.
func LenUTF8(str string) int {
	return strutil.Utf8Len(str)
}

// IsLenValid check string length is valid.
func IsLenValid(str string, length int) bool {
	return len(str) <= length
}

// IsLenValidUTF8 check utf8 string length is valid.
func IsLenValidUTF8(str string, length int) bool {
	return LenUTF8(str) <= length
}
