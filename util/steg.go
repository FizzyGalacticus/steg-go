package util

import (
	"errors"
	"fmt"
	"math"
	"strings"
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

	filenameSizeStart := 0
	filenameStart := 8
	sizeStart := filenameStart + len(filenameBits)
	fileStart := sizeStart + len(sizeBits)

	copy(bits[filenameSizeStart:filenameStart], filenameSizeBits)
	copy(bits[filenameStart:sizeStart], filenameBits)
	copy(bits[sizeStart:fileStart], sizeBits)
	copy(bits[fileStart:], fileBits)

	return bits
}

func GetFileFromPayload(bits Bits) (string, Bytes) {
	bytes := GetBytesFromBits(bits)

	filenameSize := int(bytes[0])

	filenameEnd := filenameSize + 1
	filename := string(bytes[1:filenameEnd])

	sizeEnd := filenameEnd + 8
	size := GetInt64FromBytes(bytes[filenameEnd:sizeEnd])

	fileEnd := int64(sizeEnd) + size
	file := bytes[sizeEnd:fileEnd]

	return filename, file
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

func HideFilesInImage(coverImagePath string, stegImagePath string, files []string) error {
	coverImage := DecodeImage(coverImagePath)

	imageCapacity := GetImageBitCapacity(coverImage)

	// File bits + number of files
	requiredCapacity := GetRequiredSizeForFiles(files) + 8

	if imageCapacity < requiredCapacity {
		var errBuilder strings.Builder

		errBuilder.WriteString("Cannot fit files into image.\n")
		errBuilder.WriteString(fmt.Sprintf("Required space: %d bytes\n", requiredCapacity/8))
		errBuilder.WriteString(fmt.Sprintf("Available space: %d bytes\n", imageCapacity/8))

		return errors.New(errBuilder.String())
	}

	// Process as normal
	bits := make(Bits, 0)

	// Metadata for number of files
	numFiles := len(files)
	numFileBits := GetBitsFromInt8(int8(numFiles))

	bits = append(bits, numFileBits...)

	// Gather payload
	for _, filename := range files {
		bits = append(bits, BuildFilePayload(filename)...)
	}

	// Insert payload into image
	InsertBitsIntoImage(coverImage, bits)

	// Save new image
	SaveImage(coverImage, stegImagePath)

	return nil
}

func GetFilesFromImage(stegImagePath string) ([]Bytes, error) {
	return make([]Bytes, 0), nil
}
