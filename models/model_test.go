package models

import (
	"fmt"
	"goRabbitMQ/config"
	"strconv"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var testNo = 3 // please increase the index while testing
var (
	HotelID    = "hotel_test_" + strconv.Itoa(testNo)
	RoomID     = "room_test_" + strconv.Itoa(testNo)
	RatePlanID = "rate_paln_" + strconv.Itoa(testNo)
)

func TestStoreHotel(t *testing.T) {
	var err error

	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	err = StoreHotel(HotelInput{
		HotelID:   HotelID,
		Name:      "Hawthorn Suites by Wyndham Eagle CO",
		Country:   "US",
		Address:   "0315 Chambers Avenue, 81631",
		Latitude:  39.660193,
		Longitude: -106.824123,
		Telephone: "+1-970-3283000",
		Amenities: []string{
			"Business Centre",
			"Fitness Room/Gym",
			"Pet Friendly",
			"Disabled Access",
			"Air Conditioned",
			"Free WIFI",
			"Elevator / Lift",
			"Parking",
		},
		Description: "Stay a while in beautiful mountain country at this Hawthorn Suites by Wyndham Eagle CO hotel, just off Interstate 70, only 6 miles from the Vail/Eagle Airport and close to skiing, golfing, Eagle River and great restaurants. Pets are welcome at this h",
		RoomCount:   1,
		Currency:    "USD",
	})
	if err != nil {
		fmt.Println(err)
	}
	defer config.DB.Close()

}

func TestStoreRoom(t *testing.T) {
	var err error

	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	err = StoreRoom(RoomInput{
		HotelID:     HotelID,
		RoomID:      RoomID,
		Description: "JUNIOR SUITES WITH 2 QUEEN BEDS",
		Name:        "JUNIOR SUITES WITH 2 QUEEN BEDS",
		Capacity: Capacity{
			MaxAdults:     2,
			ExtraChildren: 2,
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	defer config.DB.Close()

}

func TestStoreRatePlan(t *testing.T) {
	var err error

	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	err = StoreRatePlan(RatePlanInput{
		HotelID:    HotelID,
		RatePlanID: RatePlanID,
		CancellationPolicy: CancellationPolicy{
			{
				Type:              "Free cancellation",
				ExpiresDaysBefore: 2,
			},
		},
		Name: "BEST AVAILABLE RATE",
		OtherConditions: []string{
			"CXL BY 2 DAYS PRIOR TO ARRIVAL-FEE 1 NIGHT 2 DAYS PRIOR TO ARRIVAL",
			"BEST AVAILABLE RATE",
		},
		MealPlan: "Room only",
	})
	if err != nil {
		fmt.Println(err)
	}
	defer config.DB.Close()

}
