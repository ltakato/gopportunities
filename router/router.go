package router

// jeito preguiçoso de puxar moódulos (em vez de instalar lib de fato)
// ai executo "go tidy" para resolver essa dependência
// retornando go.sum que é o "lock" de dependencias do go
// jeito legal de pensar: você não importa ARQUIVO (que nem node), mas sim o PACOTE (ref via "namespace do package")
import "github.com/gin-gonic/gin"

// maiusculo - automaticamente exportado do package (public)
func Initialize() {
	// Initialize Router
	router := gin.Default()

	// Initialize Routes
	// passing router through pointer, instead of value
	// (not sending the whole house, but the house's "address")
	initializeRoutes(router)

	// Running server
	router.Run(":8080")
}
