package config

import (
	"fmt"
	"os"
)

func getParams() (string, string, []string) {
	return os.Args[1], os.Args[2], os.Args[3:]
}

// Arguments - Input arguments to use for
// steg processing
type Arguments struct {
	CoverFile string   `json:"coverFile"`
	StegFile  string   `json:"stegFile"`
	Files     []string `json:"files"`
}

// GetArguments - Retrieves the full
// Arguments object
func GetArguments() Arguments {
	if len(os.Args) < 4 {
		usage := "Usage:\n\t%s <cover_image> <steg_image> ...<files>\n"

		fmt.Printf(usage, os.Args[0])
		os.Exit(1)
	}

	coverFile, stegFile, files := getParams()

	return Arguments{
		CoverFile: coverFile,
		StegFile:  stegFile,
		Files:     files,
	}
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
