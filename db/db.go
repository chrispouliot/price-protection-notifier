package db

import (
	"github.com/moxuz/price-protection-notifier/config"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type DB struct {
	C *mgo.Collection
}

type Check struct {
	LastPrice float64 `bson:"price"`
	URL       string  `bson:"url"`
	Fails     int     `bson:"fails"`
	ID        int     `bson:"_Id"`
}

func NewDB() (*DB, error) {
	session, err := mgo.Dial(config.MongoURL)
	if err != nil {
		return nil, err
	}
	session.SetSafe(&mgo.Safe{})
	c := session.DB(config.MongoDB).C(config.MongoCollection)
	return &DB{C: c}, nil
}

func (d *DB) GetAll() ([]*Check, error) {
	var results []*Check
	iter := d.C.Find(bson.M{}).Iter()
	err := iter.All(&results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (d *DB) MarkFailed(check *Check) error {
	return nil
}
