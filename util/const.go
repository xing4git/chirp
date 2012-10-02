package util

const (
	_ = iota
	EXIT_STATUS_STARTUP_ERR
)

const (
	CONF_KEY_LOGPATH  = "log.path"
	CONF_KEY_LOGLEVEL = "log.level"

	CONF_KEY_MONGO_URL          = "mongo.url"
	CONF_KEY_MONGO_MAX_SESSIONS = "mongo.max.sessions"
)

// mongo database and collections name
const (
	MONGO_DATABASE                    = "chirp"
	MONGO_COLLECTION_FEED             = "feed"
	MONGO_COLLECTION_FEED_DEL         = "feed_del"
	MONGO_COLLECTION_FEED_CONTENT     = "feedContent"
	MONGO_COLLECTION_FEED_CONTENT_DEL = "feedContent_del"
	MONGO_COLLECTION_FEED_LOC         = "feedLoc"
	MONGO_COLLECTION_FEED_LOC_DEL     = "feedLoc_del"
	MONGO_COLLECTION_COMMENT          = "comment"
	MONGO_COLLECTION_COMMENT_DEL      = "comment_del"

	MONGO_COLLECTION_USER            = "user"
	MONGO_COLLECTION_USER_DEL        = "user_del"
	MONGO_COLLECTION_USER_EXPAND     = "userExpand"
	MONGO_COLLECTION_USER_EXPAND_DEL = "userExpand_del"
	MONGO_COLLECTION_USER_LOC        = "userLoc"
	MONGO_COLLECTION_USER_LOC_DEL    = "userLoc_del"

	MONGO_COLLECTION_FOLLOW = "follow"
)
