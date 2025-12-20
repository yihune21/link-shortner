package genkey

import (
	"errors"
	"math"
	"strings"
)

const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Encode(number uint64) string {
	if number == 0 {
		return string(alphabet[0])
	}

	var sb strings.Builder
	base := uint64(len(alphabet))

	for number > 0 {
		remainder := number % base
		sb.WriteByte(alphabet[remainder])
		number = number / base
	}

	return reverse(sb.String())
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func Decode(encoded string) (uint64, error) {
	var number uint64
	base := uint64(len(alphabet))
	
	for i, char := range encoded {
		idx := strings.IndexRune(alphabet, char)
		if idx == -1 {
			return 0, errors.New("invalid character in encoded string")
		}

		power := len(encoded) - 1 - i
		number += uint64(idx) * uint64(math.Pow(float64(base), float64(power)))
	}

	return number, nil
}
