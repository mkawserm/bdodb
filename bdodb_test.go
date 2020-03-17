package bdodb

import "testing"
import "github.com/blevesearch/bleve"
import "github.com/blevesearch/bleve/index/upsidedown"

func TestBleveIndex(t *testing.T) {
	index, err := BleveIndex("/tmp/bdodb", bleve.NewIndexMapping(), upsidedown.Name, nil)

	if err != nil {
		t.Errorf("%v", err)
		return
	}

	if index == nil {
		t.Error("Failed to create bleve index")
	}
}
