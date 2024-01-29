package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowOpeningHandler(ctx *gin.Context) {
	// instead of returning 200, using http module's StatusOk
	ctx.JSON(http.StatusOK, gin.H{
		"message": "GET opening",
	})
}
