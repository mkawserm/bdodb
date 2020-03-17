package bdodb_test

import (
	"fmt"
	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/index/upsidedown"
	"github.com/mkawserm/bdodb"
)

func ExampleBleveIndex() {
	// create or open bleveIndex
	index, err := bdodb.BleveIndex("/tmp/bdodb_example", bleve.NewIndexMapping(), upsidedown.Name, nil)

	if err != nil {
		panic(err)
	}

	message := struct {
		Id   string
		Body string
	}{
		Id:   "custom_id",
		Body: "bleve indexing with badger using bdodb",
	}

	// index message data
	if val, _ := index.Document(message.Id); val == nil {
		if err := index.Index(message.Id, message); err != nil {
			panic(err)
		}
	}

	// search for some text
	query := bleve.NewQueryStringQuery("bdodb")
	search := bleve.NewSearchRequest(query)
	if searchResults, err := index.Search(search); err == nil {
		fmt.Println(searchResults)
	} else {
		fmt.Println(err)
	}
}
