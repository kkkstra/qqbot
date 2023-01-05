package main

import (
	"fatsharkbot/src/config"
	"fatsharkbot/src/model"
	"fatsharkbot/src/router"
	"github.com/gin-gonic/gin"
)

func main() {
	model.InitDatabase()
	r := gin.Default()
	router.InitRouters(r)
	r.Run(config.C.App.Addr)
}
