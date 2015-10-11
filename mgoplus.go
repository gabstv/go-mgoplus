// Package mgoplus implements admin functions to use with labix.org
// mongodb driver.
package mgoplus

import (
	"github.com/jmoiron/jsonq"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

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
	err := db.Run(bson.D{
		{"collStats", collectionName},
		{"scale", 1},
		{"verbose", true},
	}, &output)

	if err != nil {
		return output, err
	}
	return output, nil
}

// Returns stats of all collections in the database.
func GetAllCollectionStats(db *mgo.Database) ([]CollectionStats, error) {
	dbnames, err := GetCollectionNames(db)
	if err != nil {
		return nil, err
	}
	clstats := make([]CollectionStats, len(dbnames))
	for k := range dbnames {
		clstats[k], err = GetCollectionStats(db, dbnames[k])
		if err != nil {
			return clstats, err
		}
	}
	return clstats, nil
}

// Returns a variety of database statistics.
//
// Based on
// http://docs.mongodb.org/manual/reference/command/dbStats/#dbcmd.dbStats
func GetDbStats(db *mgo.Database) (DbStats, error) {
	var output DbStats
	err := db.Run(bson.D{
		{"dbStats", 1},
		{"scale", 1},
	}, &output)

	if err != nil {
		return output, err
	}
	return output, nil
}
