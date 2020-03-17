package bdodb

import (
	"github.com/blevesearch/bleve/index/store"
	"github.com/dgraph-io/badger/v2"
)

// Reader implements bleve/Store/Reader interface
type Reader struct {
	iteratorOptions badger.IteratorOptions
	store           *Store
	txn             *badger.Txn
}

// Get fetch the value of the specified key from the store
func (reader *Reader) Get(k []byte) ([]byte, error) {
	item, err := reader.txn.Get(k)
	if err != nil {
		return nil, nil
	}
	return item.ValueCopy(nil)
}

// MultiGet returns multiple values for the specified keys
func (reader *Reader) MultiGet(keys [][]byte) ([][]byte, error) {
	return store.MultiGet(reader, keys)
}

// Iterator initialize a new prefix iterator
func (reader *Reader) PrefixIterator(k []byte) store.KVIterator {
	itr := reader.txn.NewIterator(reader.iteratorOptions)
	rv := Iterator{
		iterator: itr,
		prefix:   k,
	}
	rv.iterator.Seek(k)
	return &rv
}

// RangeIterator initialize a new range iterator
func (reader *Reader) RangeIterator(start, end []byte) store.KVIterator {
	itr := reader.txn.NewIterator(reader.iteratorOptions)
	rv := Iterator{
		iterator: itr,
		start:    start,
		stop:     end,
	}
	rv.iterator.Seek(start)
	return &rv
}

// Close closes the current reader and do some cleanup
func (reader *Reader) Close() error {
	return reader.txn.Commit()
}
