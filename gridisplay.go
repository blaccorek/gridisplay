package main

import (
	"crypto/md5"
	"encoding/base32"
	"flag"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"

	"github.com/jtandria/gridisplay/grid"
)

type arrayFlags []string

func (i *arrayFlags) String() string {
	return "my string representation"
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func generateFileName(img *image.NRGBA) string {
	bytes := md5.Sum([]byte(fmt.Sprintf("%#v", *img)))
	return base32.StdEncoding.EncodeToString(bytes[:])
}

func main() {
	// Get file paths
	var filePaths arrayFlags
	flag.Var(&filePaths, "f", "Image file path (must provide 4 files)")
	isPortrait := flag.Bool("portrait", false, "Use portrait orientation images")
	flag.Parse()
	log.Printf("length : %d; val: %#v", len(filePaths), filePaths)

	if len(filePaths) != 4 {
		flag.Usage()
	}
	g := grid.New(filePaths, 2, 2)
	for _, tile := range g.Tiles {
		tile.Portrait = *isPortrait
		err := tile.ReadFile()
		if err != nil {
			log.Fatal("cannot read file: " + err.Error())
		}
	}

	g.ExecOnTilePermutation(func(g *grid.Grid) {
		g.ExecOnTileFlipCombination(func(g *grid.Grid) {
			img, err := g.Merge()
			if err != nil {
				return
			}
			// save the output to png
			fileName := generateFileName(img)
			folderName := "results/" + fileName[:1]
			log.Printf("Generated filename is " + fileName)
			filePath := folderName + "/" + fileName + ".png"

			// Create your result folder if required
			if _, err := os.Stat(filePath); os.IsNotExist(err) {
				os.MkdirAll(folderName, 0700)
			}
			file, err := os.Create(filePath)
			err = png.Encode(file, img)
			if err != nil {
				log.Fatal("Cannot create new file:" + err.Error())
			}
		})
	})
}
