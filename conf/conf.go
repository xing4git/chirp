package conf

import (
	goprop "github.com/xing4git/goprop"
	"github.com/xing4git/chirp/util"
	"flag"
	"os"
	"path"
	"fmt"
)

var Conf map[string]string

func init() {
	flag.Parse()
	confPath := *flag.String("conf", path.Join(os.Getenv("HOME"), "conf/chirp.conf"), "conf file path")
	var err error
	Conf, err = goprop.Load(confPath)
	util.StartupFatalErr(err)
	fmt.Printf("%+v\n", Conf)
}
