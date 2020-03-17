package bdodb

import (
	"fmt"
	"log"

	"github.com/blevesearch/bleve/index/store"
)

// Writer implements bleve store Writer interface
type Writer struct {
	store *Store
}

// NewBatch implements NewBatch
func (writer *Writer) NewBatch() store.KVBatch {
	return store.NewEmulatedBatch(writer.store.mergeOperator)
}

// NewBatchEx implements bleve NewBatchEx
func (writer *Writer) NewBatchEx(options store.KVBatchOptions) ([]byte, store.KVBatch, error) {
	return make([]byte, options.TotalBytes), writer.NewBatch(), nil
}

// ExecuteBatch implements bleve ExecuteBatch
func (writer *Writer) ExecuteBatch(batch store.KVBatch) (err error) {
	emulatedBatch, ok := batch.(*store.EmulatedBatch)
	if !ok {
		return fmt.Errorf("wrong type of batch")
	}

	txn := writer.store.db.NewTransaction(true)

	defer (func() {
		err := txn.Commit()
		log.Printf("ERROR: %v\n", err)
	})()

	for k, mergeOps := range emulatedBatch.Merger.Merges {
		kb := []byte(k)
		item, err := txn.Get(kb)
		existingVal := []byte{}
		if err == nil {
			existingVal, _ = item.ValueCopy(nil)
		}
		mergedVal, fullMergeOk := writer.store.mergeOperator.FullMerge(kb, existingVal, mergeOps)
		if !fullMergeOk {
			return fmt.Errorf("merge operator returned failure")
		}
		err = txn.Set(kb, mergedVal)
		if err != nil {
			return err
		}
	}

	for _, op := range emulatedBatch.Ops {
		if op.V != nil {
			err = txn.Set(op.K, op.V)
			if err != nil {
				return err
			}
		} else {
			err = txn.Delete(op.K)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Close closes the current writer
func (writer *Writer) Close() error {
	return nil
}
