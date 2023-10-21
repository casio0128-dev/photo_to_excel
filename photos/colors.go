package photos

import (
	"image"
	"image/color"
)

func getColorMatrix(img image.Image) (colorMatrix [][]color.Color) {
	size := img.Bounds().Size()
	for y := 0; y < size.Y; y++ {
		y := y
		var rowColor []color.Color
		for x := 0; x < size.X; x++ {
			rowColor = append(rowColor, getImageColor(img, x, y))
		}
		colorMatrix[y] = append(colorMatrix[y], rowColor...)
	}
	return
}
