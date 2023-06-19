package godbl

import (
	"context"

	"github.com/gauravsarma1992/godbl/godbl/resource"
)

type (
	Godbl struct {
		Db resource.Db
	}

	CompliantDb struct{}
)

func NewDb(ctx context.Context, db resource.Db) (dbl *Godbl, err error) {
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

func (cdb *CompliantDb) InsertOne(resource resource.Resource) (result resource.Resource, err error) {
	return
}

func (cdb *CompliantDb) FindOne(resource resource.Resource) (result resource.Resource, err error) {
	return
}

func (cdb *CompliantDb) DeleteOne(resource resource.Resource) (result resource.Resource, err error) {
	return
}

func (cdb *CompliantDb) UpdateOne(resource resource.Resource) (result resource.Resource, err error) {
	return
}

func (cdb *CompliantDb) FindMany(resource resource.Resource) (results []resource.Resource, err error) {
	return
}

func (cdb *CompliantDb) InsertMany(resources []resource.Resource) (results []resource.Resource, err error) {
	return
}

func (cdb *CompliantDb) UpdateMany(resources []resource.Resource) (results []resource.Resource, err error) {
	return
}

func (cdb *CompliantDb) DeleteMany(resources []resource.Resource) (results []resource.Resource, err error) {
	return
}
