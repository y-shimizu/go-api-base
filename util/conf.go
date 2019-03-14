package util

import (
	"log"

	"github.com/BurntSushi/toml"
	"github.com/y-shimizu/go-api-base/bindata"
)

var Conf conf = newConf()

type conf struct {
	Logger               loggerConfig
	Server               serverConfig
}

type loggerConfig struct {
	AppLogPath      string `toml:"app_log_path"`
	AccessLogPath   string `toml:"access_log_path"`
}

type serverConfig struct {
	Port        int    `toml:"port"`
}



func newConf() conf {
	var _conf conf
	var confPath string

	switch Env {
	case Local:
		confPath = "conf/conf.local.toml"
	case Dev:
		confPath = "conf/conf.dev.toml"
	case Stg:
		confPath = "conf/conf.stg.toml"
	case Prod:
		confPath = "conf/conf.prod.toml"

	default:
		confPath = "conf/conf.local.toml"
	}

	asset, err := bindata.Asset(confPath)
	if err != nil {
		log.Fatalf("Failed to read bindata assets confPath: %s err:%s", confPath, err.Error())
	}

	_, err = toml.Decode(string(asset), &_conf)
	if err != nil {
		log.Fatalf("[CONFIGURATION FILE LOAD ERROR] confPath: %s err:%s", confPath, err.Error())
	}

	return _conf
}

