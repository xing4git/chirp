package dao

import (
	// "github.com/xing4git/chirp/log"
	"github.com/xing4git/chirp/util"
	"github.com/xing4git/chirp/mongo"
	"github.com/xing4git/chirp/model"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

func InsertFeedLoc(loc model.FeedLoc) (ret model.FeedLoc, err error) {
	s := mongo.GetSession()
	defer s.Close()
	c := collection(s, util.MONGO_COLLECTION_FEED_LOC)

	return loc, c.Insert(loc)
}

func InsertFeedLocsDel(locs []model.FeedLoc) (err error) {
	ins := make([]interface{}, len(locs), len(locs))
	for pos, v := range locs {
		ins[pos] = v
	}

	s := mongo.GetSession()
	defer s.Close()
	c := collection(s, util.MONGO_COLLECTION_FEED_LOC_DEL)
	return c.Insert(ins...)
}

func RemoveFeedLoc(id string) (err error) {
	s := mongo.GetSession()
	defer s.Close()
	c := collection(s, util.MONGO_COLLECTION_FEED_LOC)

	return c.Remove(bson.M{"fid": id})
}

func QueryFeedLoc(id string) (ret model.FeedLoc, err error) {
	s := mongo.GetSession()
	defer s.Close()
	c := collection(s, util.MONGO_COLLECTION_FEED_LOC)

	q := c.Find(bson.M{"fid": id})
	if q.Iter().Next(&ret) {
		return ret, nil
	}
	return ret, mgo.ErrNotFound
}

func WithinBoxFeedLoc(loc1 model.LatLon, loc2 model.LatLon) (rets []model.FeedLoc) {
	rets = make([]model.FeedLoc, 0, 20)
	s := mongo.GetSession()
	defer s.Close()
	c := collection(s, util.MONGO_COLLECTION_FEED_LOC)

	iter := c.Find(bson.M{"loc.loc": bson.M{"$within": bson.M{"$box": []model.LatLon{loc1, loc2}}}}).Iter()
	item := model.FeedLoc{}
	for iter.Next(&item) {
		rets = append(rets, item)
	}
	return
}

func WithinCircleFeedLoc(center model.LatLon, radius float64) (rets []model.FeedLoc) {
	rets = make([]model.FeedLoc, 0, 20)
	s := mongo.GetSession()
	defer s.Close()
	c := collection(s, util.MONGO_COLLECTION_FEED_LOC)

	iter := c.Find(bson.M{"loc.loc": bson.M{"$within": bson.M{"$center": []interface{}{center, radius}}}}).Iter()
	item := model.FeedLoc{}
	for iter.Next(&item) {
		rets = append(rets, item)
	}
	return
}
