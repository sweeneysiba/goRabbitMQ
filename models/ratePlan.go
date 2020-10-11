package models

import (
	"goRabbitMQ/config"
	"strconv"

	"github.com/sirupsen/logrus"
)

//RatePlan ...
type RatePlanInput struct {
	HotelID            string             `json:"hotel_id"`
	RatePlanID         string             `json:"rate_plan_id"`
	CancellationPolicy CancellationPolicy `json:"cancellation_policy"`
	Name               string             `json:"name"`
	OtherConditions    []string           `json:"other_conditions"`
	MealPlan           string             `json:"meal_plan"`
}

//RatePlanInput ...
type RatePlan struct {
	HotelID            string `gorm:"type:varchar(255); not null" json:"hotel_id"`
	RatePlanID         string `gorm:"primary_key"  json:"rate_plan_id"`
	CancellationPolicy string `json:"cancellation_policy"`
	Name               string `json:"name"`
	OtherConditions    string `json:"other_conditions"`
	MealPlan           string `json:"meal_plan"`
}

//CancellationPolicy ...
type CancellationPolicy []struct {
	Type              string `json:"type"`
	ExpiresDaysBefore int    `json:"expires_days_before"`
}

/*
	create table query ...
	CREATE TABLE rate_plans (
  		hotel_id varchar(255) DEFAULT NULL,
  		rate_plan_id varchar(255) DEFAULT NULL,
  		cancellation_policy varchar(255) DEFAULT NULL,
  		name varchar(255) DEFAULT NULL,
  		other_conditions varchar(255) DEFAULT NULL,
  		meal_plan varchar(255) DEFAULT NULL,
		PRIMARY KEY (hotel_id)
		FOREIGN KEY (hotel_id) REFERENCES hotel(hotel_id)
	)
*/
//RatePlan...
func StoreRatePlan(input RatePlanInput) error {
	var cancellationPolicyStr = ""

	var otherConditionsStr = ""
	for index, val := range input.OtherConditions {
		if index > 0 {
			otherConditionsStr += " , "
		}
		otherConditionsStr += val
	}

	for index, val := range input.CancellationPolicy {
		if index > 0 {
			cancellationPolicyStr += " ; "
		}
		var capacityStr = "type:" + val.Type + ", expires_days_before:" + strconv.Itoa(val.ExpiresDaysBefore)
		cancellationPolicyStr += capacityStr
	}
	data := RatePlan{
		HotelID:            input.HotelID,
		RatePlanID:         input.RatePlanID,
		CancellationPolicy: cancellationPolicyStr,
		Name:               input.Name,
		OtherConditions:    otherConditionsStr,
		MealPlan:           input.MealPlan,
	}
	logrus.WithFields(logrus.Fields{
		"HotelID":            input.HotelID,
		"RatePlanID":         input.RatePlanID,
		"CancellationPolicy": cancellationPolicyStr,
		"Name":               input.Name,
		"OtherConditions":    otherConditionsStr,
		"MealPlan":           input.MealPlan,
	}).Info("inserting a row in mysql table - rate_plans")
	if err := config.DB.Create(data).Error; err != nil {
		return err
	}
	return nil
}
