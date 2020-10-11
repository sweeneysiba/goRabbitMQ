//main.go
package main

import (
	"fmt"
	"goRabbitMQ/config"
	"goRabbitMQ/controller"
	"goRabbitMQ/models"
	Routes "goRabbitMQ/routes"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer config.DB.Close()
	config.DB.AutoMigrate(&models.Hotel{}, &models.Room{}, &models.RatePlan{})
	config.DB.Model(&models.Room{}).AddForeignKey("hotel_id", "hotels(hotel_id)", "RESTRICT", "RESTRICT")
	config.DB.Model(&models.RatePlan{}).AddForeignKey("hotel_id", "hotels(hotel_id)", "RESTRICT", "RESTRICT")
	go controller.QueueConsumer()

	r := Routes.SetupRouter()
	//running
	r.Run()
}
