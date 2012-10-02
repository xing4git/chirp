package model

import (
	"labix.org/v2/mgo/bson"
)

type Feed struct {
	Fid     bson.ObjectId "_id" // feed id
	Uid     string        // user id
	Content FeedContent   // content, contains text and img
	Refid   string        ",omitempty" // reference feed id
	Ctime   int64         // feed create time
}

type FeedContent struct {
	Text string // text content
	Img  string ",omitempty" // image url
}

type FeedLoc struct {
	Fid string // feed id
	Loc Location
}
