package models

import (
	"goRabbitMQ/config"

	"github.com/sirupsen/logrus"
)

/*
	create table query ...

	CREATE TABLE hotel (
		id int(11) NOT NULL AUTO_INCREMENT,
		hotel_id varchar(20) NOT NULL,
		name  varchar(100) NULL DEFAULT NULL,
		country varchar(40) NULL DEFAULT NULL,
		address  varchar(40) NULL DEFAULT NULL,
		latitude  DOUBLE(16,6) NULL DEFAULT NULL,
		longitude DOUBLE(16,6) NULL DEFAULT NULL,
		telephone varchar(20) NULL DEFAULT NULL,
		amenities varchar(255) NULL DEFAULT NULL,
		description varchar(255) NULL DEFAULT NULL,
		room_count int(11) NULL DEFAULT NULL,
		currency varchar(10) NULL DEFAULT NULL,
		row_created timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
		row_updated timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
		PRIMARY KEY (hotel_id)
	)

*/
//Hotel ...
type HotelInput struct {
	HotelID     string   `json:"hotel_id"`
	Name        string   `json:"name"`
	Country     string   `json:"country"`
	Address     string   `json:"address"`
	Latitude    float64  `json:"latitude"`
	Longitude   float64  `json:"longitude"`
	Telephone   string   `json:"telephone"`
	Amenities   []string `json:"amenities"`
	Description string   `json:"description"`
	RoomCount   int      `json:"room_count"`
	Currency    string   `json:"currency"`
}

//HotelInput ...
type Hotel struct {
	HotelID     string  `gorm:"primary_key" json:"hotel_id"`
	Name        string  `json:"name"`
	Country     string  `json:"country"`
	Address     string  `json:"address"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Telephone   string  `json:"telephone"`
	Amenities   string  `json:"amenities"`
	Description string  `json:"description"`
	RoomCount   int     `json:"room_count"`
	Currency    string  `json:"currency"`
}

//StoreHotel ...
func StoreHotel(input HotelInput) error {
	AmenitiesString := ""
	for index, val := range input.Amenities {
		if index > 0 {
			AmenitiesString += " , "
		}
		AmenitiesString += val
	}
	var data = Hotel{
		HotelID:     input.HotelID,
		Name:        input.Name,
		Country:     input.Country,
		Address:     input.Address,
		Latitude:    input.Latitude,
		Longitude:   input.Longitude,
		Telephone:   input.Telephone,
		Amenities:   AmenitiesString,
		Description: input.Description,
		RoomCount:   input.RoomCount,
		Currency:    input.Currency,
	}
	logrus.WithFields(logrus.Fields{
		"HotelID":     input.HotelID,
		"Name":        input.Name,
		"Country":     input.Country,
		"Address":     input.Address,
		"Latitude":    input.Latitude,
		"Longitude":   input.Longitude,
		"Telephone":   input.Telephone,
		"Amenities":   AmenitiesString,
		"Description": input.Description,
		"RoomCount":   input.RoomCount,
		"Currency":    input.Currency,
	}).Info("inserting a row in mysql table - hotels")
	if err := config.DB.Create(data).Error; err != nil {
		return err
	}
	return nil
}
