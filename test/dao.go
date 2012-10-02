package main

import (
	"github.com/xing4git/chirp/log"
	"github.com/xing4git/chirp/model"
	"github.com/xing4git/chirp/dao"
	"labix.org/v2/mgo/bson"
)

func main() {
	// testQueryFeed()
	// testRemoveFeed()
	// testQueryBatchFeed()
	// testConcurrency()
	// testInsertFeedLoc()
	// testQueryFeedLoc()
	// testRemoveFeedLoc()
	// testWithinBoxFeedLoc()
	testWithinCircleFeedLoc()
}

func testWithinCircleFeedLoc() {
	rets := dao.WithinCircleFeedLoc([2]float64{1.2, 3.3}, 0.11)
	log.Logger.Debugf("within circle feed loc rets: %+v", rets)

	rets = dao.WithinCircleFeedLoc([2]float64{1.2, 3.3}, 5)
	log.Logger.Debugf("within circle feed loc rets: %+v", rets)
}

func testWithinBoxFeedLoc() {
	rets := dao.WithinBoxFeedLoc([2]float64{1.1, 1.1}, [2]float64{3.5, 3.5})
	log.Logger.Debugf("within box feed loc rets: %+v", rets)
}

func testRemoveFeedLoc() {
	err := dao.RemoveFeedLoc("506a4df0918d4a5951000001")
	if err != nil {
		log.Logger.Fatal(err)
	}
	log.Logger.Debugf("remove feed loc success")
}

func testQueryFeedLoc() {
	ret, err := dao.QueryFeedLoc("506a4df0918d4a5951000001")
	if err != nil {
		log.Logger.Fatal(err)
	}
	log.Logger.Debugf("query feed loc result: %+v", ret)
}

func testInsertFeedLoc() {
	feed := model.Feed{Uid: "abcdef987654321", Content: model.FeedContent{Text: "first feed"}}
	feed, err := dao.InsertFeed(feed)
	if err != nil {
		log.Logger.Fatal(err)
	}
	loc := model.FeedLoc{}
	loc.Fid = feed.Fid.Hex()
	loc.Loc = model.Location{Ctime: feed.Ctime, Loc: [2]float64{3.4, 1.2}}
	loc, err = dao.InsertFeedLoc(loc)
	if err != nil {
		log.Logger.Fatal(err)
	}
	log.Logger.Debugf("insert feed loc: %+v", loc)
}

func testConcurrency() {
	ch := make(chan bool)
	for i := 0; i < 15; i++ {
		go testInsertFeed()
	}
	for i := 0; i < 15; i++ {
		go testQueryFeed()
	}
	<-ch
}

func testInsertFeed() {
	feed := model.Feed{Uid: "abcdef987654321", Content: model.FeedContent{Text: "second feed"}}
	log.Logger.Debugf("insert feed: %+v", feed)
	feed, err := dao.InsertFeed(feed)
	if err != nil {
		log.Logger.Fatal(err)
	}
	log.Logger.Debugf("insert ret: %+v", feed)
}

func testQueryFeed() {
	id := bson.ObjectIdHex("5069df87918d4a4c83000001")
	log.Logger.Debugf("query feed: %+v", id)
	feed, err := dao.QueryFeed(id)
	if err != nil {
		log.Logger.Fatal(err)
	}
	log.Logger.Debugf("query ret: %+v", feed)
}

func testQueryBatchFeed() {
	ids := []bson.ObjectId{bson.ObjectIdHex("5069d898918d4a4aa2000001"), bson.ObjectIdHex("5069df87918d4a4c83000001")}
	log.Logger.Debugf("query feed: %+v", ids)
	log.Logger.Debug("begin")
	feeds, err := dao.QueryBatchFeed(ids)
	log.Logger.Debug("end")
	if err != nil {
		log.Logger.Fatal(err)
	}
	log.Logger.Debugf("query rets: %+v", feeds)
}

func testRemoveFeed() {
	id := bson.ObjectIdHex("5069df87918d4a4c83000001")
	log.Logger.Debugf("remove feed: %+v", id)
	err := dao.RemoveFeed(id)
	if err != nil {
		log.Logger.Fatal(err)
	}
	log.Logger.Debugf("remove feed success")
}
