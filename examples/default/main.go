package main

import (
	"time"

	"github.com/enolgor/pdfsigner/examples"
	"github.com/enolgor/pdfsigner/signer"
	"github.com/enolgor/pdfsigner/signer/config"
)

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
	conf := config.New()
	if err := signer.SignVisual(cert, pdf, output, date, metadata, conf); err != nil {
		panic(err)
	}
}
