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
	inFile := config.GetInFile()
	outFile := config.GetOutFile()
	stegFiles := config.GetFiles()

	inputImage := util.DecodeImage(inFile)

	zipFile := util.ZipFiles(stegFiles)

	width, height := util.GetImageSize(inputImage)

	fmt.Println("Input: " + inFile)
	fmt.Println("Output: " + outFile)
	fmt.Println("Steg Files: " + strings.Join(stegFiles, ","))
	fmt.Printf("Image Dimensions: %dx%d\n", width, height)
	fmt.Printf("Image Bit Capacity: %d\n", util.GetImageBitCapacity(inputImage))
	fmt.Printf("Required space: %d\n", util.GetNumBitsFromBytes(zipFile))
}
