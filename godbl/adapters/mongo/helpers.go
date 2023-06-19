package mongo

import (
	godblResource "github.com/gauravsarma1992/godbl/godbl/resource"
	"github.com/gauravsarma1992/gostructs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *MongoDb) convertToBson(resource godblResource.Resource) (bsonResource bson.M) {
	bsonResource = bson.M{}
	for resKey, resVal := range resource.Attributes {
		bsonResource[resKey] = resVal
	}
	return
}

func (db *MongoDb) convertPrimaryKeyForDb(resource godblResource.Resource) (bsonResource bson.M) {
	var (
		objectID          primitive.ObjectID
		convertedResource godblResource.Resource
	)
	convertedResource = db.copyResourceToResult(resource)

	objectID, _ = primitive.ObjectIDFromHex(convertedResource.Attributes["id"].(string))
	convertedResource.Attributes = make(map[string]interface{})
	convertedResource.Attributes["_id"] = objectID

	delete(convertedResource.Attributes, "id")

	bsonResource = db.convertToBson(convertedResource)

	return
}

func (db *MongoDb) unconvertPrimaryKeyForDb(resource godblResource.Resource) (result godblResource.Resource) {
	var (
		primaryKey string
	)
	result = db.copyResourceToResult(resource)

	primaryKey = result.Attributes["_id"].(primitive.ObjectID).Hex()
	result.Attributes = make(map[string]interface{})
	result.Attributes["id"] = primaryKey
	delete(result.Attributes, "_id")

	return
}

func (db *MongoDb) convertToResource(bsonResource bson.M, name string) (resource godblResource.Resource) {
	resource = godblResource.Resource(&gostructs.DecodedResult{})
	resource.Name = name
	resource.Attributes = make(map[string]interface{})
	for resKey, resVal := range bsonResource {
		resource.Attributes[resKey] = resVal
	}
	return
}

func (db *MongoDb) copyResourceToResult(resource godblResource.Resource) (result godblResource.Resource) {
	result = godblResource.Resource(&gostructs.DecodedResult{})
	result.Name = resource.Name
	result.Attributes = make(map[string]interface{})
	for resKey, resVal := range resource.Attributes {
		result.Attributes[resKey] = resVal
	}
	return
}
