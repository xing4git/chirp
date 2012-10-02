package model

import (
	"labix.org/v2/mgo/bson"
)

type Comment struct {
	Cid     bson.ObjectId "_id" // comment id
	Uid     string        // user id
	Fid     string        // feed id
	Content string        // comment content
	Ctime   int64         // create time
}
