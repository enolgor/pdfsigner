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

package signer

import (
	"bytes"
	"image"
	"image/png"
	"io"
	"text/template"
	"time"

	"github.com/digitorus/pdfsign/sign"
	"github.com/enolgor/pdfsigner/signer/config"
	"github.com/enolgor/pdfsigner/signer/draw"
	"github.com/rotisserie/eris"
)

var drawer draw.Drawer = draw.Rectangle

type TSA = sign.TSA

type SignatureOptions struct {
	TSA TSA
}

type SignatureMetadata struct {
	Name     string
	Location string
	Reason   string
	Contact  string
}

func WithTSA(tsa TSA) func(*SignatureOptions) {
	return func(opts *SignatureOptions) {
		opts.TSA = tsa
	}
}

func Sign(cert *UnlockedCertificate, pdfReader *bytes.Reader, writer io.Writer, date time.Time, metadata *SignatureMetadata, options ...func(*SignatureOptions)) error {
	opts := &SignatureOptions{}
	for _, opt := range options {
		opt(opts)
	}
	return signPdf(pdfReader, writer, getSignData(date, cert, metadata, nil, opts))
}

func SignVisual(cert *UnlockedCertificate, pdfReader *bytes.Reader, writer io.Writer, date time.Time, metadata *SignatureMetadata, conf *config.SignatureConfiguration, options ...func(*SignatureOptions)) (err error) {
	opts := &SignatureOptions{}
	for _, opt := range options {
		opt(opts)
	}
	if conf == nil {
		conf = config.New()
	}
	imageData := new(bytes.Buffer)
	if err = DrawPngImage(imageData, date, cert, conf); err != nil {
		return
	}
	if conf.AddPage != nil {
		if pdfReader, conf.Page, err = addLastPage(pdfReader, conf.AddPage); err != nil {
			return
		}
		if conf.PosXPt == 0 && conf.PosYPt == 0 && !conf.PosStrict {
			conf.PosXPt = (conf.AddPage.Width - conf.WidthPt) / 2
			conf.PosYPt = (conf.AddPage.Height - conf.HeightPt) * 0.9
		}
	}
	return signPdf(pdfReader, writer, getSignData(date, cert, metadata, getAppearance(imageData.Bytes(), conf), opts))
}

func DrawImage(date time.Time, cert *UnlockedCertificate, conf *config.SignatureConfiguration) (image image.Image, err error) {
	if err = parseTextTemplates(cert, date, conf); err != nil {
		err = eris.Wrap(err, "failed to parse text templates")
		return
	}
	text := getTextLines(date, cert, conf)
	if len(text) == 0 && conf.Title == "" {
		err = eris.New("no text to draw")
		return
	}
	if conf.HeightPt == 0 && conf.WidthPt != 0 {
		image, err = drawWithKnownWidth(text, conf)
	} else if conf.HeightPt != 0 && conf.WidthPt == 0 {
		image, err = drawWithKnownHeight(text, conf)
	} else {
		image, err = drawWithKnownWidthAndHeight(text, conf)
	}
	if err != nil {
		err = eris.Wrap(err, "failed to draw image")
		return
	}
	if image, err = drawer.RotateImage(image, conf); err != nil {
		err = eris.Wrap(err, "failed to rotate image")
		return
	}
	if conf.Rotate == config.ROTATE_90 || conf.Rotate == config.ROTATE_270 {
		conf.HeightPt, conf.WidthPt = conf.WidthPt, conf.HeightPt
	}
	return
}

func DrawPngImage(w io.Writer, date time.Time, cert *UnlockedCertificate, conf *config.SignatureConfiguration) (err error) {
	var image image.Image
	if image, err = DrawImage(date, cert, conf); err != nil {
		return
	}
	err = eris.Wrap(png.Encode(w, image), "failed to encode image")
	return
}

func drawWithKnownWidth(text []config.TextLine, conf *config.SignatureConfiguration) (image image.Image, err error) {
	if image, err = drawer.Draw(text, conf); err != nil {
		return
	}
	widthPx, heightPx := image.Bounds().Dx(), image.Bounds().Dy()
	conf.HeightPt = float64(heightPx) * conf.WidthPt / float64(widthPx)
	return
}

