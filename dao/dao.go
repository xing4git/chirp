package dao

import (
	"github.com/xing4git/chirp/util"
	"github.com/xing4git/chirp/mongo"
	"labix.org/v2/mgo"
)

func collection(s *mongo.Session, collectionName string) *mgo.Collection {
	db := s.DB(util.MONGO_DATABASE)
	return db.C(collectionName)
}
