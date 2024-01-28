package router

// jeito preguiçoso de puxar moódulos (em vez de instalar lib de fato)
// ai executo "go tidy" para resolver essa dependência
// retornando go.sum que é o "lock" de dependencias do go
// jeito legal de pensar: você não importa ARQUIVO (que nem node), mas sim o PACOTE (ref via "namespace do package")
import "github.com/gin-gonic/gin"

// maiusculo - automaticamente exportado do package (public)
func Initialize() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run(":8080")
}
