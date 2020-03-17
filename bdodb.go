package bdodb

import (
	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/index/scorch"
	"github.com/blevesearch/bleve/mapping"
)

const (
	// EngineName is the name of this engine in blevesearch
	EngineName = "bdodb"
)

// BleveIndex a helper function that open (creates if not exists a new) bleve index
func BleveIndex(path string, mapping mapping.IndexMapping) (bleve.Index, error) {
	index, err := bleve.NewUsing(path, mapping, scorch.Name, EngineName, nil)

	if err != nil && err == bleve.ErrorIndexPathExists {
		index, err = bleve.Open(path)
	}

	return index, err
}
