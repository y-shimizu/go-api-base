package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getUser(c *gin.Context) {
	//TODO: db接続ができたらuserを受け取ってjsonで返す
	_, err := userRepository.Find(1)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, "user object")
}
