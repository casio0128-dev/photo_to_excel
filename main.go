package main

import (
	"fmt"
	"github.com/joho/godotenv"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"photo2excel/commons"
	excel2 "photo2excel/excel"
	"photo2excel/photos"
)

func init() {
	if err := godotenv.Load(commons.EnvironmentFilePath); err != nil {
		panic(err)
	}
}

func main() {
	photoFiles, err := photos.showFiles()
	if err != nil {
		panic(err)
	}
	fmt.Println(photoFiles)
	images, err := photos.openImages(photoFiles...)
	if err != nil {
		panic(err)
	}

	for _, img := range images {
		excel, err := excel2.openEXCEL(os.Getenv(commons.OutputTarget))
		if err != nil {
			panic(err)
		}
		size := (*img).Bounds().Size()
		excel2.initEXCEL(excel, size.X, size.Y)
		excel2.writeSquare(excel, *img)
	}
}
