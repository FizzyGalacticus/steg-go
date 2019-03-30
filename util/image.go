package util

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"os"
)

func init() {
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	image.RegisterFormat("jpg", "jpg", jpeg.Decode, jpeg.DecodeConfig)
}

// EditableImage is a custom struct that allows
// an image.Image to be editable
type EditableImage struct {
	image.Image
	mime      string
	customPix map[image.Point]color.Color
}

func newEditableImage(reader io.Reader, mime string) *EditableImage {
	var img image.Image
	var err error

	switch mime {
	case "image/jpeg":
		img, err = jpeg.Decode(reader)
		break
	case "image/png":
		img, err = png.Decode(reader)
		break
	default:
		panic(fmt.Sprintf("Image type not supported: %s", mime))
	}

	if err != nil {
		panic(err)
	}

	return &EditableImage{
		img,
		mime,
		map[image.Point]color.Color{},
	}
}

// Set - Sets the color of the pixel at the given position
func (e *EditableImage) Set(x, y int, c color.Color) {
	e.customPix[image.Point{x, y}] = c
}

// At - Retrieves the color of the pixel at the given position
func (e *EditableImage) At(x, y int) color.Color {
	if c := e.customPix[image.Point{x, y}]; c != nil {
		return c
	}

	return e.Image.At(x, y)
}

// Pixel - Alias for color.RGBA
type Pixel color.RGBA

// Takes a filesystem path and returns
// a decoded image
func decodeImagePath(path string) *EditableImage {
	file, err := os.Open(path)

	if err != nil {
		panic(fmt.Sprintf("Could not open image '%s'", path))
	}

	defer file.Close()

	return newEditableImage(file, GetFileContentType(path))
}

// Takes a URL and returns a decoded image
func decodeImageURL(url string) *EditableImage {
	req, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer req.Body.Close()

	return newEditableImage(req.Body, GetContentTypeFromURL(url))
}

// DecodeImage takes a filesystem path or
// a URL and returns the decoded image
func DecodeImage(path string) *EditableImage {
	if IsURL(path) {
		return decodeImageURL(path)
	}

	return decodeImagePath(path)
}

// GetImageSize - Returns the image width and height
func GetImageSize(i *EditableImage) (int, int) {
	bounds := i.Bounds()

	return bounds.Max.X, bounds.Max.Y
}

// GetImageBitCapacity - Returns how many bits can
// fit inside of an image
func GetImageBitCapacity(i *EditableImage) int64 {
	imgW, imgH := GetImageSize(i)
	numPix := imgW * imgH

	return int64(numPix * 3)
}

// GetImagePixel - Retrieves an images pixel values at
// the given position
func GetImagePixel(i *EditableImage, x int, y int) Pixel {
	r, g, b, a := (*i).At(x, y).RGBA()

	return Pixel{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: uint8(a),
	}
}

// SetImagePixel - Sets the pixel color at given coordinates
func SetImagePixel(i *EditableImage, x int, y int, p Pixel) {
	i.Set(x, y, color.RGBA{
		R: p.R,
		G: p.G,
		B: p.B,
		A: p.A,
	})
}

// SaveImage - Saves the given image to the given filename
func SaveImage(i *EditableImage, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	switch i.mime {
	case "image/jpeg":
		err = jpeg.Encode(file, i, nil)
		break
	case "image/png":
		err = png.Encode(file, i)
		break
	}

	if err != nil {
		panic(err)
	}
}
