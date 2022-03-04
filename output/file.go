package output

import (
	"crypto/md5"
	"encoding/base32"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
)

type FileOutput struct {
	ResultFolder string
}

func (output *FileOutput) Send(img image.Image) {
	filePath := output.buildResultFilePath(img)
	output.ensureResultFolderExists(filePath)
	file, err := os.Create(filePath)
	if err != nil {
		log.Printf("Cannot create new file %s; caused by: %s", filePath, err.Error())
	}
	err = png.Encode(file, img)
	if err != nil {
		log.Printf("Cannot write result in file %s; caused by: %s", filePath, err.Error())
	}
}

func (output *FileOutput) buildResultFilePath(img image.Image) string {
	fileName := generateFileName(img)
	folderName := output.ResultFolder + "/" + fileName[:1]
	filePath := folderName + "/" + fileName + ".png"
	return filePath
}

func generateFileName(img image.Image) string {
	bytes := md5.Sum([]byte(fmt.Sprintf("%#v", img)))
	return base32.StdEncoding.EncodeToString(bytes[:])
}

func (output *FileOutput) ensureResultFolderExists(fileName string) {
	folderName := output.ResultFolder + fileName[:1]
	if _, err := os.Stat(folderName); os.IsNotExist(err) {
		os.MkdirAll(folderName, 0600)
	}
}
