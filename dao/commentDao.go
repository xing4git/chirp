package dao

import (
	// "github.com/xing4git/chirp/log"
	"github.com/xing4git/chirp/util"
	"github.com/xing4git/chirp/mongo"
	"github.com/xing4git/chirp/model"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

func InsertComment(comment model.Comment) (ret model.Comment, err error) {
	s := mongo.GetSession()
	defer s.Close()
	c := collection(s, util.MONGO_COLLECTION_COMMENT)

	comment.Cid = bson.NewObjectId()
	comment.Ctime = util.UnixMillSeconds()

	return comment, c.Insert(comment)
}

func InsertCommentsDel(comments []model.Comment) (err error) {
	ins := make([]interface{}, len(comments), len(comments))
	for pos, v := range comments {
		ins[pos] = v
	}

	s := mongo.GetSession()
	defer s.Close()
	c := collection(s, util.MONGO_COLLECTION_COMMENT_DEL)
	return c.Insert(ins...)
}

func RemoveComment(id bson.ObjectId) (err error) {
	s := mongo.GetSession()
	defer s.Close()
	c := collection(s, util.MONGO_COLLECTION_COMMENT)
	return c.RemoveId(id)
}

func QueryComment(id bson.ObjectId) (ret model.Comment, err error) {
	s := mongo.GetSession()
	defer s.Close()
	c := collection(s, util.MONGO_COLLECTION_COMMENT)

	q := c.FindId(id)
	if q.Iter().Next(&ret) {
		return ret, nil
	}
	return ret, mgo.ErrNotFound
}

func QueryFeedComment(fid string) (rets []model.Comment) {
	return queryComments("fid", fid)
}

func QueryUserComment(uid string) (rets []model.Comment) {
	return queryComments("uid", uid)
}

func queryComments(idName string, id string) (rets []model.Comment) {
	rets = make([]model.Comment, 0, 20)
	s := mongo.GetSession()
	defer s.Close()
	c := collection(s, util.MONGO_COLLECTION_COMMENT)

	iter := c.Find(bson.M{idName: id}).Iter()
	item := model.Comment{}
	for iter.Next(&item) {
		rets = append(rets, item)
	}
	return
}