func drawWithKnownHeight(text []config.TextLine, conf *config.SignatureConfiguration) (image image.Image, err error) {
	copy := conf.With(config.WidthPt(200), config.HeightPt(0))
	var widthPx, heightPx float64
	if widthPx, heightPx, err = drawer.CalculateExactPixelSize(text, copy); err != nil {
		return
	}
	conf.WidthPt = widthPx * conf.HeightPt / heightPx
	image, err = drawer.Draw(text, conf)
	return
}

func drawWithKnownWidthAndHeight(text []config.TextLine, conf *config.SignatureConfiguration) (image image.Image, err error) {
	image, err = drawer.Draw(text, conf)
	return
}

func CalculateSignatureDim(date time.Time, cert *UnlockedCertificate, conf *config.SignatureConfiguration) (widthPt, heightPt float64, err error) {
	if err = parseTextTemplates(cert, date, conf); err != nil {
		err = eris.Wrap(err, "failed to parse text templates")
		return
	}
	text := getTextLines(date, cert, conf)
	if len(text) == 0 && conf.Title == "" {
		err = eris.New("no text to draw")
		return
	}
	var widthPx, heightPx float64
	if conf.HeightPt == 0 && conf.WidthPt != 0 {
		widthPx, heightPx, err = drawer.CalculateExactPixelSize(text, conf)
		widthPt = conf.WidthPt
		heightPt = heightPx * widthPt / widthPx
	} else if conf.HeightPt != 0 && conf.WidthPt == 0 {
		widthPx, heightPx, err = drawer.CalculateExactPixelSize(text, conf.With(config.WidthPt(200), config.HeightPt(0)))
		heightPt = conf.HeightPt
		widthPt = widthPx * heightPt / heightPx
	} else {
		widthPt = conf.WidthPt
		heightPt = conf.HeightPt
	}
	if conf.Rotate == config.ROTATE_90 || conf.Rotate == config.ROTATE_270 {
		widthPt, heightPt = heightPt, widthPt
	}
	return
}

func getTextLines(date time.Time, cert *UnlockedCertificate, conf *config.SignatureConfiguration) []config.TextLine {
	text := make([]config.TextLine, 0)
	if conf.IncludeSubject {
		text = append(text, config.TextLine{Key: conf.SubjectKey, Value: cert.Certificate.Subject.CommonName})
	}
	if conf.IncludeIssuer {
		text = append(text, config.TextLine{Key: conf.IssuerKey, Value: cert.Certificate.Issuer.CommonName})
	}
	if conf.IncludeDate {
		text = append(text, config.TextLine{Key: conf.DateKey, Value: date.Format(conf.DateFormat)})
	}
	if conf.ExtraLines != nil {
		text = append(text, conf.ExtraLines...)
	}
	return text
}

func parseTextTemplates(cert *UnlockedCertificate, date time.Time, conf *config.SignatureConfiguration) error {
	td := struct {
		Issuer  string
		Subject string
		Date    string
	}{
		Issuer:  cert.Certificate.Issuer.CommonName,
		Subject: cert.Certificate.Subject.CommonName,
		Date:    date.Format(conf.DateFormat),
	}
	buf := new(bytes.Buffer)
	if conf.ExtraLines != nil {
		for i := range conf.ExtraLines {
			tpl, err := template.New("extraLine").Parse(conf.ExtraLines[i].Value)
			if err != nil {
				return eris.Wrap(err, "failed to parse extra line template")
			}
			if err = tpl.Execute(buf, td); err != nil {
				return eris.Wrap(err, "failed to execute extra line template")
			}
			conf.ExtraLines[i].Value = buf.String()
			buf.Reset()
		}
	}
	if conf.Title != "" {
		tpl, err := template.New("title").Parse(conf.Title)
		if err != nil {
			return eris.Wrap(err, "failed to parse title template")
		}
		if err = tpl.Execute(buf, td); err != nil {
			return eris.Wrap(err, "failed to execute title template")
		}
		conf.Title = buf.String()
		buf.Reset()
	}
	return nil
}
