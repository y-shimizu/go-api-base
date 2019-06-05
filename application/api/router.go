package api

import (
	"net/http"
	"strconv"

	"github.com/y-shimizu/go-api-base/util"
	"github.com/brandfolder/gin-gorelic"
	"github.com/gin-gonic/gin"
)

func Main() {
	router := gin.New()

	router.Use(AccessLogger())
	if util.Env.IsLocal() {
		router.Use(gin.Logger(), Recovery())
	}

	router.Use(gorelic.Handler)
	router.Use(CORSMiddleware())

	router.GET("/", func(c *gin.Context) {
		if !util.Env.IsProd() {
			c.HTML(http.StatusOK, "index.html", gin.H{})
		} else {
			c.String(http.StatusOK, "ok")
		}
	})

	//TODO: health check用のrouter.GET()

	//TODO: cache clear用のrouter.GET()

	UserEndpoint(router)

	port := strconv.Itoa(util.Conf.Server.Port)
	router.Run(":" + port)
}
