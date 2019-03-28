package util

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"os"
	"path"
)

func addFileToZip(zipWriter *zip.Writer, filename string) error {
	fileToZip, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fileToZip.Close()

	// Get the file information
	info, err := fileToZip.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	// Using FileInfoHeader() above only uses the basename of the file. If we want
	// to preserve the folder structure we can overwrite this with the full path.
	header.Name = path.Base(filename)

	// Change to deflate to gain better compression
	// see http://golang.org/pkg/archive/zip/#pkg-constants
	header.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(writer, fileToZip)

	return err
}

func compressFiles(files []string, library string) []byte {
	var buf bytes.Buffer
	var w *zip.Writer

	switch library {
	case "zip":
		w = zip.NewWriter(&buf)
		break
	// case "gzip":
	// 	w = gzip.NewWriter(&buf)
	// 	break
	default:
		panic(fmt.Sprintf("Can't determine type of compression with type '%s'", library))
	}

	defer w.Close()

	for _, file := range files {
		err := addFileToZip(w, file)
		if err != nil {
			panic(err)
		}
	}

	return buf.Bytes()
}

// ZipFiles - Takes file paths to compress into
// a single zipfile, and returns the zipfile bytes
func ZipFiles(files []string) []byte {
	return compressFiles(files, "zip")
}

// func GZipFiles(files []string) []byte {
// 	return compressFiles(files, "gzip")
// }
