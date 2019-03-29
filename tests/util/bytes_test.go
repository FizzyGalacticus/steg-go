package tests

import (
	"steg/util"
	"testing"
)

func TestGetNumBitsFromBytes(t *testing.T) {
	type testTable struct {
		n        []byte
		expected int64
	}

	tests := []testTable{
		{[]byte("a"), 8},
		{[]byte("aa"), 16},
		{[]byte("aaa"), 24},
		{[]byte("aaaa"), 32},
	}

	for _, tt := range tests {
		actual := util.GetNumBitsFromBytes(&tt.n)

		if actual != tt.expected {
			t.Errorf("GetNumBitsFromBytes(%s): expected %d, got %d", tt.n, tt.expected, actual)
		}
	}
}

func TestGetBitsFromByte(t *testing.T) {
	type testTable struct {
		n        byte
		expected []bool
	}

	tests := []testTable{
		{[]byte("a")[0], []bool{false, true, true, false, false, false, false, true}},
	}

	for _, tt := range tests {
		actual := util.GetBitsFromByte(tt.n)

		for i, bit := range tt.expected {
			if bit != actual[i] {
				t.Errorf("GetBitsFromByte(%s): expected %t, got %t", string(tt.n), tt.expected, actual)
			}
		}
	}
}

func TestGetByteFromBits(t *testing.T) {
	type testTable struct {
		n        []bool
		expected byte
	}

	tests := []testTable{
		{[]bool{false, true, true, false, false, false, false, true}, []byte("a")[0]},
	}

	for _, tt := range tests {
		actual := util.GetByteFromBits(tt.n)

		if actual != tt.expected {
			t.Error("GetByteFromBits failed: ", util.GetBytesFromBits(tt.n), " != ", tt.expected)
		}
	}
}
