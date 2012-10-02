package model

type Follow struct {
	Uid   string // user id
	Beuid string // the user id who was followed
	Ctime int64  // create time
}
