package db

import (
	"crypto/tls"
	"net"

	"github.com/moxuz/price-protection-notifier/config"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type DB struct {
	C *mgo.Collection
}

type insert struct {
	LastPrice float64 `bson:"price"`
	URL       string  `bson:"url"`
	Fails     int     `bson:"fails"`
}

type Check struct {
	LastPrice float64 `bson:"price"`
	URL       string  `bson:"url"`
	Fails     int     `bson:"fails"`
	ID        int     `bson:"_id"`
}

func NewDB() (*DB, error) {
	dialInfo, err := mgo.ParseURL(config.MongoURL)
	tlsConfig := &tls.Config{}
	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		return conn, err
	}
	session, _ := mgo.DialWithInfo(dialInfo)
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
	err := d.C.UpdateId(check.ID, struct {
		Fails int `bson:"fails"`
	}{
		Fails: check.Fails + 1,
	})
	return err
}

func (d *DB) Insert(url string, price float64) error {
	err := d.C.Insert(insert{
		LastPrice: price,
		URL:       url,
	})
	return err
}
