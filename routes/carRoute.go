package routes

import (
	"belajar-gin/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	route := gin.Default()

	route.POST("/cars", controllers.CreateCar)
	route.PUT("/cars/:carId", controllers.UpdateCar)
	route.GET("/cars/:carId", controllers.GetCar)
	route.DELETE("cars/:carId", controllers.DeleteCar)

	return route
}
