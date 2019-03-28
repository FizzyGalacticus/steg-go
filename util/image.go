package util

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
)

func init() {
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	image.RegisterFormat("jpg", "jpg", jpeg.Decode, jpeg.DecodeConfig)
}

// Takes a filesystem path and returns
// a decoded image
func decodeImagePath(path string) image.Image {
	file, err := os.Open(path)

	if err != nil {
		panic(fmt.Sprintf("Could not open image '%s'", path))
	}

	defer file.Close()

	img, _, err := image.Decode(file)

	if err != nil {
		panic(err)
	}

	return img
}

// Takes a URL and returns a decoded image
func decodeImageURL(url string) image.Image {
	req, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer req.Body.Close()

	img, _, err := image.Decode(req.Body)

	if err != nil {
		panic(err)
	}

	return img
}

// DecodeImage takes a filesystem path or
// a URL and returns the decoded image
func DecodeImage(path string) image.Image {
	if IsURL(path) {
		return decodeImageURL(path)
	}

	return decodeImagePath(path)
}

// GetImageSize - Returns the image width and height
func GetImageSize(i image.Image) (int, int) {
	bounds := i.Bounds()

	return bounds.Max.X, bounds.Max.Y
}

// GetImageBitCapacity - Returns how many bits can
// fit inside of an image
func GetImageBitCapacity(i image.Image) int64 {
	imgW, imgH := GetImageSize(i)
	numPix := imgW * imgH

	return int64(numPix * 3)
}
