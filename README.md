bdodb 
-------------------------------------------------
[![GoDoc](https://godoc.org/github.com/mkawserm/bdodb?status.svg)](https://godoc.org/github.com/mkawserm/bdodb)
[![Build Status](https://travis-ci.com/mkawserm/bdodb.svg?branch=master)](https://travis-ci.com/mkawserm/bdodb)
[![Go Report Card](https://goreportcard.com/badge/github.com/mkawserm/bdodb)](https://goreportcard.com/report/github.com/mkawserm/bdodb)
[![Coverage Status](https://coveralls.io/repos/github/mkawserm/bdodb/badge.svg?branch=master)](https://coveralls.io/github/mkawserm/bdodb?branch=master)

-------------------------------------------------

Package bdodb is a badger based backend database for bleve
with all the customization options provided by the badger

# Usage

> `➜ go get github.com/mkawserm/bdodb`

```go

package main

import (
    "fmt"
    "github.com/blevesearch/bleve"
    "github.com/blevesearch/bleve/index/upsidedown"
    "github.com/mkawserm/bdodb"
)

func main() { 
    // create or open bleveIndex
    index, err := bdodb.BleveIndex("/tmp/bdodb", bleve.NewIndexMapping(), upsidedown.Name, nil)
    
    message := struct{
        Id   string
        Body string
    }{
        Id:   "custom_id", 
        Body: "bleve indexing with badger using bdodb",
    }
    
    // index message data 
    err = index.Index(message.Id, message)
    if err !=nil {
        panic(err)
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
```