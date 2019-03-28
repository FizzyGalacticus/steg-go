package util

// GetNumBitsFromBytes - Returns the number of
// bits that are in a given set of bytes
func GetNumBitsFromBytes(bytes []byte) int64 {
	return int64(len(bytes) * 8)
}
