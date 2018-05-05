package db

import (
	"github.com/globalsign/mgo"
	"github.com/moxuz/price-protection-notifier/config"
)

type DB struct {
	C *mgo.Collection
}

type Check struct {
	URL   string
	Fails int
}

func NewDB() (*DB, error) {
	session, err := mgo.Dial(config.MongoURL)
	if err != nil {
		return nil, err
	}
	c := session.DB(config.MongoDB).C(config.MongoCollection)
	return &DB{C: c}, nil
}

func (d *DB) GetAll() {
}
