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
	"crypto"
	"io"
	"time"

	"github.com/digitorus/pdf"
	"github.com/digitorus/pdfsign/revocation"
	"github.com/digitorus/pdfsign/sign"
	"github.com/enolgor/pdfsigner/signer/config"
)

func signPdf(pdfReader *bytes.Reader, writer io.Writer, signData *sign.SignData) error {
	size := pdfReader.Size()
	rdr, err := pdf.NewReader(pdfReader, size)
	if err != nil {
		return err
	}
	return sign.Sign(pdfReader, writer, rdr, size, *signData)
}

func getAppearance(image []byte, conf *config.SignatureConfiguration) *sign.Appearance {
	return &sign.Appearance{
		Visible:          true,
		Page:             uint32(conf.Page),
		LowerLeftX:       conf.PosXPt,
		LowerLeftY:       conf.PosYPt,
		UpperRightX:      conf.PosXPt + conf.WidthPt,
		UpperRightY:      conf.PosYPt + conf.HeightPt,
		Image:            image,
		ImageAsWatermark: false,
	}
}

func getSignData(date time.Time, unlocked *UnlockedCertificate, metadata *SignatureMetadata, appearance *sign.Appearance, options *SignatureOptions) *sign.SignData {
	data := &sign.SignData{
		Signature: sign.SignDataSignature{
			Info: sign.SignDataSignatureInfo{
				Name:        metadata.Name,
				Location:    metadata.Location,
				Reason:      metadata.Reason,
				ContactInfo: metadata.Contact,
				Date:        date,
			},
			CertType:   sign.ApprovalSignature,
			DocMDPPerm: sign.DoNotAllowAnyChangesPerms,
		},
		Signer:             unlocked.Signer,
		DigestAlgorithm:    crypto.SHA256,
		Certificate:        unlocked.Certificate,
		CertificateChains:  unlocked.Chain,
		TSA:                options.TSA,
		RevocationData:     revocation.InfoArchival{},
		RevocationFunction: nil,
	}
	if appearance != nil {
		data.Appearance = *appearance
	}
	return data
}
