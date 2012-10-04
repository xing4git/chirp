package main

import (
	"github.com/xing4git/chirp/service/feedservice"
	"github.com/xing4git/chirp/util"
	"github.com/xing4git/chirp/conf"
	"github.com/xing4git/chirp/chirplog"
	"net"
	"errors"
	"net/rpc"
	"net/rpc/jsonrpc"
)

var log = chirplog.New("server.main")

func main() {
	port, ok := conf.Conf[util.CONF_KEY_SERVER_LISTEN_PORT]
	if !ok {
		util.StartupFatalErr(errors.New("Must contain " + util.CONF_KEY_SERVER_LISTEN_PORT + " in conf"))
	}

	fs := new(feedservice.FeedService)
	rpc.Register(fs)

	l, err := net.Listen("tcp", ":"+port)
	util.StartupFatalErr(err)
	log.Info("server start at port: " + port)

	for {
		conn, err := l.Accept()
		log.Infof("conn from remote: %+v", conn.RemoteAddr())
		if err != nil {
			log.Errorf("accept error: %s", err.Error())
		}
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}

// {"method":"FeedService.Create", "id":1, "params":[{"uid":"12345", "content":{"text":"feed text", "img":"http//www.xing.com/1.jpg"}}]}

// {"method":"FeedService.Delete", "id":2, "params":["506de522918d4a653f000001"]}
