package tests

import (
	"steg/util"
	"testing"
)

func TestGetNumBitsFromBytes(t *testing.T) {
	type testTable struct {
		n        util.Bytes
		expected int64
	}

	tests := []testTable{
		{util.Bytes("a"), 8},
		{util.Bytes("aa"), 16},
		{util.Bytes("aaa"), 24},
		{util.Bytes("aaaa"), 32},
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
		expected util.Bits
	}

	tests := []testTable{
		{util.Bytes("a")[0], util.Bits{false, true, true, false, false, false, false, true}},
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
		n        util.Bits
		expected byte
	}

	tests := []testTable{
		{util.Bits{false, true, true, false, false, false, false, true}, util.Bytes("a")[0]},
	}

	for _, tt := range tests {
		actual := util.GetByteFromBits(tt.n)

		if actual != tt.expected {
			t.Error("GetByteFromBits failed: ", util.GetBytesFromBits(tt.n), " != ", tt.expected)
		}
	}
}

func TestGetBitsFromBytes(t *testing.T) {
	t.Skip()
}

func TestGetBitStringFromBytes(t *testing.T) {
	t.Skip()
}

func TestGetBitStringFromString(t *testing.T) {
	t.Skip()
}

func TestGetBitsFromBitString(t *testing.T) {
	t.Skip()
}

func TestGetBytesFromBitString(t *testing.T) {
	t.Skip()
}

func TestGetStringFromBitString(t *testing.T) {
	t.Skip()
}
