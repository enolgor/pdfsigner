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
