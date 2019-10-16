package mongodb

import (
	"gopkg.in/mgo.v2"
)

var baseSession *mgo.Session

func Connect(path string) {
	var err error
	baseSession, err = mgo.Dial(path)
	if err != nil {
		panic(err)
	}
}
