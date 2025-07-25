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

package config

import (
	"image"
	"image/color"
)

type SignatureOption func(*SignatureConfiguration)

func Page(page int) SignatureOption {
	return func(config *SignatureConfiguration) {
		config.SignaturePageConfiguration.Page = page
	}
}

func AddPage(dim *Dim) SignatureOption {
	return func(config *SignatureConfiguration) {
		config.SignaturePageConfiguration.AddPage = dim
	}
}

func Title(title string) SignatureOption {
	return func(config *SignatureConfiguration) {
		config.SignatureContentConfiguration.Title = title
	}
}

func DateFormat(format string) SignatureOption {
	return func(config *SignatureConfiguration) {
		config.SignatureContentConfiguration.DateFormat = format
	}
}

func IncludeSubject(include bool) SignatureOption {
	return func(config *SignatureConfiguration) {
		config.SignatureContentConfiguration.IncludeSubject = include
	}
}

func IncludeIssuer(include bool) SignatureOption {
	return func(config *SignatureConfiguration) {
		config.SignatureContentConfiguration.IncludeIssuer = include
	}
}

func IncludeDate(include bool) SignatureOption {
	return func(config *SignatureConfiguration) {
		config.SignatureContentConfiguration.IncludeDate = include
	}
}

func SubjectKey(key string) SignatureOption {
	return func(config *SignatureConfiguration) {
		config.SignatureContentConfiguration.SubjectKey = key
	}
}

func IssuerKey(key string) SignatureOption {
	return func(config *SignatureConfiguration) {
		config.SignatureContentConfiguration.IssuerKey = key
	}
}

func DateKey(key string) SignatureOption {
	return func(config *SignatureConfiguration) {
		config.SignatureContentConfiguration.DateKey = key
	}
}

func ExtraLine(key string, value string) SignatureOption {
	return func(config *SignatureConfiguration) {
		if config.SignatureContentConfiguration.ExtraLines == nil {
			config.SignatureContentConfiguration.ExtraLines = []TextLine{}
		}
		config.SignatureContentConfiguration.ExtraLines = append(config.ExtraLines, TextLine{Key: key, Value: value})
	}
}

func Dpi(dpi float64) SignatureOption {
	return func(config *SignatureConfiguration) {
		config.SignatureImageConfiguration.Dpi = dpi
	}
}

func BackgroundColor(color color.RGBA) SignatureOption {
	return func(config *SignatureConfiguration) {
		config.SignatureImageConfiguration.BackgroundColor = color
	}
}

func WidthPt(widthPt float64) SignatureOption {
	return func(config *SignatureConfiguration) {
		config.SignatureImageConfiguration.WidthPt = widthPt
	}
}

func HeightPt(heightPt float64) SignatureOption {
	return func(config *SignatureConfiguration) {
		config.SignatureImageConfiguration.HeightPt = heightPt
	}
}

func PosYPt(posYpt float64) SignatureOption {
	return func(config *SignatureConfiguration) {
		config.SignatureImageConfiguration.PosYPt = posYpt
	}
}

func PosXPt(posXPt float64) SignatureOption {
	return func(config *SignatureConfiguration) {
		config.SignatureImageConfiguration.PosXPt = posXPt
	}
}

func PosStrict(posStrict bool) SignatureOption {
	return func(config *SignatureConfiguration) {
		config.SignatureImageConfiguration.PosStrict = posStrict
	}
}

func Rotate(rotation Rotation) SignatureOption {
	return func(config *SignatureConfiguration) {
		config.SignatureImageConfiguration.Rotate = rotation
	}
}

func BorderSizePt(size float64) SignatureOption {
	return func(config *SignatureConfiguration) {
		config.SignatureBorderConfiguration.BorderSizePt = size
	}
}

func BorderColor(color color.RGBA) SignatureOption {
	return func(config *SignatureConfiguration) {
		config.SignatureBorderConfiguration.BorderColor = color
	}
}

func Logo(logo image.Image) SignatureOption {
	return func(config *SignatureConfiguration) {
		config.SignatureLogoConfiguration.Logo = &JImage{Image: logo}
	}
}

func LogoOpacity(opacity float64) SignatureOption {
	return func(config *SignatureConfiguration) {
		config.SignatureLogoConfiguration.LogoOpacity = opacity
	}
}

func LogoGrayscale(grayScale bool) SignatureOption {
	return func(config *SignatureConfiguration) {
		config.SignatureLogoConfiguration.LogoGrayScale = grayScale
	}
}

func LogoAlignment(position Alignment) SignatureOption {
	return func(config *SignatureConfiguration) {
		config.SignatureLogoConfiguration.LogoAlignment = position
	}
}

func EmptyLineAfterTitle(empty bool) SignatureOption {
	return func(config *SignatureConfiguration) {
		config.SignatureTextConfiguration.EmptyLineAfterTitle = empty
	}
}

func TitleAlignment(position Alignment) SignatureOption {
	return func(config *SignatureConfiguration) {
		config.SignatureTextConfiguration.TitleAlignment = position
	}
}

func LineAlignment(position Alignment) SignatureOption {
	return func(config *SignatureConfiguration) {
		config.SignatureTextConfiguration.LineAlignment = position
	}
}

func KeyAlignment(position Alignment) SignatureOption {
	return func(config *SignatureConfiguration) {
		config.SignatureTextConfiguration.KeyAlignment = position
	}
}

func ValueAlignment(position Alignment) SignatureOption {
	return func(config *SignatureConfiguration) {
		config.SignatureTextConfiguration.ValueAlignment = position
	}
}

func TitleFont(font string) SignatureOption {
	return func(config *SignatureConfiguration) {
		config.SignatureTextConfiguration.TitleFont = font
	}
}

func KeyFont(font string) SignatureOption {
	return func(config *SignatureConfiguration) {
		config.SignatureTextConfiguration.KeyFont = font
	}
}

func ValueFont(font string) SignatureOption {
	return func(config *SignatureConfiguration) {
		config.SignatureTextConfiguration.ValueFont = font
	}
}

func TitleColor(color color.RGBA) SignatureOption {
	return func(config *SignatureConfiguration) {
		config.SignatureTextConfiguration.TitleColor = color
	}
}

func KeyColor(color color.RGBA) SignatureOption {
	return func(config *SignatureConfiguration) {
		config.SignatureTextConfiguration.KeyColor = color
	}
}

func ValueColor(color color.RGBA) SignatureOption {
	return func(config *SignatureConfiguration) {
		config.SignatureTextConfiguration.ValueColor = color
	}
}
