package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ltakato/gopportunities/handler"
)

func initializeRoutes(router *gin.Engine) {
	// Initialize Handler
	handler.InitializeHandler()

	v1 := router.Group("/api/v1")
	{
		// annoymous function as handler
		v1.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "pong",
			})
		})

		v1.GET("/openings", handler.ListOpeningsHandler)

		v1.POST("/openings", handler.CreateOpeningHandler)

		v1.DELETE("/opening", handler.DeleteOpeningHandler)

		v1.PUT("/opening", handler.UpdateOpeningHandler)

		v1.GET("/opening", handler.ShowOpeningHandler)
	}
}
