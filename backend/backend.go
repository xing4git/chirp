package backend

import (
	"github.com/xing4git/chirp/model"
	"github.com/xing4git/chirp/dao"
	"github.com/xing4git/chirp/util"
	"github.com/xing4git/chirp/chirplog"
	"github.com/xing4git/chirp/dao/redisdao"
	"fmt"
)

var log = chirplog.New("backend")

func CreateFeed(feed model.Feed) {
	f := func() {
		logstr := fmt.Sprintf("create feed: %+v", feed)
		// insert into self timeline
		log.Infof("%s, insert into self timeline", logstr)
		err := redisdao.InsertTimeline(feed.Uid, feed)
		if err != nil {
			log.Errorf("%s, insert into self timeline error: %s", logstr, err.Error())
		}

		// insert into self feeds list
		log.Infof("%s, insert into self feeds list", logstr)
		err = redisdao.InsertUserFeed(feed.Uid, feed)
		if err != nil {
			log.Errorf("%s, insert into self feeds list error: %s", logstr, err.Error())
		}

		insertIntoFansTimeline(feed)

		// insert forwards list
		if feed.Refid != "" {
			log.Infof("%s, insert into feed forwards list", logstr)
			err = redisdao.InsertFeedForward(feed.Refid, feed)
			if err != nil {
				log.Errorf("%s, insert into forwards list error: %s", logstr, err.Error())
			}
		}
	}
	go util.SafeEmptyFunc(f)()
}

// insert into fans timeline
func insertIntoFansTimeline(feed model.Feed) {
	f := func() {
		logstr := fmt.Sprintf("create feed: %+v", feed)
		var startIndex int64 = 0
		var step int64 = util.BACKEND_HANDLE_STEP
		for {
			log.Infof("%s, insert into fans timeline, startIndex = %d", logstr, startIndex)
			fans, err := redisdao.UserFans(feed.Uid, startIndex, startIndex+step)
			if err != nil {
				log.Errorf("%s, insert into fans timeline error: %s", logstr, err.Error())
				break
			} else {
				insertIntoPartFansTimeline(feed, fans)
				// finish
				if len(fans) < int(step) {
					break
				}
				startIndex += startIndex + step + 1
			}
		}
	}
	go util.SafeEmptyFunc(f)()
}

func insertIntoPartFansTimeline(feed model.Feed, fans []model.User) {
	f := func() {
		for _, fan := range fans {
			fanuid := fan.Uid.Hex()
			err := redisdao.InsertTimeline(fanuid, feed)
			if err != nil {
				log.Errorf("insert into fans(%s) timeline error: %s", fanuid, err.Error())
			}
		}
	}
	go util.SafeEmptyFunc(f)()
}

func DeleteFeed(feed model.Feed) {
	f := func() {
		logstr := fmt.Sprintf("delete feed: %+v", feed)

		// insert feed del
		log.Infof("%s, insert into feed del", logstr)
		feeds := make([]model.Feed, 1)
		feeds[1] = feed
		err := dao.InsertFeedsDel(feeds)
		if err != nil {
			log.Errorf("%s, insert into feed del error: %s", logstr, err.Error())
		}

		// query and delete feed loc
		log.Infof("%s, query and feed loc", logstr)
		feedloc, err := dao.QueryFeedLoc(feed.Fid.Hex())
		if err != nil {
			log.Errorf("%s, query feed loc error: %s", logstr, err.Error())
		} else {
			err = dao.RemoveFeedLoc(feedloc.Fid)
			if err != nil {
				log.Errorf("%s, delete feed loc error: %s", logstr, err.Error())
			}

			locs := make([]model.FeedLoc, 1)
			locs[1] = feedloc
			err = dao.InsertFeedLocsDel(locs)
			if err != nil {
				log.Errorf("%s, insert into feed loc del error: %s", logstr, err.Error())
			}
		}

		// remove feed comment
		deleteFeedComments(feed)
	}
	go util.SafeEmptyFunc(f)()
}

func deleteFeedComments(feed model.Feed) {
	f := func() {
		logstr := fmt.Sprintf("delete feed: %+v", feed)
		var err error

		// remove feed comment
		var startIndex int64 = 0
		var step int64 = util.BACKEND_HANDLE_STEP
		for {
			log.Infof("%s, remove feed comments, startIndex = %d", logstr, startIndex)
			comments, err := redisdao.FeedComments(feed.Fid.Hex(), startIndex, startIndex+step)
			if err != nil {
				log.Errorf("%s, remove feed comments error: %s", logstr, err.Error())
				break
			} else {
				DeleteComments(comments)
				if len(comments) < int(step) {
					break
				}
				startIndex += startIndex + step + 1
			}
		}

		// remove feed_comments key
		err = redisdao.RemoveKey(util.REDIS_FEED_COMMENTS + feed.Fid.Hex())
		if err != nil {
			log.Errorf("%s, remove feed_comments key error: %s", logstr, err.Error())
		}
	}
	go util.SafeEmptyFunc(f)()
}

func DeleteComments(comments []model.Comment) {
	f := func() {
		var err error
		var deletedComments []model.Comment = make([]model.Comment, 0, len(comments))

		for _, comment := range comments {
			logstr := fmt.Sprintf("delete comment: %s", comment.Cid.Hex())
			comment, err = dao.QueryComment(comment.Cid)
			if err != nil {
				log.Errorf("%s, query comment error: %s", logstr, err.Error())
			}
			deletedComments = append(deletedComments, comment)

			err = dao.RemoveComment(comment.Cid)
			if err != nil {
				log.Errorf("%s, remove comment error: %s", logstr, err.Error())
			}
		}

		err = dao.InsertCommentsDel(deletedComments)
		if err != nil {
			log.Errorf("insert comments del error: %s", err.Error())
		}
	}
	go util.SafeEmptyFunc(f)()
}
