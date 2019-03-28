package util

import (
	"strings"
)

// GetNumBitsFromBytes - Returns the number of
// bits that are in a given set of bytes
func GetNumBitsFromBytes(bytes *[]byte) int64 {
	return int64(len(*bytes) * 8)
}

// GetBitsFromByte - Retrieves the bits of a byte
// as an array of bool
func GetBitsFromByte(b byte) []bool {
	bits := make([]bool, GetNumBitsFromBytes(&[]byte{b}))

	for i := 1; i < 9; i++ {
		val := int(b) & (1 << uint(8-i))

		bits[i-1] = val > 0
	}

	return bits
}

// GetBitsFromBytes - Retrieves the bits of a byte
// array as an array of bool
func GetBitsFromBytes(bytes []byte) []bool {
	bits := make([]bool, GetNumBitsFromBytes(&bytes))

	for i, b := range bytes {
		for j, bit := range GetBitsFromByte(b) {
			bits[i*8+j] = bit
		}
	}

	return bits
}

// GetBitStringFromBytes - Retrieves the bits of a byte
// array as a string
func GetBitStringFromBytes(bytes []byte) string {
	var builder strings.Builder
	bits := GetBitsFromBytes(bytes)

	for _, bit := range bits {
		if bit {
			builder.WriteString("1")
		} else {
			builder.WriteString("0")
		}
	}

	return builder.String()
}
