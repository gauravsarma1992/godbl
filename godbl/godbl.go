package godbl

import (
	"context"

	"github.com/gauravsarma1992/gostructs"
)

type (
	Godbl struct {
		Db Db
	}

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

	CompliantDb struct{}
)

func NewDb(ctx context.Context, db Db) (dbl *Godbl, err error) {
	dbl = &Godbl{
		Db: db,
	}
	if err = db.Setup(); err != nil {
		return
	}
	return
}

func (cdb *CompliantDb) Setup() error {
	return nil
}

func (cdb *CompliantDb) InsertOne(resource Resource) (result Resource, err error) {
	return
}

func (cdb *CompliantDb) FindOne(resource Resource) (result Resource, err error) {
	return
}

func (cdb *CompliantDb) DeleteOne(resource Resource) (result Resource, err error) {
	return
}

func (cdb *CompliantDb) UpdateOne(resource Resource) (result Resource, err error) {
	return
}

func (cdb *CompliantDb) FindMany(resource Resource) (results []Resource, err error) {
	return
}

func (cdb *CompliantDb) InsertMany(resources []Resource) (results []Resource, err error) {
	return
}

func (cdb *CompliantDb) UpdateMany(resources []Resource) (results []Resource, err error) {
	return
}

func (cdb *CompliantDb) DeleteMany(resources []Resource) (results []Resource, err error) {
	return
}
