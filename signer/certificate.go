package signer

import (
	"crypto"
	"crypto/x509"

	"software.sslmate.com/src/go-pkcs12"
)

type UnlockedCertificate struct {
	Signer      crypto.Signer
	Chain       [][]*x509.Certificate
	Certificate *x509.Certificate
}

func UnlockCertificate(pkcs12data []byte, passphrase string) (*UnlockedCertificate, error) {
	key, certificate, parents, err := pkcs12.DecodeChain(pkcs12data, passphrase)
	if err != nil {
		return nil, err
	}
	chain := []*x509.Certificate{certificate}
	chain = append(chain, parents...)
	return &UnlockedCertificate{
		Signer:      key.(crypto.Signer),
		Chain:       [][]*x509.Certificate{chain},
		Certificate: certificate,
	}, nil
}
