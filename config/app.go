package config

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

// Arguments - Input arguments to use for
// steg processing
type Arguments struct {
	CoverFile string   `json:"coverFile"`
	StegFile  string   `json:"stegFile"`
	Mode      string   `json:"mode"`
	Files     []string `json:"files"`
}

var coverFlag = flag.String("c", "", "The image to use as source when encoding.")
var stegFlag = flag.String("s", "", "The path to save the image with payload to.\n\tNote: must be lostless image format.")
var modeFlag = flag.String("m", "e", "The mode to run in: 'e' for 'encode', 'd' for 'decode'")

func init() {
	flag.Parse()
}

// GetArguments - Retrieves the full
// Arguments object
func GetArguments() Arguments {
	filesFlag := flag.Args()

	inputErr := false

	switch *modeFlag {
	case "e":
		noCover := *coverFlag == ""
		noSteg := *stegFlag == ""
		noFiles := len(filesFlag) == 0

		inputErr = noCover || noSteg || noFiles
		break
	case "d":
		noSteg := *stegFlag == ""

		inputErr = noSteg
		break
	default:
		inputErr = true
		fmt.Println("Mode must be 'e' or 'd'")
	}

	if inputErr {
		flag.PrintDefaults()
		os.Exit(1)
	}

	cover, err := filepath.Abs(*coverFlag)
	if err != nil {
		panic(err)
	}

	steg, err := filepath.Abs(*stegFlag)
	if err != nil {
		panic(err)
	}

	files := make([]string, len(filesFlag))
	for i, file := range filesFlag {
		path, err := filepath.Abs(file)
		if err != nil {
			panic(err)
		}

		files[i] = path
	}

	return Arguments{
		CoverFile: cover,
		StegFile:  steg,
		Files:     files,
		Mode:      *modeFlag,
	}
}

func GetMode() string {
	args := GetArguments()

	if args.Mode == "e" {
		return "encoding"
	}

	return "decoding"
}

// GetCoverFile - Retrieves the path to the
// image file to use as a source
func GetCoverFile() string {
	args := GetArguments()

	return args.CoverFile
}

// GetStegFile - Retrieves the path to the
// image file to use as the output
func GetStegFile() string {
	args := GetArguments()

	return args.StegFile
}

// GetFiles - Retrieves the list of
// paths of files to store in the image
func GetFiles() []string {
	args := GetArguments()

	return args.Files
}
