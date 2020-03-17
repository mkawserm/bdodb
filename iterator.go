package bdodb

import (
	"bytes"

	"github.com/dgraph-io/badger/v2"
)

// Iterator implements bleve store iterator
type Iterator struct {
	iterator *badger.Iterator
	prefix   []byte

	start []byte
	stop  []byte
}

// Seek advance the iterator to the specified key
func (i *Iterator) Seek(key []byte) {
	if len(i.prefix) > 0 {
		if bytes.Compare(key, i.prefix) < 0 {
			i.iterator.Seek(i.prefix)
			return
		}
	} else if len(i.start) > 0 {
		if bytes.Compare(key, i.start) < 0 {
			i.iterator.Seek(i.start)
			return
		}
	}

	i.iterator.Seek(key)
}

// Next advance the iterator to the next step
func (i *Iterator) Next() {
	i.iterator.Next()
}

// Current returns the key & value of the current step
func (i *Iterator) Current() ([]byte, []byte, bool) {
	if i.Valid() {
		return i.Key(), i.Value(), true
	}
	return nil, nil, false
}

// Key return the key of the current step
func (i *Iterator) Key() []byte {
	return i.iterator.Item().KeyCopy(nil)
}

// Value returns the value of the current step
func (i *Iterator) Value() []byte {
	v, _ := i.iterator.Item().ValueCopy(nil)
	return v
}

// Valid whether the current iterator step is valid or not
func (i *Iterator) Valid() bool {
	// for prefix iterator
	if len(i.prefix) > 0 {
		return i.iterator.ValidForPrefix(i.prefix)
	}

	// for range based iterator
	if !i.iterator.Valid() {
		return false
	}
	if i.stop == nil || len(i.stop) == 0 {
		return true
	}
	if bytes.Compare(i.stop, i.iterator.Item().Key()) <= 0 {
		return false
	}
	return true
}

// Close closes the current iterator and commit its transaction
func (i *Iterator) Close() error {
	i.iterator.Close()
	return nil
}
