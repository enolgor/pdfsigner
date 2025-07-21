package draw

import (
	"errors"
	"math"

	"image"
	"image/color"
	"image/draw"

	xdraw "golang.org/x/image/draw"
)

func redrawLogo(img image.Image, parentBounds image.Rectangle, opacity float64, grayscale bool) (image.Image, error) {
	if opacity < 0 || opacity > 1 {
		return nil, errors.New("opacity must be between 0 and 1")
	}
	origBounds := img.Bounds()
	origWidth := origBounds.Dx()
	origHeight := origBounds.Dy()
	newWidth, newHeight := fitInside(parentBounds.Dx(), parentBounds.Dy(), origWidth, origHeight)
	resized := image.NewNRGBA(image.Rect(0, 0, newWidth, newHeight))
	xdraw.CatmullRom.Scale(resized, resized.Bounds(), img, origBounds, draw.Over, nil)
	if grayscale {
		resized = toGrayscaleWithAlpha(resized)
	}
	final := image.NewNRGBA(resized.Bounds())
	alphaMask := image.NewUniform(color.Alpha{A: uint8(opacity * 255)})
	draw.DrawMask(final, final.Bounds(), resized, image.Point{}, alphaMask, image.Point{}, draw.Over)
	return final, nil
}

func toGrayscaleWithAlpha(src *image.NRGBA) *image.NRGBA {
	bounds := src.Bounds()
	dst := image.NewNRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := src.NRGBAAt(x, y)
			gray := uint8(0.299*float64(c.R) + 0.587*float64(c.G) + 0.114*float64(c.B))
			dst.SetNRGBA(x, y, color.NRGBA{R: gray, G: gray, B: gray, A: c.A})
		}
	}
	return dst
}

func fitInside(parentW, parentH, childW, childH int) (newW, newH int) {
	scaleX := float64(parentW) / float64(childW)
	scaleY := float64(parentH) / float64(childH)

	scale := math.Min(scaleX, scaleY)

	newW = int(float64(childW) * scale)
	newH = int(float64(childH) * scale)
	return
}
