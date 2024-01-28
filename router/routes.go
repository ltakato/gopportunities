package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		// annoymous function as handler
		v1.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "pong",
			})
		})

		v1.GET("/openings", func(ctx *gin.Context) {
			// instead of returning 200, using http module's StatusOk
			ctx.JSON(http.StatusOK, gin.H{
				"message": "GET opening",
			})
		})
	}
}
