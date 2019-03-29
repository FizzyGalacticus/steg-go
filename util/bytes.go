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
		val := int8(b) & (1 << uint8(8-i))

		bits[i-1] = val > 0
	}

	return bits
}

// GetByteFromBits - Converts 8-bit array to
// a byte
func GetByteFromBits(bits []bool) byte {
	var b int8

	for i, bit := range bits {
		if bit {
			b = b | (1 << uint8(7-i))
		}
	}

	return byte(b)
}

// GetBytesFromBits - Converts a bit array to
// a byte array
func GetBytesFromBits(bits []bool) []byte {
	bytes := make([]byte, len(bits)/8)

	for i := 0; i < len(bits); i += 8 {
		bytes[i/8] = GetByteFromBits(bits[i : i+8])
	}

	return bytes
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

// GetBitStringFromBytes - Retrieves the bit string
// of a byte array
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

// GetBitStringFromString - Retrieves the decoded string
// of a bit string
func GetBitStringFromString(str string) string {
	return GetBitStringFromBytes([]byte(str))
}

// GetBitsFromBitString - Converts a bit string to a
// bit array
func GetBitsFromBitString(bitStr string) []bool {
	bits := make([]bool, len(bitStr))

	for i, char := range bitStr {
		if string(char) == "1" {
			bits[i] = true
		} else {
			bits[i] = false
		}
	}

	return bits
}

// GetBytesFromBitString - Converts a bit string to a
// byte array
func GetBytesFromBitString(bitStr string) []byte {
	bytes := make([]byte, len(bitStr)/8)

	for i := 0; i < len(bitStr); i += 8 {
		bits := GetBitsFromBitString(bitStr[i : i+8])
		bytes[i/8] = GetByteFromBits(bits)
	}

	return bytes
}

// GetStringFromBitString - Converts a bit string to a
// string
func GetStringFromBitString(bitStr string) string {
	return string(GetBytesFromBitString(bitStr))
}
