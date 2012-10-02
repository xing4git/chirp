package dao

import (
	// "github.com/xing4git/chirp/log"
	"github.com/xing4git/chirp/util"
	"github.com/xing4git/chirp/mongo"
	"github.com/xing4git/chirp/model"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

func InsertFollow(follow model.Follow) (ret model.Follow, err error) {
	s := mongo.GetSession()
	defer s.Close()
	c := collection(s, util.MONGO_COLLECTION_FOLLOW)

	follow.Ctime = util.UnixMillSeconds()
	return follow, c.Insert(follow)
}

func RemoveFollow(uid string, beuid string) (err error) {
	s := mongo.GetSession()
	defer s.Close()
	c := collection(s, util.MONGO_COLLECTION_FOLLOW)

	return c.Remove(bson.M{"uid": uid, "beuid": beuid})
}

func QueryFollow(uid string, beuid string) (ret model.Follow, err error) {
	s := mongo.GetSession()
	defer s.Close()
	c := collection(s, util.MONGO_COLLECTION_FOLLOW)

	q := c.Find(bson.M{"uid": uid, "beuid": beuid})
	if q.Iter().Next(&ret) {
		return ret, nil
	}
	return ret, mgo.ErrNotFound
}
