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

package examples

import (
	"bytes"
	_ "embed" // for embedding the certificate
	"io"
	"os"

	"github.com/enolgor/pdfsigner/signer"
)

const passphrase = "bji&M7^#fpEBJAs53JXYf7!3v6MGTucT"

//go:embed cert.p12
var certificateData []byte

//go:embed test.pdf
var testPDFData []byte

func GetCertificate() *signer.UnlockedCertificate {
	cert, err := signer.UnlockCertificate(certificateData, passphrase)
	if err != nil {
		panic(err)
	}
	return cert
}

func GetTestPDF() *bytes.Reader {
	return bytes.NewReader(testPDFData)
}

func GetOutputPDF() io.WriteCloser {
	if out, err := os.Create("output.pdf"); err != nil {
		panic(err)
	} else {
		return out
	}
}
