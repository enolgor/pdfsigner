package db

import (
	"fmt"
	"os"

	"github.com/dgraph-io/badger/v4"
	"github.com/pkg/errors"
)

type DB struct {
	path string
	open bool
	*badger.DB
}

func New(path string) *DB {
	return &DB{
		path: path,
		open: false,
	}
}

func (db *DB) Open(key []byte) (err error) {
	opts := badger.DefaultOptions(db.path)
	if key != nil {
		opts = opts.WithEncryptionKey(key)
	}
	db.DB, err = badger.Open(opts)
	return
}

func (db *DB) ReencryptDB(key []byte) (err error) {
	fmt.Println("reencrypting")
	opts := badger.DefaultOptions(db.path + ".tmp")
	if key != nil {
		opts = opts.WithEncryptionKey(key)
	}
	fmt.Println("1")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("db is closed", db.DB.IsClosed())
	if err = db.StreamDB(opts); err != nil {
		fmt.Println("fuck this shit")
		return errors.Wrap(err, "error streaming db")
	}
	fmt.Println("2")
	// if err = db.Close(); err != nil {

	// 	return
	// }
	fmt.Println("3")
	if err = os.Rename(db.path+".tmp", db.path); err != nil {
		return
	}
	return db.Open(key)
}

func (db *DB) ReadTest() (value string) {
	var val []byte
	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("test"))
		if err != nil {
			return err
		}
		val, err = item.ValueCopy(nil)
		return err
	})
	if err != nil {
		val = []byte{}
	}
	return string(val)
}

func (db *DB) SetTest(val string) {
	db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte("test"), []byte(val))
	})
}

func (db *DB) IsClosed() bool {
	return db.DB == nil || db.DB.IsClosed()
}

func (db *DB) Close() error {
	return db.DB.Close()
}
