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

	// CONF_KEY_REDIS_SYNC        = "redis.sync"
	CONF_KEY_REDIS_MAX_CLIENTS = "redis.max.clients"

	CONF_KEY_SERVER_LISTEN_PORT = "server.listen.port"
)

// mongo
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

// redis
const (
	REDIS_DATABASE = 1

	REDIS_USER_TIMELINE = "user:timeline:"
	REDIS_USER_AT       = "user:at:"
	REDIS_USER_FANS     = "user:fans:"
	REDIS_USER_FOLLOWS  = "user:follows:"
	REDIS_USER_FEEDS    = "user:feeds:"
	REDIS_FEED_COMMENTS = "feed:comments:"
	REDIS_FEED_FORWARDS = "feed:forwards:"
)

const (
	BACKEND_HANDLE_STEP = 100
)
