package r

import (
	"github.com/alphazero/redis"
	"github.com/xing4git/chirp/util"
	"github.com/xing4git/chirp/conf"
	"github.com/xing4git/chirp/chirplog"
	"errors"
	"strconv"
	"time"
)

var log = chirplog.New("r/r.go")

type pool struct {
	clients     []redis.Client
	avaliableCh chan redis.Client
}

type Client struct {
	start time.Time
	redis.Client
	sp *pool
}

func (c *Client) Quit() {
	if c.Client != nil {
		c.sp.avaliableCh <- c.Client
		c.Client = nil
		log.Debugf("client usage time: %f ms", time.Now().Sub(c.start).Seconds()*1e3)
	}
}

var sp *pool

func GetClient() *Client {
	start := time.Now()
	c := &Client{}
	c.Client = <-sp.avaliableCh
	c.sp = sp
	end := time.Now()
	c.start = end
	log.Debugf("fetch client time: %f ms", end.Sub(start).Seconds()*1e3)
	return c
}

func init() {
	sp = &pool{}

	maxClientsStr, ok := conf.Conf[util.CONF_KEY_REDIS_MAX_CLIENTS]
	if !ok {
		util.StartupFatalErr(errors.New("Must contain " + util.CONF_KEY_REDIS_MAX_CLIENTS + " in conf"))
	}
	maxClients, err := strconv.Atoi(maxClientsStr)
	util.StartupFatalErr(err)
	if maxClients <= 0 {
		util.StartupFatalErr(errors.New(util.CONF_KEY_REDIS_MAX_CLIENTS + " must larger then 0"))
	}

	sp.clients = make([]redis.Client, maxClients)
	sp.avaliableCh = make(chan redis.Client, maxClients)

	spec := redis.DefaultSpec().Db(util.REDIS_DATABASE)
	var c redis.Client
	for i := 0; i < maxClients; i++ {
		c, err = redis.NewSynchClientWithSpec(spec)
		util.StartupFatalErr(err)
		sp.clients = append(sp.clients, c)
		sp.avaliableCh <- c
	}
}

func Shutdown() {
	for _, c := range sp.clients {
		c.Quit()
	}
	close(sp.avaliableCh)
}
