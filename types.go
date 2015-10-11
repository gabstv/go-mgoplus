package mgoplus

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

// The DbStats struct contains storage statistics for a given database.
// All size values are in bytes.
//
// Members:
//     DB
//         Contains the name of the database.
//     NCollections
//         Contains a count of the number of collections in that database.
//     NObjects
//         Contains a count of the number of objects (i.e. documents) in
//         the database across all collections.
//     AvgObjSize
//         The average size of each document in bytes. This is the dataSize
//         divided by the number of documents.
//     DataSize
//         The total size in bytes of the data held in this database including
//         the padding factor. The dataSize will not decrease when documents
//         shrink, but will decrease when you remove documents.
//     StorageSize
//         The total amount of space in bytes allocated to collections in this
//         database for document storage. The storageSize does not decrease as
//         you remove or shrink documents.
//     NumExtents
//         Contains a count of the number of extents in the database across all collections.
//     NIndexes
//         Contains a count of the total number of indexes across all collections
//         in the database.
//     IndexSize
//         The total size in bytes of all indexes created on this database.
//     FileSize
//         The total size in bytes of the data files that hold the database.
//         This value includes preallocated space and the padding factor.
//         The value of fileSize only reflects the size of the data files for
//         the database and not the namespace file.
//         Only present when using the mmapv1 storage engine.
//     NsSizeMB
//         The total size of the namespace files (i.e. that end with .ns) for
//         this database. You cannot change the size of the namespace file
//         after creating a database, but you can change the default size for
//         all new namespace files with the nsSize runtime option.
//         Only present when using the mmapv1 storage engine.
//     DataFileVersion
//         See DbStatsDataFileVersion
//     ExtentFreeList
//         See DbStatsExtentFreeList
type DbStats struct {
	DB              string                 `bson:"db"`
	NCollections    int                    `bson:"collections"`
	NObjects        int                    `bson:"objects"`
	AvgObjSize      int64                  `bson:"avgObjSize"`
	DataSize        int64                  `bson:"dataSize"`
	StorageSize     int64                  `bson:"storageSize"`
	NumExtents      int                    `bson:"numExtents"`
	NIndexes        int                    `bson:"indexes"`
	IndexSize       int64                  `bson:"indexSize"`
	FileSize        int64                  `bson:"fileSize"`
	NsSizeMB        int64                  `bson:"nsSizeMB"`
	DataFileVersion DbStatsDataFileVersion `bson:"dataFileVersion"`
	ExtentFreeList  DbStatsExtentFreeList  `bson:"extentFreeList"`
}

// Document that contains information about the on-disk format of the data files
// for the database. Only present when using the mmapv1 storage engine.
type DbStatsDataFileVersion struct {
	Major int `bson:"major"`
	Minor int `bson:"minor"`
}

// Only present when using the mmapv1 storage engine.
type DbStatsExtentFreeList struct {
	Num  int   `bson:"num"`
	Size int64 `bson:"size"`
}
