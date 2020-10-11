package models

import "goRabbitMQ/config"

//OffersInput ...
type OffersInput struct {
	Offers Offers `json:"offers"`
}

//Offers ...
type Offers []struct {
	CmOfferID    string        `json:"cm_offer_id"`
	Hotel        HotelInput    `json:"hotel"`
	Room         RoomInput     `json:"room"`
	RatePlan     RatePlanInput `json:"rate_plan"`
	OriginalData OriginalData  `json:"original_data"`
	Capacity     Capacity      `json:"capacity"`
	Number       int           `json:"number"`
	Price        int           `json:"price"`
	Currency     string        `json:"currency"`
	CheckIn      string        `json:"check_in"`
	CheckOut     string        `json:"check_out"`
	Fees         Fees          `json:"fees"`
}

//OriginalData ...
type OriginalData struct {
	GuaranteePolicy GuaranteePolicy `json:"GuaranteePolicy"`
}

//GuaranteePolicy ...package models
type GuaranteePolicy struct {
	Required bool `json:"Required"`
}

//Fees ...
type Fees []struct {
	Type        string  `json:"type"`
	Description string  `json:"description"`
	Included    bool    `json:"included"`
	Percent     float64 `json:"percent"`
}

//CreateOffers ...
func CreateOffers(OffersInput *OffersInput) (err error) {
	if err = config.DB.Create(OffersInput).Error; err != nil {
		return err
	}
	return nil
}
