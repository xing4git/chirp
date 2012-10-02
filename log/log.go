package log

import (
	"github.com/xing4git/chirp/conf"
	"github.com/xing4git/chirp/util"
	"github.com/xing4git/golog"
	"errors"
	"os"
	"strconv"
)

var Logger *golog.Logger
var logfile *os.File

func init() {
	logpath, ok := conf.Conf[util.CONF_KEY_LOGPATH]
	if !ok {
		util.StartupFatalErr(errors.New("Must contain " + util.CONF_KEY_LOGPATH + " in conf"))
	}
	var err error
	logfile, err = os.OpenFile(logpath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0664)
	util.StartupFatalErr(err)

	loglevel, ok := conf.Conf[util.CONF_KEY_LOGLEVEL]
	if !ok {
		util.StartupFatalErr(errors.New("Must contain " + util.CONF_KEY_LOGLEVEL + " in conf"))
	}
	var level int
	level, err = strconv.Atoi(loglevel)
	util.StartupFatalErr(err)

	Logger = golog.NewLogger(logfile, "", level, golog.FLAG_LstdFlags)
}

func Shutdown() {
	logfile.Close()
}
