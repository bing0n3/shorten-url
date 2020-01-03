package utils

import (
	"errors"
	"math"
	"strings"
)

var (
	// CharacterSet consists of 62 characters [0-9][A-Z][a-z].
	base int64 = 62
	characterSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

// Encode return a base representation as string of the given integer number.
func Encode(num int64) string {
	b := make([]byte, 0)
	if num == 0 {
		return string(characterSet[0])
	}
	for num > 0 {
		// remain
		rem := num % base
		num /= base

		b = append([]byte{characterSet[rem]}, b...)
	}


	return string(b)
}


// Decode return a integer of a base62 encoded string.
func Decode(s string) (int64, error){

	var (
		r int64
		pos int
		pow int
	)

	// loop through input
	for i, v := range s {
		pow = len(s) - (i + 1)

		pos = strings.IndexRune(characterSet, v)

		if pos == -1 {
			return int64(pos), errors.New("invalid character: "+ string(v))
		}
		r += int64(pos) * int64(math.Pow(float64(base),float64(pow)))
	}

	return r, nil
}