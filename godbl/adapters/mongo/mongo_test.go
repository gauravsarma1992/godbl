package mongo

import (
	"context"
	"testing"

	"github.com/gauravsarma1992/gostructs"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetTestResource() (resource *gostructs.DecodedResult) {
	resource = &gostructs.DecodedResult{
		Name: "test_a",
		Attributes: map[string]interface{}{
			"a": "b",
		},
	}
	return
}

func GetTestMongoDb() (db *MongoDb, err error) {
	db, err = NewMongoDb(context.Background(), nil)
	db.db.Collection("test_a").Drop(context.Background())
	return
}

func TestMongoNew(t *testing.T) {
	var (
		db  *MongoDb
		err error
	)

	if db, err = NewMongoDb(context.Background(), nil); err != nil {
		t.Error(err)
	}
	assert.NotNil(t, db)
}

func TestMongoInsert(t *testing.T) {
	resource := GetTestResource()
	mongodb, err := GetTestMongoDb()
	result, err := mongodb.InsertOne(resource)
	assert.Nil(t, err)
	assert.NotEmptyf(t, result.Attributes["_id"], "ID is empty")
}

func TestMongoInsertAndFindOne(t *testing.T) {
	resource := GetTestResource()
	mongodb, _ := GetTestMongoDb()
	resource, err := mongodb.InsertOne(resource)

	findResource := &gostructs.DecodedResult{
		Attributes: make(map[string]interface{}),
	}
	findResource.Name = resource.Name
	findResource.Attributes["id"] = resource.Attributes["_id"].(primitive.ObjectID).Hex()

	findResult, err := mongodb.FindOne(findResource)

	assert.Nil(t, err)
	assert.NotNil(t, findResult.Attributes["_id"])
}

func TestMongoInsertAndFindMany(t *testing.T) {
	resource := GetTestResource()
	mongodb, _ := GetTestMongoDb()
	for idx := 0; idx < 10; idx++ {
		resource.Attributes["a"] = "b"
		mongodb.InsertOne(resource)
	}
	findResult, err := mongodb.FindMany(resource)
	assert.Nil(t, err)
	assert.Equal(t, 10, len(findResult))
}

func TestMongoUpdate(t *testing.T) {
	resource := GetTestResource()

	mongodb, _ := GetTestMongoDb()
	resource, _ = mongodb.InsertOne(resource)

	updateResource := &gostructs.DecodedResult{
		Attributes: make(map[string]interface{}),
	}
	updateResource.Name = resource.Name
	updateResource.Attributes["id"] = resource.Attributes["_id"].(primitive.ObjectID).Hex()

	updateResource.Attributes["a"] = "c"
	_, err := mongodb.UpdateOne(updateResource)

	assert.Nil(t, err)

	findResource := &gostructs.DecodedResult{
		Name:       resource.Name,
		Attributes: make(map[string]interface{}),
	}
	findResource.Attributes["id"] = resource.Attributes["_id"].(primitive.ObjectID).Hex()
	findResource, _ = mongodb.FindOne(findResource)
	assert.Equal(t, "c", findResource.Attributes["a"])
}

func TestMongoDelete(t *testing.T) {
	resource := GetTestResource()

	mongodb, _ := GetTestMongoDb()
	resource, _ = mongodb.InsertOne(resource)

	deleteResource := &gostructs.DecodedResult{
		Attributes: make(map[string]interface{}),
	}
	deleteResource.Name = resource.Name
	deleteResource.Attributes["id"] = resource.Attributes["_id"].(primitive.ObjectID).Hex()

	_, err := mongodb.DeleteOne(deleteResource)
	assert.Nil(t, err)

	findResult, _ := mongodb.FindMany(resource)
	assert.Equal(t, 0, len(findResult))
}
