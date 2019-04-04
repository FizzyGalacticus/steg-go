package main

// package main

// Following:
// https://stackoverflow.com/questions/33186783/get-a-pixel-array-from-from-golang-image-image

import (
	"steg/config"
	"steg/util"
)

func main() {
	mode := config.GetMode()
	var err error

	switch mode {
	case "encoding":
		coverFile := config.GetCoverFile()
		stegFile := config.GetStegFile()
		files := config.GetFiles()
		err = util.HideFilesInImage(coverFile, stegFile, files)
		break
	case "decoding":
		stegFile := config.GetStegFile()
		_, err = util.GetFilesFromImage(stegFile)
		break
	}

	if err != nil {
		panic(err)
	}
}
