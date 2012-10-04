package redisdao

import (
	"github.com/xing4git/chirp/util"
	"github.com/xing4git/chirp/model"
	"github.com/xing4git/chirp/r"
	"github.com/xing4git/chirp/chirplog"
	"labix.org/v2/mgo/bson"
	// "github.com/alphazero/redis"
	"errors"
)

var log = chirplog.New("redisdao")

func RemoveKey(key string) (err error) {
	log.Warnf("remove key: %s", key)
	c := r.GetClient()
	defer c.Quit()

	success, err := c.Del(key)
	if err != nil {
		return err
	} else if !success {
		return errors.New("delete key " + key + " failed")
	}
	return nil
}

func InsertTimeline(uid string, feed model.Feed) (err error) {
	key := util.REDIS_USER_TIMELINE + uid
	return redisInsertFeed(key, feed)
}

// from < to
func TimelineFeeds(uid string, from int64, to int64) (rets []model.Feed, err error) {
	key := util.REDIS_USER_TIMELINE + uid
	return redisFeeds(key, from, to)
}

func InsertAt(uid string, feed model.Feed) (err error) {
	key := util.REDIS_USER_AT + uid
	return redisInsertFeed(key, feed)
}

// from < to
func AtFeeds(uid string, from int64, to int64) (rets []model.Feed, err error) {
	key := util.REDIS_USER_AT + uid
	return redisFeeds(key, from, to)
}

func InsertUserFeed(uid string, feed model.Feed) (err error) {
	key := util.REDIS_USER_FEEDS + uid
	return redisInsertFeed(key, feed)
}

// from < to
func UserFeeds(uid string, from int64, to int64) (rets []model.Feed, err error) {
	key := util.REDIS_USER_FEEDS + uid
	return redisFeeds(key, from, to)
}

func InsertFeedForward(fid string, feed model.Feed) (err error) {
	key := util.REDIS_FEED_FORWARDS + fid
	return redisInsertFeed(key, feed)
}

// from < to
func FeedForwards(fid string, from int64, to int64) (rets []model.Feed, err error) {
	key := util.REDIS_FEED_FORWARDS + fid
	return redisFeeds(key, from, to)
}

func InsertUserFan(uid string, user model.User) (err error) {
	key := util.REDIS_USER_FANS + uid
	return redisInsertUser(key, user)
}

// from < to
func UserFans(uid string, from int64, to int64) (rets []model.User, err error) {
	key := util.REDIS_USER_FANS + uid
	return redisUsers(key, from, to)
}

func InsertUserFollow(uid string, user model.User) (err error) {
	key := util.REDIS_USER_FOLLOWS + uid
	return redisInsertUser(key, user)
}

// from < to
func UserFollows(uid string, from int64, to int64) (rets []model.User, err error) {
	key := util.REDIS_USER_FOLLOWS + uid
	return redisUsers(key, from, to)
}

func InsertFeedComment(fid string, comment model.Comment) (err error) {
	key := util.REDIS_FEED_COMMENTS + fid
	return redisInsertComment(key, comment)
}

// from < to
func FeedComments(fid string, from int64, to int64) (rets []model.Comment, err error) {
	key := util.REDIS_FEED_COMMENTS + fid
	return redisComments(key, from, to)
}

func redisFeeds(key string, from int64, to int64) (rets []model.Feed, err error) {
	var strs []string
	strs, err = redisStrs(key, from, to)
	if err != nil {
		return nil, err
	}

	rets = make([]model.Feed, len(strs))
	for pos, str := range strs {
		feed := model.Feed{}
		feed.Fid = bson.ObjectIdHex(str)
		rets[pos] = feed
	}
	return
}

func redisUsers(key string, from int64, to int64) (rets []model.User, err error) {
	var strs []string
	strs, err = redisStrs(key, from, to)
	if err != nil {
		return nil, err
	}

	rets = make([]model.User, len(strs))
	for pos, str := range strs {
		user := model.User{}
		user.Uid = bson.ObjectIdHex(str)
		rets[pos] = user
	}
	return
}

func redisComments(key string, from int64, to int64) (rets []model.Comment, err error) {
	var strs []string
	strs, err = redisStrs(key, from, to)
	if err != nil {
		return nil, err
	}

	rets = make([]model.Comment, len(strs))
	for pos, str := range strs {
		comment := model.Comment{}
		comment.Cid = bson.ObjectIdHex(str)
		rets[pos] = comment
	}
	return
}

func redisInsertFeed(key string, feed model.Feed) (err error) {
	return redisInsert(key, feed.Ctime, []byte(feed.Fid.Hex()))
}

func redisInsertUser(key string, user model.User) (err error) {
	return redisInsert(key, user.Ctime, []byte(user.Uid.Hex()))
}

func redisInsertComment(key string, comment model.Comment) (err error) {
	return redisInsert(key, comment.Ctime, []byte(comment.Cid.Hex()))
}

func redisInsert(key string, score int64, data []byte) (err error) {
	c := r.GetClient()
	defer c.Quit()

	success, err := c.Zadd(key, float64(score), data)
	if err != nil {
		return err
	} else if !success {
		return errors.New("insert " + key + " failed")
	}
	return
}

func redisStrs(key string, from int64, to int64) (rets []string, err error) {
	c := r.GetClient()
	defer c.Quit()

	var bytess [][]byte
	bytess, err = c.Zrevrange(key, from, to)
	if err != nil {
		return
	}

	rets = make([]string, len(bytess))
	for pos, v := range bytess {
		rets[pos] = string(v)
	}
	log.Infof("size: %d, ids: %+v", len(rets), rets)
	return rets, nil
}
