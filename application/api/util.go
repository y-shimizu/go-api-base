package api

import (
	"os"

	"github.com/y-shimizu/go-api-base/util"

	log "github.com/Sirupsen/logrus"
	"github.com/doloopwhile/logrusltsv"
)


func NewAccessLogger() *log.Logger {
	l := log.New()
	f, err := os.OpenFile(util.Conf.Logger.AccessLogPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	l.Out = f
	l.Formatter = &logrusltsv.Formatter{}
	return l
}
