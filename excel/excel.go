package excel

import (
	"github.com/xuri/excelize/v2"
	"image"
	"os"
	"photo2excel/commons"
)

//func setupt() {
//	excelize.Style{
//		Border:        nil,
//		Fill:          excelize.Fill{},
//		Font:          nil,
//		Alignment:     nil,
//		Protection:    nil,
//		NumFmt:        0,
//		DecimalPlaces: nil,
//		CustomNumFmt:  nil,
//		NegRed:        false,
//	}
//}

func WriteSquare(target *excelize.File, img image.Image) error {
	imgSize := img.Bounds().Size()
	for yIndex := 0; yIndex < imgSize.Y; yIndex++ {
		for xIndex := 0; xIndex < imgSize.X; xIndex++ {
			target.SetCellStyle(target.GetSheetName(0), string(yIndex), string(xIndex))
		}
	}
	return nil
}

//
//func InitEXCEL(excel *excelize.File, maxX, maxY int) error {
//	colName, err := excelize.ColumnNumberToName(maxY)
//	if err != nil {
//		return err
//	}
//	excel.SetColWidth(excel.GetSheetList()[0], "A", colName, 200)
//	excel.SetRowHeight(excel.GetSheetList()[0], maxX, 200)
//
//	return nil
//}

func OpenEXCEL(path string) (*excelize.File, error) {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			file := excelize.NewFile(excelize.Options{})
			if err := file.SaveAs(path + commons.ExcelExtension); err != nil {
				return nil, err
			} else {
				return file, nil
			}
		} else if os.IsExist(err) {
			file, err := excelize.OpenFile(path)
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
