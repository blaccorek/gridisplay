package grid

import "image"

type Output interface {
	Send(image.Image)
}

type GridTransformer struct {
	output Output
	grid   Grid
}

func Transform(grid Grid, output Output) *GridTransformer {
	return &GridTransformer{
		output: output,
		grid:   grid,
	}
}

func (tilePermuter *GridTransformer) GetAllTilePermutationMergedAsFile(width int, height int) {
	tilePermuter.executeActionOnEachPermutation(func(g *GridTransformer) {
		g.ExecuteActionOnTileFlipCombinations(func(g *Grid) {
			img := g.MergeTilesToSingleImage(width, height)
			tilePermuter.output.Send(img)
		})
	})
}

func (transformer *GridTransformer) ExecuteActionOnTileFlipCombinations(action func(*Grid)) {
	length := transformer.grid.Size()
	for mask := 0; mask < 2<<length; mask++ {
		transformer.flipAccordingToMask(mask)
		action(&transformer.grid)
	}
}

func (transformer *GridTransformer) flipAccordingToMask(mask int) {
	for i := 0; i < transformer.grid.Size(); i++ {
		currentTile := transformer.grid.tiles[i]
		if flippedTileRequired(mask, i) != currentTile.IsFlipped() {
			currentTile.Flip()
		}
	}
}

func flippedTileRequired(mask int, tileIndex int) bool {
	return mask&(1<<tileIndex) != 0
}

func (tilePermuter *GridTransformer) executeActionOnEachPermutation(action func(*GridTransformer)) {
	tilePermuter.recursivelyExecuteOnPermutation(action, 0)
}

func (transformer *GridTransformer) recursivelyExecuteOnPermutation(action func(*GridTransformer), tileIndex int) {
	if tileIndex > transformer.grid.Size() {
		action(transformer)
	} else {
		transformer.recursivelyExecuteOnPermutation(action, tileIndex+1)
		transformer.permuteThenExecuteOnTile(action, tileIndex)
	}
}

func (transformer *GridTransformer) permuteThenExecuteOnTile(action func(*GridTransformer), currentTileIndex int) {
	for nextTileIndex := currentTileIndex + 1; nextTileIndex < transformer.grid.Size(); nextTileIndex++ {
		transformer.grid.Permute(currentTileIndex, nextTileIndex)
		transformer.recursivelyExecuteOnPermutation(action, currentTileIndex+1)
		transformer.grid.Permute(currentTileIndex, nextTileIndex)
	}
}
