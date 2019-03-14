package api

import "github.com/gin-gonic/gin"

func UserEndpoint(router *gin.Engine) {

	router.Group("user/:user_id").GET("", getUser)

}
