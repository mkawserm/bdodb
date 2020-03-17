package bdodb

import "os"
import "testing"
import "github.com/blevesearch/bleve"
import "github.com/blevesearch/bleve/index/upsidedown"

func cleanupDb(t *testing.T, path string) {
	err := os.RemoveAll(path)
	if err != nil {
		t.Fatal(err)
	}
}

func TestBleveIndex(t *testing.T) {
	dbPath := "/tmp/bdodb"
	index, err := BleveIndex(dbPath, bleve.NewIndexMapping(), upsidedown.Name, nil)

	if err != nil {
		t.Errorf("%v", err)
		return
	}

	if index == nil {
		t.Error("Failed to create bleve index")
	}

	cleanupDb(t, dbPath)
}

func TestBleveIndexWithEncryptionEnabled(t *testing.T) {
	dbPath := "/tmp/bdodb_encrypted"
	encryptionKey := []byte("67356274875244489356392574264952")

	{
		// create index if not exists
		index, err := BleveIndex(dbPath, bleve.NewIndexMapping(), upsidedown.Name, map[string]interface{}{
			"BdodbConfig": Config{
				EncryptionKey: encryptionKey,
				Logger:        nil,
			},
		})

		if err != nil {
			t.Errorf("%v", err)
			return
		}

		if index == nil {
			t.Error("Failed to create bleve index")
		}

		err = index.Close()
		if err != nil {
			t.Errorf("%v", err)
			return
		}
	}

	// index file should already exists now open it
	{
		index, err := BleveIndex(dbPath, bleve.NewIndexMapping(), upsidedown.Name, map[string]interface{}{
			"BdodbConfig": Config{
				EncryptionKey: encryptionKey,
				Logger:        nil,
			},
		})

		if err != nil {
			t.Errorf("%v", err)
			return
		}

		if index == nil {
			t.Error("Failed to create bleve index")
		}

		err = index.Close()
		if err != nil {
			t.Errorf("%v", err)
			return
		}
	}

	cleanupDb(t, dbPath)
}
