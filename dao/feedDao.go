package dao

import (
	// "github.com/xing4git/chirp/log"
	"github.com/xing4git/chirp/util"
	"github.com/xing4git/chirp/mongo"
	"github.com/xing4git/chirp/model"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

func InsertFeed(feed model.Feed) (ret model.Feed, err error) {
	s := mongo.GetSession()
	defer s.Close()
	c := collection(s, util.MONGO_COLLECTION_FEED)

	feed.Fid = bson.NewObjectId()
	feed.Ctime = util.UnixMillSeconds()

	return feed, c.Insert(feed)
}

func InsertFeedsDel(feeds []model.Feed) (err error) {
	ins := make([]interface{}, len(feeds), len(feeds))
	for pos, v := range feeds {
		ins[pos] = v
	}

	s := mongo.GetSession()
	defer s.Close()
	c := collection(s, util.MONGO_COLLECTION_FEED_DEL)
	return c.Insert(ins...)
}

func RemoveFeed(id bson.ObjectId) (err error) {
	s := mongo.GetSession()
	defer s.Close()
	c := collection(s, util.MONGO_COLLECTION_FEED)

	return c.RemoveId(id)
}

func QueryFeed(id bson.ObjectId) (ret model.Feed, err error) {
	s := mongo.GetSession()
	defer s.Close()
	c := collection(s, util.MONGO_COLLECTION_FEED)

	q := c.FindId(id)
	if q.Iter().Next(&ret) {
		return ret, nil
	}
	return ret, mgo.ErrNotFound
}

func QueryBatchFeed(ids []bson.ObjectId) (rets []model.Feed, err error) {
	s := mongo.GetSession()
	defer s.Close()
	c := collection(s, util.MONGO_COLLECTION_FEED)

	iter := c.Find(bson.M{"_id": bson.M{"$in": ids}}).Iter()
	rets = make([]model.Feed, 0, len(ids))
	item := model.Feed{}
	for iter.Next(&item) {
		rets = append(rets, item)
	}
	return rets, nil
}
