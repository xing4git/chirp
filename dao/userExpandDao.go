package dao

import (
	// "github.com/xing4git/chirp/log"
	"github.com/xing4git/chirp/util"
	"github.com/xing4git/chirp/mongo"
	"github.com/xing4git/chirp/model"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

func InsertUserExpand(ue model.UserExpand) (ret model.UserExpand, err error) {
	s := mongo.GetSession()
	defer s.Close()
	c := collection(s, util.MONGO_COLLECTION_USER_EXPAND)
	return ue, c.Insert(ue)
}

func InsertUserExpandsDel(ues []model.UserExpand) (err error) {
	ins := make([]interface{}, len(ues), len(ues))
	for pos, v := range ues {
		ins[pos] = v
	}

	s := mongo.GetSession()
	defer s.Close()
	c := collection(s, util.MONGO_COLLECTION_USER_EXPAND_DEL)
	return c.Insert(ins...)
}

func RemoveUserExpand(id string) (err error) {
	s := mongo.GetSession()
	defer s.Close()
	c := collection(s, util.MONGO_COLLECTION_USER_EXPAND)
	return c.Remove(bson.M{"uid": id})
}

func UpdateUserExpand(ue model.UserExpand) (ret model.UserExpand, err error) {
	s := mongo.GetSession()
	defer s.Close()
	c := collection(s, util.MONGO_COLLECTION_USER_EXPAND)

	cmd := bson.D{{"findAndModify", util.MONGO_COLLECTION_USER_EXPAND}, {"new", true}, {"query", bson.M{"uid": ue.Uid}}, {"update", bson.M{"$set": ue.UserExpandToBson()}}}
	err = c.Database.Run(cmd, &ret)
	return
}

func UserLogin(uid string) (err error) {
	s := mongo.GetSession()
	defer s.Close()
	c := collection(s, util.MONGO_COLLECTION_USER_EXPAND)
	err = c.Update(bson.M{"uid": uid}, bson.M{"$set": bson.M{"lltime": util.UnixMillSeconds()}, "$inc": bson.M{"logincnt": 1}})
	return
}

func QueryUserExpand(id string) (ret model.UserExpand, err error) {
	s := mongo.GetSession()
	defer s.Close()
	c := collection(s, util.MONGO_COLLECTION_USER_EXPAND)

	q := c.Find(bson.M{"uid": id})
	if q.Iter().Next(&ret) {
		return ret, nil
	}
	return ret, mgo.ErrNotFound
}

func QueryBatchUserExpand(ids []string) (rets []model.UserExpand, err error) {
	s := mongo.GetSession()
	defer s.Close()
	c := collection(s, util.MONGO_COLLECTION_USER_EXPAND)

	iter := c.Find(bson.M{"uid": bson.M{"$in": ids}}).Iter()
	rets = make([]model.UserExpand, 0, len(ids))
	item := model.UserExpand{}
	for iter.Next(&item) {
		rets = append(rets, item)
	}
	return rets, nil
}
