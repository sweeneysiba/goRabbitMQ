package controller

import (
	"fmt"
	"goRabbitMQ/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//CreateUser ... Create User
func CreateOffers(c *gin.Context) {
	var offers models.OffersInput
	c.BindJSON(&offers)
	err := models.CreateOffers(&offers)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, offers)
	}
}
