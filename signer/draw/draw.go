// MIT License
//
// Copyright (c) 2025 @enolgor
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

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
