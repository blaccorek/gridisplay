package main

import (
	"github.com/jtandria/gridisplay/cli"
	"github.com/jtandria/gridisplay/grid"
	"github.com/jtandria/gridisplay/output"
)

func main() {
	options := cli.ParseOptions()
	currentGrid := grid.LoadFromFiles(options.FilePaths, 2, 2)
	currentGrid.HomogeniseTileOrientation(options.Portrait)
	fileOuput := output.FileOutput{
		ResultFolder: "./output",
	}
	grid.Transform(currentGrid, &fileOuput).GetAllTilePermutationMergedAsFile(1920, 1080)
}
