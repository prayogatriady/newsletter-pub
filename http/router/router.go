package router

import (
	"github.com/gin-gonic/gin"
)

func ApiRoutes(app *gin.RouterGroup) {
	injectEmail := InitEmail()
	injectPing := InitPing()

	api := app.Group("/api")
	{
		api.GET("/ping", injectPing.Ping)

		emailGroup := api.Group("/email")
		{
			emailGroup.POST("/sendEmail", injectEmail.SendEmail)
		}
	}
}
