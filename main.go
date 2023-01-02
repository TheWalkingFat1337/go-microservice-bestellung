package main

import (
	"github.com/TheWalkingFat1337/go-microservice-bestellung/config"
	"github.com/TheWalkingFat1337/go-microservice-bestellung/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	config.Connect()
	routes.BestellungRoute(router)
	router.Run(":5678")
}
