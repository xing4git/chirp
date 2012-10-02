package dao

import (
	// "github.com/xing4git/chirp/log"
	"github.com/xing4git/chirp/util"
	"github.com/xing4git/chirp/mongo"
	"github.com/xing4git/chirp/model"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

func InsertUser(user model.User) (ret model.User, err error) {
	s := mongo.GetSession()
	defer s.Close()
	c := collection(s, util.MONGO_COLLECTION_USER)

	user.Uid = bson.NewObjectId()
	user.Ctime = util.UnixMillSeconds()

	return user, c.Insert(user)
}

func InsertUsersDel(users []model.User) (err error) {
	ins := make([]interface{}, len(users), len(users))
	for pos, v := range users {
		ins[pos] = v
	}

	s := mongo.GetSession()
	defer s.Close()
	c := collection(s, util.MONGO_COLLECTION_USER_DEL)
	return c.Insert(ins...)
}

func RemoveUser(id bson.ObjectId) (err error) {
	s := mongo.GetSession()
	defer s.Close()
	c := collection(s, util.MONGO_COLLECTION_USER)

	return c.RemoveId(id)
}

func QueryUser(id bson.ObjectId) (ret model.User, err error) {
	s := mongo.GetSession()
	defer s.Close()
	c := collection(s, util.MONGO_COLLECTION_USER)

	q := c.FindId(id)
	if q.Iter().Next(&ret) {
		return ret, nil
	}
	return ret, mgo.ErrNotFound
}

func QueryBatchUser(ids []bson.ObjectId) (rets []model.User, err error) {
	s := mongo.GetSession()
	defer s.Close()
	c := collection(s, util.MONGO_COLLECTION_USER)

	iter := c.Find(bson.M{"_id": bson.M{"$in": ids}}).Iter()
	rets = make([]model.User, 0, len(ids))
	item := model.User{}
	for iter.Next(&item) {
		rets = append(rets, item)
	}
	return rets, nil
}
