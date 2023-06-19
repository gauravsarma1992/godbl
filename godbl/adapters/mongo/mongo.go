package mongo

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gauravsarma1992/godbl/godbl"
	"github.com/gauravsarma1992/gostructs"
)

type (
	MongoDb struct {
		ctx    context.Context
		config *MongoConfig
		db     *mongo.Database
	}

	MongoConfig struct {
		Username   string `json:"username"`
		Password   string `json:"password"`
		Host       string `json:"host"`
		Port       string `json:"port"`
		DbName     string `json:"db_name"`
		ReplicaSet string `json:"replica_set"`
	}
)

func NewMongoDb(ctx context.Context, config *MongoConfig) (db *MongoDb, err error) {
	db = &MongoDb{
		ctx:    ctx,
		config: config,
	}
	if config == nil {
		if db.config, err = NewMongoConfig("config.json"); err != nil {
			return
		}
	}
	if err = db.Setup(); err != nil {
		return
	}
	return
}

func NewMongoConfig(fileName string) (mongoConfig *MongoConfig, err error) {
	var (
		fileB []byte
	)
	mongoConfig = &MongoConfig{}
	if fileB, err = ioutil.ReadFile(fileName); err != nil {
		log.Println("Error reading config file", err)
	}
	if err = json.Unmarshal(fileB, mongoConfig); err != nil {
		log.Println("Error unmarshalling config file", err)
	}
	if mongoConfig.Host == "" {
		mongoConfig.Host = "localhost"
	}
	if mongoConfig.Port == "" {
		mongoConfig.Port = "27017"
	}
	if mongoConfig.DbName == "" {
		mongoConfig.DbName = "dev"
	}
	err = nil

	return
}

func (mongoConfig *MongoConfig) GetUrl() (url string) {
	userCreds := ""
	if mongoConfig.Username != "" && mongoConfig.Password != "" {
		userCreds = fmt.Sprintf("%s:%s@", mongoConfig.Username, mongoConfig.Password)
	}
	url = fmt.Sprintf("mongodb://%s%s:%s/%s?replicaSet=%s",
		userCreds,
		mongoConfig.Host,
		mongoConfig.Port,
		mongoConfig.DbName,
		mongoConfig.ReplicaSet,
	)
	return
}

func (db *MongoDb) connectToDb() (err error) {
	var (
		client *mongo.Client
	)
	client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(db.config.GetUrl()))
	if err != nil {
		return
	}
	if err = client.Ping(context.Background(), nil); err != nil {
		return
	}
	db.db = client.Database(db.config.DbName)
	return
}

func (db *MongoDb) Setup() (err error) {
	if err = db.connectToDb(); err != nil {
		return
	}
	return
}

func (db *MongoDb) convertToBson(resource godbl.Resource) (bsonResource bson.M) {
	bsonResource = bson.M{}
	for resKey, resVal := range resource.Attributes {
		bsonResource[resKey] = resVal
	}
	return
}

func (db *MongoDb) convertToResource(bsonResource bson.M, name string) (resource godbl.Resource) {
	resource = godbl.Resource(&gostructs.DecodedResult{})
	resource.Name = name
	resource.Attributes = make(map[string]interface{})
	for resKey, resVal := range bsonResource {
		resource.Attributes[resKey] = resVal
	}
	return
}

func (db *MongoDb) copyResourceToResult(resource godbl.Resource) (result godbl.Resource) {
	result = godbl.Resource(&gostructs.DecodedResult{})
	result.Name = resource.Name
	result.Attributes = make(map[string]interface{})
	for resKey, resVal := range resource.Attributes {
		result.Attributes[resKey] = resVal
	}
	return
}

func (db *MongoDb) InsertOne(resource godbl.Resource) (result godbl.Resource, err error) {
	var (
		insertResult *mongo.InsertOneResult
	)
	currCollection := db.db.Collection(resource.Name)
	if insertResult, err = currCollection.InsertOne(db.ctx, db.convertToBson(resource)); err != nil {
		return
	}
	result = &gostructs.DecodedResult{}
	result = db.copyResourceToResult(resource)
	result.Attributes["_id"] = insertResult.InsertedID
	return
}

func (db *MongoDb) FindOne(resource godbl.Resource) (result godbl.Resource, err error) {
	result = godbl.Resource(&gostructs.DecodedResult{})
	singleResult := &mongo.SingleResult{}
	currCollection := db.db.Collection(resource.Name)
	opts := options.FindOne()
	if singleResult = currCollection.FindOne(context.TODO(), db.convertToBson(resource), opts); singleResult.Err() != nil {
		err = singleResult.Err()
		return
	}
	if err = singleResult.Decode(&result.Attributes); err != nil {
		return
	}
	return
}

func (db *MongoDb) DeleteOne(resource godbl.Resource) (result godbl.Resource, err error) {
	currCollection := db.db.Collection(resource.Name)
	if _, err = currCollection.DeleteOne(db.ctx, db.convertToBson(resource)); err != nil {
		return
	}
	return
}

func (db *MongoDb) UpdateOne(resource godbl.Resource) (result godbl.Resource, err error) {
	currCollection := db.db.Collection(resource.Name)
	if _, err = currCollection.UpdateOne(db.ctx, bson.M{"_id": resource.Attributes["_id"]}, db.convertToBson(resource)); err != nil {
		return
	}
	return
}

func (db *MongoDb) FindMany(resource godbl.Resource) (results []godbl.Resource, err error) {
	var (
		cursor      *mongo.Cursor
		bsonResults []bson.M
	)
	currCollection := db.db.Collection(resource.Name)
	if cursor, err = currCollection.Find(db.ctx, db.convertToBson(resource)); err != nil {
		return
	}
	if err = cursor.All(db.ctx, &bsonResults); err != nil {
		return
	}
	for _, bsonResult := range bsonResults {
		results = append(results, db.convertToResource(bsonResult, resource.Name))
	}

	return
}

func (db *MongoDb) InsertMany(resources []godbl.Resource) (results []godbl.Resource, err error) {
	return
}

func (db *MongoDb) UpdateMany(resources []godbl.Resource) (results []godbl.Resource, err error) {
	return
}

func (db *MongoDb) DeleteMany(resources []godbl.Resource) (results []godbl.Resource, err error) {
	return
}
