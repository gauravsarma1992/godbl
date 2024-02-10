package resource

import (
	"time"

	"github.com/gauravsarma1992/gostructs"
)

type (
	ResourceHelpers struct {
		Uuid      uint64
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	Resource struct {
		DecodedResult *gostructs.DecodedResult
	}

	Session struct{}

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

		// Callbacks
		BeforeSave(Resource) (Resource, error)
		AfterSave(Resource) (Resource, error)

		// Transactions
		StartSession() (Session, error)
		EndSession(Session) error
	}
)
