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
//   --cert ../cert.p12 --passphrase "bji&M7^#fpEBJAs53JXYf7!3v6MGTucT" \
//   --out output.pdf --page 1 --force --visible \
//   --width 750 --rotate 270 --xpos 10 --ypos 50 \
//   --datetime-format "02 Jan 2006 at 15:04:05" --border-size 0 \
//   --no-date --no-issuer --no-subject --no-empty-line-after-title \
//   --background-color "transparent" --title-color "rgba(150, 150, 150, 130)" \
//   --title "Signed by {{.Subject}} on {{.Date}}. Certificate issued by {{.Issuer}}." \
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
		config.WidthPt(750),
		config.Rotate(config.ROTATE_270),
		config.Title("Signed by {{.Subject}} on {{.Date}}. Certificate issued by {{.Issuer}}."),
		config.DateFormat("02 Jan 2006 at 15:04:05"),
		config.IncludeDate(false),
		config.IncludeIssuer(false),
		config.IncludeSubject(false),
		config.BackgroundColor(color.RGBA{R: 255, G: 255, B: 255, A: 0}),
		config.BorderSizePt(0),
		config.TitleColor(color.RGBA{R: 150, G: 150, B: 150, A: 130}),
		config.Page(1),
		config.AddPage(nil),
		config.PosXPt(10),
		config.PosYPt(50),
		config.EmptyLineAfterTitle(false),
	)
	if err := signer.SignVisual(cert, pdf, output, date, metadata, conf); err != nil {
		panic(err)
	}
}
