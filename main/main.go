package main

// package main

// Following:
// https://stackoverflow.com/questions/33186783/get-a-pixel-array-from-from-golang-image-image

import (
	"fmt"
	"steg/config"
	"steg/util"
	"strings"
)

func main() {
	coverFile := config.GetCoverFile()
	stegFile := config.GetStegFile()
	files := config.GetFiles()

	coverImage := util.DecodeImage(coverFile)

	imageCapacity := util.GetImageBitCapacity(coverImage)

	// File bits + number of files
	requiredCapacity := util.GetRequiredSizeForFiles(files) + 8

	if imageCapacity < requiredCapacity {
		var errBuilder strings.Builder

		errBuilder.WriteString("Cannot fit files into image.\n")
		errBuilder.WriteString(fmt.Sprintf("Required space: %d bytes\n", requiredCapacity/8))
		errBuilder.WriteString(fmt.Sprintf("Available space: %d bytes\n", imageCapacity/8))

		panic(errBuilder.String())
	} else {
		// Process as normal
		bits := make(util.Bits, 0)

		// Metadata for number of files
		numFiles := len(files)
		numFileBits := util.GetBitsFromInt8(int8(numFiles))

		bits = append(bits, numFileBits...)

		// Gather payload
		for _, filename := range files {
			bits = append(bits, util.BuildFilePayload(filename)...)
		}

		// Insert payload into image
		util.InsertBitsIntoImage(coverImage, bits)

		// Save new image
		util.SaveImage(coverImage, stegFile)
	}
}
