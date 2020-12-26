package grid

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/disintegration/imaging"
)

// ImageTile Image Grid component
type ImageTile struct {
	ImageFilePath string
	Image         image.Image
	Portrait      bool
	Flipped       bool
}

// ReadFile Get Image from file
func (it *ImageTile) ReadFile() error {
	absPath, err := filepath.Abs(it.ImageFilePath)
	if err != nil {
		return err
	}

	imgFile, err := os.Open(absPath)
	if err != nil {
		return err
	}
	defer imgFile.Close()

	splittedPath := strings.Split(it.ImageFilePath, ".")
	ext := splittedPath[len(splittedPath)-1]

	if ext == "jpg" || ext == "jpeg" {
		it.Image, err = jpeg.Decode(imgFile)
	} else {
		it.Image, err = png.Decode(imgFile)
	}
	it.resize()
	return err
}

func (it *ImageTile) resize() {
	if it.Portrait {
		it.Image = imaging.Resize(it.Image, 270, 480, imaging.Lanczos)
	} else {
		it.Image = imaging.Resize(it.Image, 480, 270, imaging.Lanczos)
	}
}

// Upturn Rotate image to 180°
func (it *ImageTile) Upturn() {
	it.Image = imaging.Rotate180(it.Image)
	it.Flipped = !it.Flipped
}
