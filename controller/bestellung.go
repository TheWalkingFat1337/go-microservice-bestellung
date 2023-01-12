package controller

import (
	"time"

	"github.com/TheWalkingFat1337/go-microservice-bestellung/config"
	"github.com/TheWalkingFat1337/go-microservice-bestellung/models"
	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
)

func GetBestellungen(c *gin.Context) {
	bestellungen := []models.Bestellung{}
	config.DB.Find(&bestellungen)
	c.JSON(200, &bestellungen)
}

func GetBestellung(c *gin.Context) {
	var bestellung models.Bestellung
	config.DB.Where("id = ?", c.Param("id")).Find(&bestellung)
	c.JSON(200, &bestellung)
}

func CreateBestellung(c *gin.Context) {
	//var name string = http.GET("http://restapi_user:1234/name/")
	bestellung := models.Bestellung{
		Bestelldatum: datatypes.Date(time.Now()),
		//Kundennr: name
	}
	c.BindJSON(&bestellung)
	config.DB.Create(&bestellung)
	c.JSON(200, &bestellung)
}

func DeleteBestellung(c *gin.Context) {
	var bestellung models.Bestellung
	config.DB.Where("id = ?", c.Param("id")).Delete(&bestellung)
	c.JSON(200, &bestellung)
}

func UpdateBestellung(c *gin.Context) {
	var bestellung models.Bestellung
	config.DB.Where("id = ?", c.Param("id")).First(&bestellung)
	c.BindJSON(&bestellung)
	config.DB.Save(&bestellung)
	c.JSON(200, &bestellung)
}
