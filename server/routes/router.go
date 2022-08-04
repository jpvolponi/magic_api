package routes

import "github.com/gin-gonic/gin"

func ConfigureRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("ap1/v1")
	{
		cards := main.Group("cards")
		{
			cards.GET("/", controllers.Showcards)
		}
	}
	return router
}
