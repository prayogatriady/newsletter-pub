package router

import "github.com/gin-gonic/gin"

func ApiRoutes(app *gin.RouterGroup) {
	injectEmail := InitEmail()

	api := app.Group("/api")
	{
		emailGroup := api.Group("/email")
		{
			emailGroup.POST("/sendEmail", injectEmail.SendEmail)
		}
	}
}
