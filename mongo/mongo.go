package mongo

import (
	"github.com/xing4git/chirp/util"
	"github.com/xing4git/chirp/conf"
	"github.com/xing4git/chirp/chirplog"
	"errors"
	"strconv"
	"time"
	"labix.org/v2/mgo"
)

var log = chirplog.New("mongo")

type pool struct {
	sessions    []*mgo.Session
	avaliableCh chan *mgo.Session
}

type Session struct {
	start time.Time
	*mgo.Session
	sp *pool
}

func (s *Session) Close() {
	if s.Session != nil {
		s.sp.avaliableCh <- s.Session
		s.Session = nil
		log.Debugf("session usage time: %f ms", time.Now().Sub(s.start).Seconds()*1e3)
	}
}

var sp *pool

// wrap a mgo.Session to mongo.Session
func GetSession() *Session {
	start := time.Now()
	s := &Session{}
	s.Session = <-sp.avaliableCh
	s.sp = sp
	end := time.Now()
	s.start = end
	log.Debugf("fetch session time: %f ms", end.Sub(start).Seconds()*1e3)
	return s
}

func init() {
	sp = &pool{}

	url, ok := conf.Conf[util.CONF_KEY_MONGO_URL]
	if !ok {
		util.StartupFatalErr(errors.New("Must contain " + util.CONF_KEY_MONGO_URL + " in conf"))
	}

	maxSessonsStr, ok := conf.Conf[util.CONF_KEY_MONGO_MAX_SESSIONS]
	if !ok {
		util.StartupFatalErr(errors.New("Must contain " + util.CONF_KEY_MONGO_MAX_SESSIONS + " in conf"))
	}
	maxSessons, err := strconv.Atoi(maxSessonsStr)
	util.StartupFatalErr(err)
	if maxSessons <= 0 {
		util.StartupFatalErr(errors.New(util.CONF_KEY_MONGO_MAX_SESSIONS + " must larger then 0"))
	}

	sp.sessions = make([]*mgo.Session, 0, maxSessons)
	sp.avaliableCh = make(chan *mgo.Session, maxSessons)

	s, err := mgo.Dial(url)
	util.StartupFatalErr(err)
	s.SetSafe(&mgo.Safe{})
	sp.sessions = append(sp.sessions, s)
	sp.avaliableCh <- s

	for i := 1; i < maxSessons; i++ {
		news := s.Copy()
		sp.sessions = append(sp.sessions, news)
		sp.avaliableCh <- news
	}
}

func Shutdown() {
	for _, s := range sp.sessions {
		s.Close()
	}
	close(sp.avaliableCh)
}
