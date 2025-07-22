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
	"image/color"
)

type SignatureConfiguration struct {
	SignaturePageConfiguration
	SignatureContentConfiguration
	SignatureImageConfiguration
	SignatureBorderConfiguration
	SignatureLogoConfiguration
	SignatureTextConfiguration
}

type SignaturePageConfiguration struct {
	Page    int  `json:"page"`
	AddPage *Dim `json:"addPage,omitempty"`
}

type SignatureContentConfiguration struct {
	Title          string     `json:"title"`
	DateFormat     string     `json:"dateFormat"`
	IncludeSubject bool       `json:"includeSubject"`
	IncludeIssuer  bool       `json:"includeIssuer"`
	IncludeDate    bool       `json:"includeDate"`
	SubjectKey     string     `json:"subjectKey"`
	IssuerKey      string     `json:"issuerKey"`
	DateKey        string     `json:"dateKey"`
	ExtraLines     []TextLine `json:"extraLines,omitempty"`
}

type SignatureImageConfiguration struct {
	Dpi             float64    `json:"dpi"`
	BackgroundColor color.RGBA `json:"backgroundColor"`
	WidthPt         float64    `json:"widthPt"`
	HeightPt        float64    `json:"heightPt"`
	PosXPt          float64    `json:"posXPt"`
	PosYPt          float64    `json:"posYPt"`
	PosStrict       bool       `json:"posStrict"`
	Rotate          Rotation   `json:"rotate"`
}

type SignatureBorderConfiguration struct {
	BorderSizePt float64    `json:"borderSizePt"`
	BorderColor  color.RGBA `json:"borderColor"`
}

type SignatureLogoConfiguration struct {
	Logo          *JImage   `json:"logo,omitempty"`
	LogoOpacity   float64   `json:"logoOpacity"`
	LogoGrayScale bool      `json:"logoGrayScale"`
	LogoAlignment Alignment `json:"logoAlignment"`
}

type SignatureTextConfiguration struct {
	EmptyLineAfterTitle bool       `json:"emptyLineAfterTitle"`
	TitleAlignment      Alignment  `json:"titleAlignment"`
	LineAlignment       Alignment  `json:"lineAlignment"`
	KeyAlignment        Alignment  `json:"keyAlignment"`
	ValueAlignment      Alignment  `json:"valueAlignment"`
	TitleFont           string     `json:"titleFont"`
	KeyFont             string     `json:"keyFont"`
	ValueFont           string     `json:"valueFont"`
	TitleColor          color.RGBA `json:"titleColor"`
	KeyColor            color.RGBA `json:"keyColor"`
	ValueColor          color.RGBA `json:"valueColor"`
}

func New(option ...SignatureOption) *SignatureConfiguration {
	config := &SignatureConfiguration{}
	config.Page = 0
	config.AddPage = PaperSize["A4"]
	config.Title = "DIGITALLY SIGNED"
	config.DateFormat = "2006-01-02 15:04:05 -07:00"
	config.IncludeSubject = true
	config.IncludeIssuer = true
	config.IncludeDate = true
	config.SubjectKey = "Subject:"
	config.IssuerKey = "Issuer:"
	config.DateKey = "Date:"
	config.ExtraLines = make([]TextLine, 0)
	config.Dpi = 300
	config.BackgroundColor = color.RGBA{255, 255, 255, 255}
	config.WidthPt = 200
	config.HeightPt = 0
	config.PosXPt = 0
	config.PosYPt = 0
	config.PosStrict = false
	config.Rotate = ROTATE_0
	config.BorderSizePt = 1
	config.BorderColor = color.RGBA{0, 0, 0, 255}
	config.Logo = nil
	config.LogoOpacity = 0.25
	config.LogoGrayScale = false
	config.LogoAlignment = CENTER
	config.EmptyLineAfterTitle = true
	config.TitleAlignment = CENTER
	config.LineAlignment = CENTER
	config.KeyAlignment = LEFT
	config.ValueAlignment = RIGHT
	config.TitleFont = "RobotoMono-Bold"
	config.KeyFont = "RobotoMono-SemiBold"
	config.ValueFont = "RobotoMono-Regular"
	config.TitleColor = color.RGBA{0, 0, 0, 255}
	config.KeyColor = color.RGBA{0, 0, 0, 255}
	config.ValueColor = color.RGBA{0, 0, 0, 255}
	for _, opt := range option {
		opt(config)
	}
	return config
}

func (sc *SignatureConfiguration) With(option ...SignatureOption) *SignatureConfiguration {
	copy := *sc
	for _, opt := range option {
		opt(&copy)
	}
	return &copy
}
