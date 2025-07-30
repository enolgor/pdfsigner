package signer

import "github.com/enolgor/pdfsigner/signer"

type StoredCertificate struct {
	Passphrase string
	Data       []byte
	Issuer     string
	Subject    string
}

func NewStoredCertificate(data []byte, passphrase string) (*StoredCertificate, error) {
	sc := &StoredCertificate{}
	sc.Data = data
	sc.Passphrase = passphrase
	unlocked, err := sc.Unlock()
	if err != nil {
		return nil, err
	}
	sc.Issuer = unlocked.Certificate.Issuer.CommonName
	sc.Subject = unlocked.Certificate.Subject.CommonName
	return sc, nil
}

func (sc *StoredCertificate) Unlock() (*signer.UnlockedCertificate, error) {
	return signer.UnlockCertificate(sc.Data, sc.Passphrase)
}
