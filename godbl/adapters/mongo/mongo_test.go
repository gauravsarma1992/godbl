package mongo

import (
	"context"
	"testing"

	"github.com/gauravsarma1992/gostructs"
	"github.com/stretchr/testify/assert"
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

func TestMongoInsertAndFineOne(t *testing.T) {
	resource := GetTestResource()
	mongodb, _ := GetTestMongoDb()
	_, err := mongodb.InsertOne(resource)
	findResult, err := mongodb.FindOne(resource)
	assert.Nil(t, err)
	assert.NotNil(t, findResult.Attributes["_id"])
}

func TestMongoInsertAndFineMany(t *testing.T) {
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
