package main

import (
	"log"
	lost "newsletter-pub/api/v1/lostinspace"
	"newsletter-pub/http/router"
	"newsletter-pub/utils/config"

	"github.com/gin-gonic/gin"
)

var appConfig = config.AppCfg

func main() {

	app := gin.Default()
	app.Use(gin.Logger())
	app.Use(gin.Recovery())

	app.NoRoute(lost.LostInSpace)
	router.ApiRoutes(&app.RouterGroup)

	log.Fatal(app.Run(":" + appConfig.App.Port))
}
