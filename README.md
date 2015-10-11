# mgoplus
### Admin functions for the labix.org mongodb driver.

[![GoDoc](https://godoc.org/github.com/gabstv/go-mgoplus?status.svg)](https://godoc.org/github.com/gabstv/go-mgoplus)

#### Example
```go
package main

import(
	"github.com/gabstv/go-mgoplus"
	"labix.org/v2/mgo"
	"log"
)

func main() {
	// connect to mongodb
	sess, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatalln(err)
	}

	db := sess.DB("my_collection")

	// retrieve all collections
	colls, err := mgoplus.GetCollectionNames(db)
	if err != nil {
		log.Fatalln(err)
	}

	for _, v := range colls {
		log.Println(v)
	}

	// retrieve information of a collection
	info, err := mgoplus.GetCollectionStats(db, "users")
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Collection '%s' size is %v bytes.\n", "users", info.Size)
}
```