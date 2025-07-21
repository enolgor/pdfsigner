package draw

import (
	"image"

	"github.com/enolgor/pdfsigner/signer/config"
)

type Drawer interface {
	Draw(text []config.TextLine, conf *config.SignatureConfiguration) (image.Image, error)
	CalculateExactPixelSize(text []config.TextLine, conf *config.SignatureConfiguration) (float64, float64, error)
	RotateImage(img image.Image, conf *config.SignatureConfiguration) (image.Image, error)
}

func PtsToPixels(pts float64, dpi float64) float64 {
	return pts * dpi / 72
}

func PixelsToPts(pix float64, dpi float64) float64 {
	return pix * 72 / dpi
}

var lineKey = func(line config.TextLine) string { return line.Key }
var lineValue = func(line config.TextLine) string { return line.Value }

var Rectangle Drawer = &rect{ypadpt: 5, xpadpt: 5, vspacept: 3}
