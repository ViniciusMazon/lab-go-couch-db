package main

import (
	"context"
	"fmt"

	_ "github.com/go-kivik/couchdb/v4" // The CouchDB driver
	kivik "github.com/go-kivik/kivik/v4"
)

func main() {

	client, err := kivik.New("couch", "http://docker:123456@localhost:5984/")
	if err != nil {
		panic(err)
	}

	db := client.DB("test")

	doc := map[string]interface{}{
		"_id":  "user",
		"name": "vinicius mazon",
		"age":  27,
	}

	rev, err := db.Put(context.TODO(), "user", doc)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Cow inserted with revision %s\n", rev)
}
