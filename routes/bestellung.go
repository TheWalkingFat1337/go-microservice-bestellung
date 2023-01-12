package routes

import (
	"github.com/TheWalkingFat1337/go-microservice-bestellung/controller"
	"github.com/gin-gonic/gin"
)

func BestellungRoute(router *gin.Engine) {
	router.GET("/:id", controller.GetBestellung)
	router.GET("/", controller.GetBestellungen)
	router.POST("/", controller.CreateBestellung)
	router.DELETE("/:id", controller.DeleteBestellung)
	router.PUT("/:id", controller.UpdateBestellung)
}
