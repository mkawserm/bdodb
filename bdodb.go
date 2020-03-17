package bdodb

import (
	"log"

	"github.com/blevesearch/bleve/registry"
)

const (
	// EngineName is the name of this engine in blevesearch
	EngineName = "bdodb"
)

// init add the engine name to blevesearch
func init() {
	log.SetPrefix("bdodb")
	registry.RegisterKVStore(EngineName, NewStore)
}
