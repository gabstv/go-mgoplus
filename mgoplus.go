package mgoplus

import (
	"github.com/jmoiron/jsonq"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
)

type CollectionStats struct {
	// The total size of all indexes
	TotalIndexSize int64 `bson:"totalIndexSize"`
	LastExtentSize int64 `bson:"lastExtentSize"`
	AvgObjSize     int64 `bson:"avgObjSize"`
	StorageSize    int64 `bson:"storageSize"`
	UserFlags      int   `bson:"userFlags"`
	Count          int   `bson:"count"`
	Size           int64 `bson:"size"`
	// The namespace of the current collection, which follows
	// the format [database].[collection].
	NS string `bson:"ns"`
	// Capped is a fixed-sized collection that automatically overwrites
	// its oldest entries when it reaches its maximum size.
	Capped bool `bson:"capped"`
	// The amount of indexes
	NIndexes int `bson:"nindexes"`
	// This field specifies the key and size of every existing
	// index on the collection.
	IndexSizes map[string]int64 `bson:"indexSizes"`
	// I decided to leave the 3.0 deprecated members out:
	//IndexDetails map[string]interface{} `bson:"indexDetails"`
	//Extents map[string]interface{} `bson:"extents"`
	//NumExtents int `bson:"numExtents"`
	//LastExtentSize int64 `bson:"lastExtentSize"`
	//paddingFactor is deprecated
	//ok int `bson:"ok"`
}

// Similar to
// http://docs.mongodb.org/manual/reference/method/db.getCollectionNames/
func GetCollectionNames(db *mgo.Database) ([]string, error) {
	raw := make(map[string]interface{})
	err := db.Run(bson.D{{"listCollections", 1}}, &raw)
	if err != nil {
		return nil, err
	}
	jq := jsonq.NewQuery(raw)
	items, err := jq.ArrayOfObjects("cursor", "firstBatch")
	if err != nil {
		return nil, err
	}
	output := make([]string, len(items))
	for k, v := range items {
		output[k], _ = v["name"].(string)
	}
	return output, nil
}

func GetCollectionStats(db *mgo.Database, collectionName string) (CollectionStats, error) {
	var output CollectionStats
	//raw := make(map[string]interface{})
	err := db.Run(bson.D{
		{"collStats", collectionName},
		{"scale", 1},
		{"verbose", true},
	}, &output)

	if err != nil {
		return output, err
	}
	log.Println(output)
	return output, nil
}
