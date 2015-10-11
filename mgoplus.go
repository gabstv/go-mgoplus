// Package mgoplus implements admin functions to use with labix.org
// mongodb driver.
package mgoplus

import (
	"github.com/jmoiron/jsonq"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
)

// CollectionStats contains mostly size measurements (in bytes) of a collection.
// Use GetCollectionStats to retrieve this struct.
//
// Members:
//     TotalIndexSize
//         The size of all indexes, combined.
//     LastExtentSize
//         The size of the last extent allocated. The scale argument affects this value.
//         Only present when using the mmapv1 storage engine.
//     AvgObjSize
//         Average size of a document/object.
//     StorageSize
//         The total amount of storage allocated to this collection for document storage.
//     UserFlags
//         A number that indicates the user-set flags on the collection.
//         userFlags only appears when using the mmapv1 storage engine.
//     Count
//         The number of objects or documents in this collection.
//     Size
//         The total size in memory of all records in a collection. This value does not
//         include the record header, which is 16 bytes per record, but does include the
//         record’s padding. Additionally size does not include the size of any indexes
//         associated with the collection, which the TotalIndexSize field reports.
//     NS
//         The namespace of the current collection, which follows the format
//         [database].[collection].
//     Capped
//         Capped is a fixed-sized collection that automatically overwrites
//         its oldest entries when it reaches its maximum size.
//     NIndexes
//         The amount of indexes (including _id_).
//     IndexSizes
//         This field specifies the key and size of every existing index on the collection.
type CollectionStats struct {
	TotalIndexSize int64            `bson:"totalIndexSize"`
	LastExtentSize int64            `bson:"lastExtentSize"`
	AvgObjSize     int64            `bson:"avgObjSize"`
	StorageSize    int64            `bson:"storageSize"`
	UserFlags      int              `bson:"userFlags"`
	Count          int              `bson:"count"`
	Size           int64            `bson:"size"`
	NS             string           `bson:"ns"`
	Capped         bool             `bson:"capped"`
	NIndexes       int              `bson:"nindexes"`
	IndexSizes     map[string]int64 `bson:"indexSizes"`
	// I decided to leave the 3.0 deprecated members out:
	//IndexDetails map[string]interface{} `bson:"indexDetails"`
	//Extents map[string]interface{} `bson:"extents"`
	//NumExtents int `bson:"numExtents"`
	//LastExtentSize int64 `bson:"lastExtentSize"`
	//paddingFactor is deprecated
	//ok int `bson:"ok"`
}

// Returns a slice containing the names of all collections in the current database.
//
// Based on
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

// Returns a variety of storage statistics for a given collection.
//
// Based on
// http://docs.mongodb.org/manual/reference/method/db.collection.stats/
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
