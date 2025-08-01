package store

import (
	"time"

	"github.com/Jipok/go-persist"
	"github.com/enolgor/pdfsigner/desktop/app/certs"
	"github.com/rotisserie/eris"
)

var ErrDatabaseLocked error = eris.New("database is locked")
var ErrInvalidPassword error = eris.New("invalid password")

type DB struct {
	store      *persist.Store
	sequence   *persist.PersistMap[Sequence]
	internal   Bucket[string]
	canary     EncryptedBucket[[]byte]
	flags      Bucket[bool]
	testBucket EncryptedBucket[string]
	certs      EncryptedBucket[certs.StoredCertificate]
	protected  bool

	locked bool
}

func New(path string) (db *DB, err error) {
	store := persist.New()
	if err = store.Open(path); err != nil {
		return
	}
	if err = store.StartAutoShrink(time.Minute, 2); err != nil {
		return
	}
	db = &DB{store: store}
	if db.sequence, err = persist.Map[Sequence](store, "$sequence"); err != nil {
		return
	}
	if db.internal, err = NewBucket[string](store, "$internal", db.sequence); err != nil {
		return
	}
	if db.flags, err = NewBucket[bool](store, "flags", db.sequence); err != nil {
		return
	}
	if db.testBucket, err = NewEncryptedBucket[string](store, "data", db.sequence); err != nil {
		return
	}
	if db.certs, err = NewEncryptedBucket[certs.StoredCertificate](store, "certs", db.sequence); err != nil {
		return
	}
	if db.canary, err = NewEncryptedBucket[[]byte](store, "$canary", db.sequence); err != nil {
		return
	}
	salt, _ := db.internal.Get("salt")
	db.locked = salt != ""
	db.protected = salt != ""
	if !db.locked {
		db.Unlock("")
	}
	err = db.store.Shrink()
	return
}

func (db *DB) Flags() Bucket[bool] {
	return db.flags
}

func (db *DB) TestBucket() Bucket[string] {
	return db.testBucket
}

func (db *DB) Certs() Bucket[certs.StoredCertificate] {
	return db.certs
}

func (db *DB) Unlock(password string) (err error) {
	var cipher Cipher
	if password != "" {
		var salt string
		var key []byte
		if salt, err = db.internal.Get("salt"); err != nil {
			return
		}
		if key, _, err = Derive(password, salt); err != nil {
			return
		}
		if cipher, err = NewAESGCMCipher(key); err != nil {
			return
		}
		db.protected = true
	} else {
		cipher = NewNopCipher()
		db.protected = false
	}
	db.testBucket.Unlock(cipher)
	db.certs.Unlock(cipher)
	db.canary.Unlock(cipher)
	if _, err := db.canary.Get("canary"); err != nil {
		if IsNotExist(err) {
			rng, _ := randomBytes(16)
			if err = db.canary.Set("canary", rng); err != nil {
				return err
			}
		} else {
			return ErrInvalidPassword
		}
	}
	db.locked = false
	return
}

func (db *DB) Reencrypt(password string) (err error) {
	if db.locked {
		return ErrDatabaseLocked
	}
	var newcipher Cipher
	if password == "" {
		newcipher = NewNopCipher()
		db.internal.Set("salt", "")
		db.protected = false
	} else {
		var key []byte
		var salt string
		if key, salt, err = Derive(password, ""); err != nil {
			return
		}
		db.internal.Set("salt", salt)
		if newcipher, err = NewAESGCMCipher(key); err != nil {
			return
		}
		db.protected = true
	}
	if err = db.testBucket.Reencrypt(newcipher); err != nil {
		return
	}
	if err = db.certs.Reencrypt(newcipher); err != nil {
		return
	}
	if err = db.canary.Reencrypt(newcipher); err != nil {
		return
	}
	err = db.store.Shrink()
	return
}

func (db *DB) IsLocked() bool {
	return db.locked
}

func (db *DB) Close() error {
	return db.store.Close()
}

func (db *DB) IsProtected() bool {
	return db.protected
}

func IsNotExist(err error) bool {
	return eris.Is(err, ErrItemNotExist)
}

func IsInvalidPassword(err error) bool {
	return eris.Is(err, ErrInvalidPassword)
}
