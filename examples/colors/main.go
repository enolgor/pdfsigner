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

package main

import (
	"image/color"
	"time"

	"github.com/enolgor/pdfsigner/examples"
	"github.com/enolgor/pdfsigner/signer"
	"github.com/enolgor/pdfsigner/signer/config"
)

// with command line:
// pdfsigner sign \
//   -c ../cert.p12 -s "bji&M7^#fpEBJAs53JXYf7!3v6MGTucT" \
//   -o output.pdf -p 1 -f -v \
//   -w 300 -x 150 -y 500 \
//   --rs 2 \
//   --rc "rgba(255,0,0,255)" \
//   --bc "rgba(0,255,125,125)" \
//   --tc "rgba(255,255,255,255)" \
//   --kc "rgba(0,255,0,255)" \
//   --vc "rgba(255,255,0,255)" \
//   ../test.pdf

func main() {
	cert := examples.GetCertificate()
	pdf := examples.GetTestPDF()
	date := time.Now()
	output := examples.GetOutputPDF()
	defer output.Close()
	metadata := &signer.SignatureMetadata{
		Name:     "John Doe",
		Contact:  "john.doe@example.com",
		Location: "New York, USA",
		Reason:   "Document verification",
	}
	conf := config.New(
		config.AddPage(nil),
		config.Page(1),
		config.AddPage(nil),
		config.PosXPt(150),
		config.PosYPt(500),
		config.WidthPt(300),
		config.BorderSizePt(2),
		config.BorderColor(color.RGBA{R: 255, G: 0, B: 0, A: 255}),
		config.BackgroundColor(color.RGBA{R: 0, G: 255, B: 125, A: 125}),
		config.TitleColor(color.RGBA{R: 255, G: 255, B: 255, A: 255}),
		config.KeyColor(color.RGBA{R: 0, G: 255, B: 0, A: 255}),
		config.ValueColor(color.RGBA{R: 255, G: 255, B: 0, A: 255}),
	)
	if err := signer.SignVisual(cert, pdf, output, date, metadata, conf); err != nil {
		panic(err)
	}
}
