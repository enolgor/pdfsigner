package store

import (
	"errors"
	"fmt"
	"time"

	"github.com/Jipok/go-persist"
)

type DB struct {
	locked   bool
	cipher   Cipher
	store    *persist.Store
	data     map[string]*persist.PersistMap[Encrypted[any]]
	internal *persist.PersistMap[string]
}

func New(path string) (db *DB, err error) {
	store := persist.New()
	if err = store.Open(path); err != nil {
		return
	}
	if err = store.StartAutoShrink(time.Minute, 2); err != nil {
		return
	}
	db = &DB{
		store: store,
		data:  make(map[string]*persist.PersistMap[Encrypted[any]]),
	}
	if db.internal, err = persist.Map[string](db.store, "$"); err != nil {
		return
	}
	salt, ok := db.internal.Get("salt")
	db.locked = ok && salt != ""
	if !db.locked {
		db.cipher = NewNopCipher()
	}
	err = db.store.Shrink()
	return
}

func (db *DB) Unlock(password string) (err error) {
	if !db.locked {
		return
	}
	var salt string
	var ok bool
	var key []byte
	if salt, ok = db.internal.Get("salt"); !ok {
		err = errors.New("salt not found")
		return
	}
	if key, _, err = Derive(password, salt); err != nil {
		return
	}
	if db.cipher, err = NewAESGCMCipher(key); err != nil {
		return
	}
	db.locked = false
	return
}

func (db *DB) ReencryptDB(password string) (err error) {
	if db.locked {
		return errors.New("database is locked")
	}
	var newcipher Cipher
	if password == "" {
		newcipher = NewNopCipher()
		db.internal.Set("salt", "")
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
	}
	for _, pmap := range db.data {
		if err = rangeReencrypt(pmap, db.cipher, newcipher); err != nil {
			return
		}
	}
	db.cipher = newcipher
	err = db.store.Shrink()
	return
}

func rangeReencrypt[T any](pmap *persist.PersistMap[Encrypted[T]], old, new Cipher) (err error) {
	pmap.Range(func(key string, item Encrypted[T]) bool {
		if err = item.Open(old); err != nil {
			return false
		}
		if err = item.Seal(new); err != nil {
			return false
		}
		pmap.Set(key, item)
		return true
	})
	return
}

func Read[T any](db *DB, bucket, key string) (t T, err error) {
	pmap, ok := db.data[bucket]
	if !ok {
		if pmap, err = persist.Map[Encrypted[any]](db.store, bucket); err != nil {
			return
		}
		db.data[bucket] = pmap
	}
	item, found := pmap.Get(key)
	if !found {
		err = fmt.Errorf("item %s/%s not found", bucket, key)
		return
	}
	if err = item.Open(db.cipher); err != nil {
		return
	}
	if t, ok = (*item.Value).(T); !ok {
		err = fmt.Errorf("wrong type for item %s/%s", bucket, key)
	}
	return
}

func Set(db *DB, bucket, key string, val any) (err error) {
	pmap, ok := db.data[bucket]
	if !ok {
		if pmap, err = persist.Map[Encrypted[any]](db.store, bucket); err != nil {
			return
		}
		db.data[bucket] = pmap
	}
	item := NewEncrypted(&val)
	if err = item.Seal(db.cipher); err != nil {
		return
	}
	pmap.Set(key, item)
	return
}

func (db *DB) IsLocked() bool {
	return db.locked
}

func (db *DB) Close() error {
	return db.store.Close()
}
