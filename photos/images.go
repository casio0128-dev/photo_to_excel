package photos

import (
	"image"
	"image/color"
	"os"
	"path/filepath"
	"photo2excel/commons"
	"photo2excel/settings"
	"strings"
)

func getImageColor(img image.Image, x, y int) color.Color {
	return img.At(x, y)
}

func withCompatibleExt(name string) bool {
	for _, ext := range commons.CompatibleExtensions {
		if strings.Contains(filepath.Ext(name), ext) {
			return true
		}
	}
	return false
}

func ShowFiles() (map[string][]string, error) {
	setting, err := settings.New()
	if err != nil {
		return nil, err
	}

	photoDir, photoOK := setting.Get(commons.PhotoDir).([]string)
	if !photoOK {
		return nil, commons.FailedTypeCastError{}
	}

	result := make(map[string][]string)
	for _, pDir := range photoDir {
		targetDir := commons.CreateFilePath(pDir)
		dir, err := os.ReadDir(targetDir)
		if err != nil {
			return nil, err
		}

		for _, f := range dir {
			if f.IsDir() || !withCompatibleExt(f.Name()) {
				continue
			}
			result[pDir] = append(result[pDir], commons.CreateFilePath(pDir, f.Name()))
		}
	}
	return result, nil
}

func OpenImages(files ...string) ([]*image.Image, error) {
	var result []*image.Image
	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			return nil, err
		}

		img, _, err := image.Decode(f) //(b)
		if err != nil {
			return nil, err
		}
		result = append(result, &img)
	}
	return result, nil
}
