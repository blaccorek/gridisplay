package grid

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"

	"github.com/disintegration/imaging"
)

type ImageTile struct {
	image    image.Image
	rotation int
}

// ReadFile Get Image from file
func (tile *ImageTile) LoadImageFromFile(filePath string) error {
	absolutePath, err := filepath.Abs(filePath)
	if err != nil {
		return err
	}

	imageFile, err := os.Open(absolutePath)
	if err != nil {
		return err
	}
	defer imageFile.Close()

	fileExtension := filepath.Ext(filePath)

	tile.image, err = decode(imageFile, fileExtension)
	if err != nil {
		return err
	}
	return nil
}

func decode(file *os.File, fileExtension string) (image.Image, error) {
	if isExtensionJPG(fileExtension) {
		return jpeg.Decode(file)
	} else if isExtensionPNG(fileExtension) {
		return png.Decode(file)
	} else {
		return nil, fmt.Errorf("extension not supported: %s", fileExtension)
	}
}

func isExtensionJPG(fileExtension string) bool {
	return fileExtension == "jpg" || fileExtension == "jpeg"
}

func isExtensionPNG(fileExtension string) bool {
	return fileExtension == "png"
}

func (tile *ImageTile) Resize(length int, height int) {
	tile.image = imaging.Resize(tile.image, length, height, imaging.Lanczos)
}

func (tile *ImageTile) SetPortrait(isPortrait bool) {
	if tile.IsPortrait() != isPortrait {
		tile.image = imaging.Rotate90(tile.image)
		tile.rotation = (tile.rotation + 90) % 360
	}
}

func (tile *ImageTile) IsPortrait() bool {
	return tile.GetWidth() > tile.GetHeight()
}

func (tile *ImageTile) GetWidth() int {
	return tile.image.Bounds().Dx()
}

func (tile *ImageTile) GetHeight() int {
	return tile.image.Bounds().Dy()
}

func (tile *ImageTile) Flip() {
	tile.image = imaging.Rotate180(tile.image)
	tile.rotation = (tile.rotation + 180) % 360
}

func (tile *ImageTile) IsFlipped() bool {
	return tile.rotation > 90
}
