package dao

import (
	// "github.com/xing4git/chirp/log"
	"github.com/xing4git/chirp/util"
	"github.com/xing4git/chirp/mongo"
	"github.com/xing4git/chirp/model"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

func InsertUserLoc(loc model.UserLoc) (ret model.UserLoc, err error) {
	s := mongo.GetSession()
	defer s.Close()
	c := collection(s, util.MONGO_COLLECTION_USER_LOC)

	return loc, c.Insert(loc)
}

func InsertUserLocsDel(locs []model.UserLoc) (err error) {
	ins := make([]interface{}, len(locs), len(locs))
	for pos, v := range locs {
		ins[pos] = v
	}

	s := mongo.GetSession()
	defer s.Close()
	c := collection(s, util.MONGO_COLLECTION_USER_LOC_DEL)
	return c.Insert(ins...)
}

func RemoveUserLoc(id string) (err error) {
	s := mongo.GetSession()
	defer s.Close()
	c := collection(s, util.MONGO_COLLECTION_USER_LOC)

	return c.Remove(bson.M{"uid": id})
}

func QueryUserLoc(id string) (ret model.UserLoc, err error) {
	s := mongo.GetSession()
	defer s.Close()
	c := collection(s, util.MONGO_COLLECTION_USER_LOC)

	q := c.Find(bson.M{"uid": id})
	if q.Iter().Next(&ret) {
		return ret, nil
	}
	return ret, mgo.ErrNotFound
}

func WithinBoxUserLoc(loc1 model.LatLon, loc2 model.LatLon) (rets []model.UserLoc) {
	rets = make([]model.UserLoc, 0, 20)
	s := mongo.GetSession()
	defer s.Close()
	c := collection(s, util.MONGO_COLLECTION_USER_LOC)

	iter := c.Find(bson.M{"loc.loc": bson.M{"$within": bson.M{"$box": []model.LatLon{loc1, loc2}}}}).Iter()
	item := model.UserLoc{}
	for iter.Next(&item) {
		rets = append(rets, item)
	}
	return
}

func WithinCircleUserLoc(center model.LatLon, radius float64) (rets []model.UserLoc) {
	rets = make([]model.UserLoc, 0, 20)
	s := mongo.GetSession()
	defer s.Close()
	c := collection(s, util.MONGO_COLLECTION_USER_LOC)

	iter := c.Find(bson.M{"loc.loc": bson.M{"$within": bson.M{"$center": []interface{}{center, radius}}}}).Iter()
	item := model.UserLoc{}
	for iter.Next(&item) {
		rets = append(rets, item)
	}
	return
}
