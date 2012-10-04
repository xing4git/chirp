package chirplog

import (
	"github.com/xing4git/chirp/conf"
	"github.com/xing4git/chirp/util"
	"github.com/xing4git/golog"
	"errors"
	"os"
	"strconv"
	"strings"
)

var logfile *os.File
var loglevel int

func init() {
	logpath, ok := conf.Conf[util.CONF_KEY_LOGPATH]
	if !ok {
		util.StartupFatalErr(errors.New("Must contain " + util.CONF_KEY_LOGPATH + " in conf"))
	}
	var err error
	logfile, err = os.OpenFile(logpath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0664)
	util.StartupFatalErr(err)

	loglevelstr, ok := conf.Conf[util.CONF_KEY_LOGLEVEL]
	if !ok {
		util.StartupFatalErr(errors.New("Must contain " + util.CONF_KEY_LOGLEVEL + " in conf"))
	}
	loglevel, err = strconv.Atoi(loglevelstr)
	util.StartupFatalErr(err)
}

func New(prefix string) *golog.Logger {
	return golog.NewLogger(logfile, strings.ToUpper(prefix), loglevel, golog.FLAG_LstdFlags)
}

func Shutdown() {
	logfile.Close()
}
