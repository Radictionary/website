package database

import (
	"fmt"
	"log"

	"github.com/dgraph-io/badger/v4"
)

func CallDatabase() *badger.DB {
	db, err := badger.Open(badger.DefaultOptions("/tmp/badgerv4"))
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func ViewDatabase(db *badger.DB) {
	_ = db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
		  item := it.Item()
		  k := item.Key()
		  err := item.Value(func(v []byte) error {
			fmt.Printf("key=%s, value=%s\n", k, v)
			return nil
		  })
		  if err != nil {
			return err
		  }
		}
		return nil
	  })
	  
}
