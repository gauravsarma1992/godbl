package resource

import (
	"github.com/gauravsarma1992/gostructs"
)

type (
	Resource *gostructs.DecodedResult

	Db interface {
		Setup() error
		InsertOne(Resource) (Resource, error)
		FindOne(Resource) (Resource, error)
		DeleteOne(Resource) (Resource, error)
		UpdateOne(Resource) (Resource, error)

		FindMany(Resource) ([]Resource, error)
		InsertMany([]Resource) ([]Resource, error)
		UpdateMany([]Resource) ([]Resource, error)
		DeleteMany([]Resource) ([]Resource, error)
	}
)
