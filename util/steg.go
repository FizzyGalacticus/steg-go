package util

import (
	"fmt"
	"math"
)

func GetRequiredSizeForFile(filename string) int64 {
	size := GetFileSize(filename)

	return (size + int64(len(filename)+8+1)) * 8
}

func GetRequiredSizeForFiles(filenames []string) int64 {
	size := int64(0)

	for _, filename := range filenames {
		size += GetRequiredSizeForFile(filename)
	}

	return size
}

func BuildFilePayload(filename string) Bits {
	size := GetFileSize(filename)

	sizeBits := GetBitsFromInt64(size)
	filenameBits := GetBitsFromBytes(Bytes(filename))
	filenameSizeBits := GetBitsFromInt8(int8(len(filename)))
	fileBits := GetFileBits(filename)

	bits := make(Bits, GetRequiredSizeForFile(filename))

	filenameStart := 0
	filenameSizeStart := len(filenameBits)
	sizeStart := filenameSizeStart + len(filenameSizeBits)
	fileStart := sizeStart + len(sizeBits)

	copy(bits[filenameStart:filenameSizeStart], filenameBits)
	copy(bits[filenameSizeStart:sizeStart], filenameSizeBits)
	copy(bits[sizeStart:fileStart], sizeBits)
	copy(bits[fileStart:], fileBits)

	return bits
}

func insertBitsIntoPixel(bits Bits, pix Pixel) Pixel {
	if len(bits) > 3 {
		panic(fmt.Sprintf("Too many bits for a pixel: %d", len(bits)))
	}

	rgbVals := []uint32{pix.R, pix.G, pix.B}

	for i, val := range rgbVals {
		if i < len(bits) {
			if bits[i] {
				val = val & 1
			} else if val%2 == 1 {
				val--
			}
		}
	}

	return Pixel{
		R: rgbVals[0],
		G: rgbVals[1],
		B: rgbVals[2],
		A: pix.A,
	}
}

func InsertBitsIntoImage(img *EditableImage, bits Bits) {
	imgWidth, _ := GetImageSize(img)

	bitsPerWidth := imgWidth * 3

	for i := 0; i < len(bits); i += 3 {
		lastIndex := i + 3

		if lastIndex > len(bits) {
			lastIndex = len(bits)
		}

		x := i % bitsPerWidth
		y := int(math.Floor(float64(i / bitsPerWidth)))

		pix := GetPixelFromRGBA(img.At(x, y).RGBA())

		newPix := insertBitsIntoPixel(bits[i:lastIndex], pix)

		img.Set(x, y, &newPix)
	}
}
