package util

import (
	"io"
	"os"

	log "github.com/Sirupsen/logrus"
)

var ActionLogger *log.Logger

func NewAppLogWriter() io.Writer {
	if Env.IsLocal() {
		return io.MultiWriter(openLogFile(Conf.Logger.AppLogPath), os.Stdout)
	}
	return openLogFile(Conf.Logger.AppLogPath)
}

func openLogFile(path string) io.Writer {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	return f
}
