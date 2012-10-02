package model

import (
	"labix.org/v2/mgo/bson"
)

type User struct {
	Uid      bson.ObjectId "_id" // user id
	Username string        // unique name
	Email    string        // unique email
	Pwd      string        // encrypted password
	Avatar   string        // avatar url
	Sex      int           // 0 means man, 1 means woman, 2 means others
	Ctime    int64         // register time, millisecond
}

type UserExpand struct {
	Uid         string // user id, required
	Blog        string ",omitempty" // blog url
	Address     string ",omitempty" // user address
	Birthday    string ",omitempty" // yyyy-mm-dd
	Phone       string ",omitempty" // phone number
	QQ          string ",omitempty" // qq number
	Msn         string ",omitempty" // msn account
	Description string ",omitempty" // self description
	Logincnt    int    ",omitempty" // login count
	Lltime      int64  ",omitempty" // last login time
}

type UserLoc struct {
	Uid   string     // user id
	Lloc  Location   // last location
	Hlocs []Location ",omitempty" // history locations
}

func (ue UserExpand) UserExpandToBson() (ret bson.M) {
	ret = bson.M{}
	if ue.Blog != "" {
		ret["blog"] = ue.Blog
	}
	if ue.Address != "" {
		ret["address"] = ue.Address
	}
	if ue.Birthday != "" {
		ret["birthday"] = ue.Birthday
	}
	if ue.Phone != "" {
		ret["phone"] = ue.Phone
	}
	if ue.QQ != "" {
		ret["qq"] = ue.QQ
	}
	if ue.Msn != "" {
		ret["msn"] = ue.Msn
	}
	if ue.Description != "" {
		ret["description"] = ue.Description
	}
	if ue.Logincnt != 0 {
		ret["logincnt"] = ue.Logincnt
	}
	if ue.Lltime != 0 {
		ret["lltime"] = ue.Lltime
	}
	return
}
