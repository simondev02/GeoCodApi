package routers

import (
	"SimonBK_Vehiculo/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter configura las rutas de la aplicación
func SetupRouter(r *gin.Engine) {

	// Validacion de acces Token
	r.Use(middleware.ValidateTokenMiddleware())

	// Grupo de rutas para vehiculos
	geoCod := r.Group("/vehicle")
	{
		geoCod.GET("/geocoding", controllers.GeocodingController)
	}

}
