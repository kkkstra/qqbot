package main

import (
	"fatsharkbot/src/config"
	"fatsharkbot/src/cron"
	"fatsharkbot/src/model"
	"fatsharkbot/src/router"
	"github.com/gin-gonic/gin"
)

func main() {
	model.InitDatabase()
	cron.InitCron()

	r := gin.Default()
	router.InitRouters(r)
	r.Run(config.C.App.Addr)
}
