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

func getSignData(date time.Time, unlocked *UnlockedCertificate, metadata *SignatureMetadata, appearance *sign.Appearance) *sign.SignData {
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
		TSA:                sign.TSA{},
		RevocationData:     revocation.InfoArchival{},
		RevocationFunction: nil,
	}
	if appearance != nil {
		data.Appearance = *appearance
	}
	return data
}
