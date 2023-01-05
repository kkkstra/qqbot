package router

import (
	"fatsharkbot/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRouters(r *gin.Engine) {
	r.POST("/", controller.EventHandler)
}
