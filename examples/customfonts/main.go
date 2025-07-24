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
	"github.com/enolgor/pdfsigner/signer/fonts"
)

// with command line:
// pdfsigner sign \
//   -c ../cert.p12 -s "bji&M7^#fpEBJAs53JXYf7!3v6MGTucT" \
//   -o output.pdf -p 1 -f -v \
//   -w 300 -x 150 -y 225 \
//   --rs 0 \
//   --nd --ni --nt --sk "Signed by: " \
//   --bc "rgba(230,230,230,255)" \
//   --lf fonts/Corinthia-Bold.ttf --lf fonts/Corinthia-Regular.ttf \
//   --kf "Corinthia-Bold" --vf "Corinthia-Regular" \
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
	fonts.LoadCustomFont("fonts/Corinthia-Bold.ttf")
	fonts.LoadCustomFont("fonts/Corinthia-Regular.ttf")
	conf := config.New(
		config.Page(1),
		config.AddPage(nil),
		config.PosXPt(150),
		config.PosYPt(225),
		config.WidthPt(300),
		config.BorderSizePt(0),
		config.BackgroundColor(color.RGBA{R: 230, G: 230, B: 230, A: 255}),
		config.SubjectKey("Signed by:  "),
		config.IncludeIssuer(false),
		config.IncludeDate(false),
		config.Title(""),
		config.KeyFont("Corinthia-Bold"),
		config.ValueFont("Corinthia-Regular"),
	)
	if err := signer.SignVisual(cert, pdf, output, date, metadata, conf); err != nil {
		panic(err)
	}
}
