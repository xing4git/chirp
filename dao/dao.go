package dao

import (
	"github.com/xing4git/chirp/util"
	"github.com/xing4git/chirp/mongo"
	"github.com/xing4git/chirp/chirplog"
	"labix.org/v2/mgo"
)

var log = chirplog.New("dao")

func collection(s *mongo.Session, collectionName string) *mgo.Collection {
	db := s.DB(util.MONGO_DATABASE)
	return db.C(collectionName)
}
