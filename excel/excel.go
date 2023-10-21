package excel

import (
	"github.com/xuri/excelize/v2"
	"image"
	"os"
)

func writeSquare(target *excelize.File, img image.Image) error {
	imgSize := img.Bounds().Size()
	for yIndex := 0; yIndex < imgSize.Y; yIndex++ {
		for xIndex := 0; xIndex < imgSize.X; xIndex++ {
			target.SetCellStyle(target.GetSheetName(0), string(yIndex), string(xIndex), 0)
		}
	}
	return nil
}

func initEXCEL(excel *excelize.File, maxX, maxY int) error {
	colName, err := excelize.ColumnNumberToName(maxY)
	if err != nil {
		return err
	}
	excel.SetColWidth(excel.GetSheetList()[0], "A", colName, 2)
	excel.SetRowHeight(excel.GetSheetList()[0], maxX, 2)

	return nil
}

func openEXCEL(path string) (*excelize.File, error) {
	if fi, err := os.Stat(path); err != nil {
		if fi.IsDir() {
			return nil, err
		}
		if os.IsNotExist(err) {
			file := excelize.NewFile()
			if err := file.SaveAs(path, excelize.Options{}); err != nil {
				return nil, err
			} else {
				return file, nil
			}
		} else if os.IsExist(err) {
			file, err := excelize.OpenFile(path, excelize.Options{})
			if err != nil {
				return nil, err
			}
			if err := file.SaveAs(path, excelize.Options{}); err != nil {
				return nil, err
			}
			return file, nil
		}
	}
	return nil, nil
}
