package bdodb

import (
	"github.com/blevesearch/bleve/index/store"
	"github.com/dgraph-io/badger/v2"
	"github.com/dgraph-io/badger/v2/options"
	"os"
	"time"
)

// Store implements bleve store
type Store struct {
	path          string
	db            *badger.DB
	mergeOperator store.MergeOperator
}

// NewStore creates a new store instance
func NewStore(mergeOperator store.MergeOperator, config map[string]interface{}) (store.KVStore, error) {
	path, ok := config["path"].(string)
	if !ok {
		return nil, os.ErrInvalid
	}
	if path == "" {
		return nil, os.ErrInvalid
	}

	opt := badger.DefaultOptions(path)
	opt.ReadOnly = false
	opt.Truncate = true
	opt.TableLoadingMode = options.LoadToRAM
	opt.ValueLogLoadingMode = options.MemoryMap
	opt.Compression = options.Snappy

	//if Dir, ok := config["Dir"].(string); ok {
	//	opt.Dir = Dir
	//}
	//
	//if ValueDir, ok := config["ValueDir"].(string); ok {
	//	opt.ValueDir = ValueDir
	//}

	/* usually modified options */

	// SyncWrites
	if SyncWrites, ok := config["SyncWrites"].(bool); ok {
		opt.SyncWrites = SyncWrites
	}
	// NumVersionsToKeep
	if NumVersionsToKeep, ok := config["NumVersionsToKeep"].(int); ok {
		opt.NumVersionsToKeep = NumVersionsToKeep
	}
	// ReadOnly
	if ReadOnly, ok := config["ReadOnly"].(bool); ok {
		opt.ReadOnly = ReadOnly
	}
	// Truncate
	if Truncate, ok := config["Truncate"].(bool); ok {
		opt.Truncate = Truncate
	}
	// Logger
	if Logger, ok := config["Logger"].(badger.Logger); ok {
		opt.Logger = Logger
	}
	// Compression
	if Compression, ok := config["Compression"].(options.CompressionType); ok {
		opt.Compression = Compression
	}
	// EventLogging
	if EventLogging, ok := config["EventLogging"].(bool); ok {
		opt.EventLogging = EventLogging
	}
	// InMemory
	if InMemory, ok := config["InMemory"].(bool); ok {
		opt.InMemory = InMemory
	}

	/* encryption related options */

	// EncryptionKey
	if EncryptionKey, ok := config["EncryptionKey"].([]byte); ok {
		opt.EncryptionKey = EncryptionKey
	}
	// EncryptionKeyRotationDuration
	if EncryptionKeyRotationDuration, ok := config["EncryptionKeyRotationDuration"].(time.Duration); ok {
		opt.EncryptionKeyRotationDuration = EncryptionKeyRotationDuration
	}

	/* fine tuning options */

	// MaxTableSize
	if MaxTableSize, ok := config["MaxTableSize"].(int64); ok {
		opt.MaxTableSize = MaxTableSize
	}
	// LevelSizeMultiplier
	if LevelSizeMultiplier, ok := config["LevelSizeMultiplier"].(int); ok {
		opt.LevelSizeMultiplier = LevelSizeMultiplier
	}
	// MaxLevels
	if MaxLevels, ok := config["MaxLevels"].(int); ok {
		opt.MaxLevels = MaxLevels
	}
	// ValueThreshold
	if ValueThreshold, ok := config["ValueThreshold"].(int); ok {
		opt.ValueThreshold = ValueThreshold
	}
	// NumMemtables
	if NumMemtables, ok := config["NumMemtables"].(int); ok {
		opt.NumMemtables = NumMemtables
	}
	// BlockSize
	if BlockSize, ok := config["BlockSize"].(int); ok {
		opt.BlockSize = BlockSize
	}
	// BloomFalsePositive
	if BloomFalsePositive, ok := config["BloomFalsePositive"].(float64); ok {
		opt.BloomFalsePositive = BloomFalsePositive
	}
	// KeepL0InMemory
	if KeepL0InMemory, ok := config["KeepL0InMemory"].(bool); ok {
		opt.KeepL0InMemory = KeepL0InMemory
	}
	// MaxCacheSize
	if MaxCacheSize, ok := config["MaxCacheSize"].(int64); ok {
		opt.MaxCacheSize = MaxCacheSize
	}

	// NumLevelZeroTables
	if NumLevelZeroTables, ok := config["NumLevelZeroTables"].(int); ok {
		opt.NumLevelZeroTables = NumLevelZeroTables
	}
	// NumLevelZeroTablesStall
	if NumLevelZeroTablesStall, ok := config["NumLevelZeroTablesStall"].(int); ok {
		opt.NumLevelZeroTablesStall = NumLevelZeroTablesStall
	}

	// LevelOneSize
	if LevelOneSize, ok := config["LevelOneSize"].(int64); ok {
		opt.LevelOneSize = LevelOneSize
	}
	// ValueLogFileSize
	if ValueLogFileSize, ok := config["ValueLogFileSize"].(int64); ok {
		opt.ValueLogFileSize = ValueLogFileSize
	}
	// ValueLogMaxEntries
	if ValueLogMaxEntries, ok := config["ValueLogMaxEntries"].(uint32); ok {
		opt.ValueLogMaxEntries = ValueLogMaxEntries
	}

	// NumCompactors
	if NumCompactors, ok := config["NumCompactors"].(int); ok {
		opt.NumCompactors = NumCompactors
	}
	// CompactL0OnClose
	if CompactL0OnClose, ok := config["CompactL0OnClose"].(bool); ok {
		opt.CompactL0OnClose = CompactL0OnClose
	}
	// LogRotatesToFlush
	if LogRotatesToFlush, ok := config["LogRotatesToFlush"].(int32); ok {
		opt.LogRotatesToFlush = LogRotatesToFlush
	}
	// ZSTDCompressionLevel
	if ZSTDCompressionLevel, ok := config["ZSTDCompressionLevel"].(int); ok {
		opt.ZSTDCompressionLevel = ZSTDCompressionLevel
	}

	// VerifyValueChecksum
	if VerifyValueChecksum, ok := config["VerifyValueChecksum"].(bool); ok {
		opt.VerifyValueChecksum = VerifyValueChecksum
	}
	// ChecksumVerificationMode
	if ChecksumVerificationMode, ok := config["ChecksumVerificationMode"].(options.ChecksumVerificationMode); ok {
		opt.ChecksumVerificationMode = ChecksumVerificationMode
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, os.FileMode(0700))
		if err != nil {
			return nil, err
		}
	}

	db, err := badger.Open(opt)
	if err != nil {
		return nil, err
	}

	rv := Store{
		path:          path,
		db:            db,
		mergeOperator: mergeOperator,
	}
	return &rv, nil
}

// Close cleanup and close the current store
func (store *Store) Close() error {
	return store.db.Close()
}

// Reader initialize a new store.Reader
func (store *Store) Reader() (store.KVReader, error) {
	return &Reader{
		iteratorOptions: badger.DefaultIteratorOptions,
		store:           store,
		txn:             store.db.NewTransaction(false),
	}, nil
}

// Writer initialize a new store.Writer
func (store *Store) Writer() (store.KVWriter, error) {
	return &Writer{
		store: store,
	}, nil
}
