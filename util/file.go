package util

import (
	"io/ioutil"
	"net/http"
	"os"
)

// GetFile - Returns a pointer to a file reader
func GetFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	return file
}

// GetFileBytes - Returns the bytes of a file
func GetFileBytes(filename string) Bytes {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return file
}

// GetFileBits - Returns the bits of a file
func GetFileBits(filename string) Bits {
	bytes := GetFileBytes(filename)
	return GetBitsFromBytes(bytes)
}

// GetFileSize - Returns the size of the file
// with the given path
func GetFileSize(filename string) int64 {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	info, err := file.Stat()
	if err != nil {
		panic(err)
	}

	return info.Size()
}

// GetFileSizes - Returns the total size (in bytes)
// of the given filenames
func GetFileSizes(filenames []string) int64 {
	var total int64

	for _, filename := range filenames {
		total += GetFileSize(filename)
	}

	return total
}

func getContentTypeFromReader(reader func([]byte) (int, error)) string {
	// Only the first 512 bytes are used to sniff the content type.
	buffer := make(Bytes, 512)

	_, err := reader(buffer)
	if err != nil {
		panic(err)
	}

	return http.DetectContentType(buffer)
}

// GetContentTypeFromURL - Gets the content/mime type
// of a file from a URL
func GetContentTypeFromURL(url string) string {
	req, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer req.Body.Close()

	return getContentTypeFromReader(req.Body.Read)
}

// GetFileContentType - Gets the content/mime type
// of a file from a file path
func GetFileContentType(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	return getContentTypeFromReader(file.Read)
}
