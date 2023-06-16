package router

import (
	"github.com/gin-gonic/gin"
	"github.com/travor-backend/controller"
	"github.com/travor-backend/db"
)

func destinationRoutes(router *gin.RouterGroup) {
	destinations := router.Group("/destinations")

	{
		destinations.GET("/", controller.GetDestinations(db.DB))
		destinations.GET("/:id", controller.GetDestinationById(db.DB))
		destinations.POST("/", controller.CreateDestination(db.DB))
		destinations.PUT("/:id", controller.UpdateDestination(db.DB))
		destinations.DELETE("/:id", controller.DeleteDestination(db.DB))
	}
}
