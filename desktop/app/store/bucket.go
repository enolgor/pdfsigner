package store

import (
	"github.com/Jipok/go-persist"
	"github.com/rotisserie/eris"
)

var ErrItemNotExist error = eris.New("item does not exist")
var ErrBucketLocked error = eris.New("bucket is locked")

type Entry[T any] struct {
	Key   string
	Value T
}

type EncryptedBucket[T any] interface {
	Bucket[T]
	Unlock(cipher Cipher)
	Reencrypt(newcipher Cipher) error
	Locked() bool
}

type Bucket[T any] interface {
	Get(key string) (T, error)
	Set(key string, value T) error
	Delete(key string) error
	All() ([]Entry[T], error)
	Keys() []string
}

type bucket[T any] struct {
	name   string
	pmap   *persist.PersistMap[T]
	seqmap *persist.PersistMap[Sequence]
}

func NewBucket[T any](store *persist.Store, name string, seqmap *persist.PersistMap[Sequence]) (Bucket[T], error) {
	pmap, err := persist.Map[T](store, name)
	if err != nil {
		return nil, err
	}
	_, found := seqmap.Get(name)
	if !found {
		seqmap.Set(name, make(Sequence, 0))
	}
	return &bucket[T]{name, pmap, seqmap}, nil
}

func (b *bucket[T]) Get(key string) (T, error) {
	value, ok := b.pmap.Get(key)
	if !ok {
		return value, eris.Wrapf(ErrItemNotExist, "item %s/%s not found", b.name, key)
	}
	return value, nil
}

func (b *bucket[T]) Set(key string, value T) error {
	b.pmap.Set(key, value)
	sequence, _ := b.seqmap.Get(b.name)
	sequence.Add(key)
	b.seqmap.Set(b.name, sequence)
	return nil
}

func (b *bucket[T]) Delete(key string) error {
	if ok := b.pmap.Delete(key); !ok {
		return eris.Wrapf(ErrItemNotExist, "item %s/%s not found", b.name, key)
	}
	sequence, _ := b.seqmap.Get(b.name)
	sequence.Remove(key)
	b.seqmap.Set(b.name, sequence)
	return nil
}

func (b *bucket[T]) Keys() []string {
	sequence, _ := b.seqmap.Get(b.name)
	return sequence
}

func (b *bucket[T]) All() ([]Entry[T], error) {
	all := []Entry[T]{}
	b.pmap.Range(func(key string, t T) bool {
		all = append(all, Entry[T]{key, t})
		return true
	})
	sequence, _ := b.seqmap.Get(b.name)
	sort(sequence, all)
	return all, nil
}

type encryptedBucket[T any] struct {
	name   string
	bucket Bucket[Encrypted[T]]
	cipher Cipher
	locked bool
}

func NewEncryptedBucket[T any](store *persist.Store, name string, seqmap *persist.PersistMap[Sequence]) (EncryptedBucket[T], error) {
	bucket, err := NewBucket[Encrypted[T]](store, name, seqmap)
	if err != nil {
		return nil, err
	}
	return &encryptedBucket[T]{name, bucket, nil, true}, nil
}

func (b *encryptedBucket[T]) Get(key string) (t T, err error) {
	if b.locked {
		err = ErrBucketLocked
		return
	}
	var item Encrypted[T]
	if item, err = b.bucket.Get(key); err != nil {
		return
	}
	if err = item.Open(b.cipher); err != nil {
		err = eris.Wrapf(err, "item %s/%s could not be decrypted", b.name, key)
		return
	}
	t = *item.Value
	return
}

func (b *encryptedBucket[T]) Set(key string, value T) error {
	if b.locked {
		return ErrBucketLocked
	}
	item := NewEncrypted(&value)
	if err := item.Seal(b.cipher); err != nil {
		return eris.Wrapf(err, "item %s/%s could not be encrypted", b.name, key)
	}
	return b.bucket.Set(key, item)
}

func (b *encryptedBucket[T]) Delete(key string) error {
	if b.locked {
		return ErrBucketLocked
	}
	return b.bucket.Delete(key)
}

func (b *encryptedBucket[T]) All() ([]Entry[T], error) {
	if b.locked {
		return nil, ErrBucketLocked
	}
	allEnc, err := b.bucket.All()
	if err != nil {
		return nil, err
	}
	all := make([]Entry[T], len(allEnc))
	for i := range allEnc {
		if err = allEnc[i].Value.Open(b.cipher); err != nil {
			err = eris.Wrapf(err, "item %s/%s could not be decrypted", b.name, allEnc[i].Key)
			return nil, err
		}
		all[i] = Entry[T]{allEnc[i].Key, *allEnc[i].Value.Value}
	}
	return all, nil
}

func (b *encryptedBucket[T]) Unlock(cipher Cipher) {
	b.cipher = cipher
	b.locked = false
}

func (b *encryptedBucket[T]) Reencrypt(newcipher Cipher) error {
	if b.locked {
		return ErrBucketLocked
	}
	all, err := b.All()
	if err != nil {
		return err
	}
	b.cipher = newcipher
	for _, entry := range all {
		if err = b.Set(entry.Key, entry.Value); err != nil {
			return err
		}
	}
	return nil
}

func (b *encryptedBucket[T]) Locked() bool {
	return b.locked
}

func (b *encryptedBucket[T]) Keys() []string {
	return b.bucket.Keys()
}
