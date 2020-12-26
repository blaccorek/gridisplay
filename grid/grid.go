package grid

import (
	"image"
	"image/draw"
)

// Grid List of ImageTile
type Grid struct {
	Tiles    []*ImageTile
	RowNb    int
	ColumnNb int
}

// New Create new Grid
func New(imageFilePaths []string, row int, column int) *Grid {
	g := &Grid{
		RowNb:    row,
		ColumnNb: column,
	}
	for _, path := range imageFilePaths {
		tile := &ImageTile{
			ImageFilePath: path,
			Flipped:       false,
		}
		g.Tiles = append(g.Tiles, tile)
	}
	return g
}

// Merge bnlabla
func (g *Grid) Merge() (*image.NRGBA, error) {
	var canvas *image.NRGBA

	imageBoundX := g.Tiles[0].Image.Bounds().Dx()
	imageBoundY := g.Tiles[0].Image.Bounds().Dy()

	canvasBoundX := g.RowNb * imageBoundX
	canvasBoundY := g.ColumnNb * imageBoundY

	canvasMaxPoint := image.Point{canvasBoundX, canvasBoundY}
	canvasRect := image.Rectangle{image.Point{0, 0}, canvasMaxPoint}
	canvas = image.NewNRGBA(canvasRect)

	// draw grids one by one
	for i, tile := range g.Tiles {
		img := tile.Image
		x := i % g.RowNb
		y := i / g.ColumnNb
		minPoint := image.Point{x * imageBoundX, y * imageBoundY}
		maxPoint := minPoint.Add(image.Point{imageBoundX, imageBoundY})
		nextGridRect := image.Rectangle{minPoint, maxPoint}
		draw.Draw(canvas, nextGridRect, img, image.Point{}, draw.Src)
	}

	return canvas, nil
}

// ExecOnTilePermutation Execute f function on each permutation of tiles
func (g *Grid) ExecOnTilePermutation(f func(*Grid)) {
	perm(g, f, 0)
}

// Permute the values at index i to len(a)-1.
func perm(grid *Grid, f func(*Grid), i int) {
	if i > len(grid.Tiles) {
		f(grid)
	} else {
		perm(grid, f, i+1)
		for j := i + 1; j < len(grid.Tiles); j++ {
			grid.Tiles[i], grid.Tiles[j] = grid.Tiles[j], grid.Tiles[i]
			perm(grid, f, i+1)
			grid.Tiles[i], grid.Tiles[j] = grid.Tiles[j], grid.Tiles[i]
		}
	}
}

func (g *Grid) flipAccordingToMask(mask int) {
	for i := 0; i < len(g.Tiles); i++ {
		currentTile := g.Tiles[i]
		flipped := mask&(1<<i) != 0
		if flipped != currentTile.Flipped {
			currentTile.Upturn()
		}
	}
}

// ExecOnTileFlipCombination Execute f function on each flip combination of tiles
func (g *Grid) ExecOnTileFlipCombination(f func(*Grid)) {
	length := len(g.Tiles)
	for mask := 0; mask < 2<<length; mask++ {
		g.flipAccordingToMask(mask)
		f(g)
	}
}
