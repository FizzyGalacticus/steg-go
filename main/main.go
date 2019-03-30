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

	// File bytes + a single byte for how many files
	fileSizes := util.GetFileSizes(files)
	requiredCapacity := (fileSizes + 1) * 8

	for _, filename := range files {
		// filename bytes
		requiredCapacity += int64(len(filename) * 8)
	}

	if imageCapacity < requiredCapacity {
		var errBuilder strings.Builder

		errBuilder.WriteString("Cannot fit files into image.\n")
		errBuilder.WriteString(fmt.Sprintf("Required space: %d bytes\n", requiredCapacity/8))
		errBuilder.WriteString(fmt.Sprintf("Available space: %d bytes\n", imageCapacity/8))

		panic(errBuilder.String())
	} else {
		// Process as normal

		// Save new image
		util.SaveImage(coverImage, stegFile)
	}
}
