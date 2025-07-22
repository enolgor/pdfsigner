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
