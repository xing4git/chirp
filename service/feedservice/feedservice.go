package feedservice

import (
	"github.com/xing4git/chirp/model"
	"github.com/xing4git/chirp/dao"
	"github.com/xing4git/chirp/backend"
	"github.com/xing4git/chirp/chirplog"
	"labix.org/v2/mgo/bson"
	"fmt"
)

var log = chirplog.New("FeedService")

type FeedService struct{}

func (fs *FeedService) Create(feed model.Feed, ret *model.Feed) error {
	logstr := fmt.Sprintf("create feed: %+v", feed)
	log.Info(logstr)
	tmp, err := dao.InsertFeed(feed)
	if err != nil {
		log.Errorf("%s, error: ", logstr, err.Error())
		return err
	}
	*ret = tmp
	log.Infof("create feed result: %+v", tmp)
	backend.CreateFeed(tmp)
	return nil
}

func (fs *FeedService) Delete(fid string, ret *model.Feed) error {
	logstr := fmt.Sprintf("delete feed: %s", fid)
	log.Info(logstr)
	bid := bson.ObjectIdHex(fid)
	removed, err := dao.QueryFeed(bid)
	if err != nil {
		log.Errorf("%s, error: %s", logstr, err.Error())
		return err
	}
	err = dao.RemoveFeed(bid)
	if err != nil {
		log.Errorf("%s, error: %s", logstr, err.Error())
		return err
	}
	*ret = removed
	backend.DeleteFeed(removed)
	return nil
}
