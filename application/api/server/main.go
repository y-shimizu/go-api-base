package main

import (
	"github.com/y-shimizu/go-api-base/application/api"
	"github.com/y-shimizu/go-api-base/util"
	"log"
)


func main() {
	log.SetOutput(util.NewAppLogWriter())
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("start: application/api")

	api.Main()
}
