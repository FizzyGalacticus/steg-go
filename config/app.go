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
	InFile  string   `json:"inFile"`
	OutFile string   `json:"outFile"`
	Files   []string `json:"files"`
}

// GetArguments - Retrieves the full
// Arguments object
func GetArguments() Arguments {
	if len(os.Args) < 4 {
		usage := "Usage:\n\t%s <input_image> <output_image> ...<steg_files>\n"

		fmt.Printf(usage, os.Args[0])
		os.Exit(1)
	}

	inFile, outFile, files := getParams()

	return Arguments{
		InFile:  inFile,
		OutFile: outFile,
		Files:   files,
	}
}

// GetInFile - Retrieves the path to the
// image file to use as a source
func GetInFile() string {
	args := GetArguments()

	return args.InFile
}

// GetOutFile - Retrieves the path to the
// image file to use as the output
func GetOutFile() string {
	args := GetArguments()

	return args.OutFile
}

// GetFiles - Retrieves the list of
// paths of files to store in the image
func GetFiles() []string {
	args := GetArguments()

	return args.Files
}
