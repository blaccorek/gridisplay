package grid

import (
	"image"
	"image/draw"
	"log"
)

// Grid List of ImageTile
type Grid struct {
	tiles       []*ImageTile
	rowCount    int
	columnCount int
}

// LoadFromFiles Create new Grid
func LoadFromFiles(imageFilePaths []string, row int, column int) Grid {
	grid := Grid{
		rowCount:    row,
		columnCount: column,
	}
	for _, path := range imageFilePaths {
		tile := &ImageTile{
			rotation: 0,
		}
		err := tile.LoadImageFromFile(path)
		if err != nil {
			log.Printf("Cannot load image from file nammed %s", path)
		} else {
			grid.AddTile(tile)
		}
	}
	return grid
}

func (grid *Grid) Size() int {
	return len(grid.tiles)
}

func (grid *Grid) MaxSize() int {
	return grid.rowCount * grid.columnCount
}

func (grid *Grid) AddTile(tile *ImageTile) {
	grid.tiles = append(grid.tiles, tile)
}

func (grid *Grid) MergeTilesToSingleImage(width int, height int) *image.NRGBA {
	canvas := createCanvas(width, height)
	for _, tile := range grid.tiles {
		tile.Resize(width/grid.rowCount, height/grid.columnCount)
		position := &image.Point{
			X: grid.rowCount * tile.GetWidth(),
			Y: grid.columnCount * tile.GetHeight(),
		}
		drawTileInCanvas(position, tile, canvas)
	}
	return canvas
}

func createCanvas(width int, height int) *image.NRGBA {
	canvasMaxPoint := image.Point{width, height}
	canvasRect := image.Rectangle{image.Point{0, 0}, canvasMaxPoint}
	return image.NewNRGBA(canvasRect)
}

func drawTileInCanvas(position *image.Point, tile *ImageTile, canvas *image.NRGBA) {
	rectangleToDraw := getTileAreaInCanvas(position, tile)
	draw.Draw(canvas, rectangleToDraw, tile.image, image.Point{}, draw.Src)
}

func getTileAreaInCanvas(position *image.Point, tile *ImageTile) image.Rectangle {
	tileWidth := tile.GetWidth()
	tileHeight := tile.GetHeight()
	startPointInCanvas := image.Point{position.X * tileWidth, position.Y * tileHeight}
	endPointInCanvas := startPointInCanvas.Add(image.Point{tileWidth, tileHeight})
	return image.Rectangle{startPointInCanvas, endPointInCanvas}
}

func (grid *Grid) HomogeniseTileOrientation(isPortrait bool) {
	for _, tile := range grid.tiles {
		tile.SetPortrait(isPortrait)
	}
}

func (grid *Grid) Permute(src int, dest int) {
	grid.tiles[src], grid.tiles[dest] = grid.tiles[dest], grid.tiles[src]
}
