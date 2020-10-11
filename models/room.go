package models

import (
	"goRabbitMQ/config"
	"strconv"
)

/*
	CREATE TABLE `rooms` (
		`hotel_id` varchar(255) DEFAULT NULL,
		`room_id` varchar(255) DEFAULT NULL,
		`description` varchar(255) DEFAULT NULL,
		`name` varchar(255) DEFAULT NULL,
		`capacity` varchar(255) DEFAULT NULL
	)
*/
//Room ...
type RoomInput struct {
	HotelID     string   `json:"hotel_id"`
	RoomID      string   `json:"room_id"`
	Description string   `json:"description"`
	Name        string   `json:"name"`
	Capacity    Capacity `json:"capacity"`
}

//RoomInput ...
type Room struct {
	HotelID     string `gorm:"type:varchar(255); not null" json:"hotel_id"`
	RoomID      string `gorm:"primary_key" json:"room_id"`
	Description string `json:"description"`
	Name        string `json:"name"`
	Capacity    string `json:"capacity"`
}

//Capacity ...
type Capacity struct {
	MaxAdults     int `json:"max_adults"`
	ExtraChildren int `json:"extra_children"`
}

//StoreRoom ...
func StoreRoom(input RoomInput) error {
	var capacityStr = "max_adults:" + strconv.Itoa(input.Capacity.MaxAdults) + ", extra_children:" + strconv.Itoa(input.Capacity.ExtraChildren)

	data := Room{
		HotelID:     input.HotelID,
		RoomID:      input.RoomID,
		Description: input.Description,
		Name:        input.Name,
		Capacity:    capacityStr,
	}
	if err := config.DB.Create(data).Error; err != nil {
		return err
	}
	return nil

}
