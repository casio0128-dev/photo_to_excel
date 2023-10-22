package main

import (
	"fmt"
	"github.com/joho/godotenv"
	_ "image/jpeg"
	_ "image/png"
	"photo2excel/commons"
	excel2 "photo2excel/excel"
	"photo2excel/photos"
	"photo2excel/settings"
)

func init() {
	if err := godotenv.Load(commons.EnvironmentFilePath); err != nil {
		panic(err)
	}
}

func main() {
	photoFiles, err := photos.ShowFiles()
	if err != nil {
		panic(err)
	}
	fmt.Println(photoFiles)
	for k, v := range photoFiles {
		images, err := photos.OpenImages(v...)
		if err != nil {
			panic(err)
		}

		setting, err := settings.New()
		if err != nil {
			panic(err)
		}

		for _, img := range images {
			dist := setting.Get(commons.OutputDir).(string)
			_, err := excel2.OpenEXCEL(commons.CreateFilePath(dist, k))
			if err != nil {
				panic(err)
			}
			//size := (*img).Bounds().Size()
			//excel2.InitEXCEL(excel, size.X, size.Y)
			excel2.WriteSquare(excel, *img)
		}
	}
}
