package routes

import (
	"testing"
)

var err error

func TestStoreHotel(t *testing.T) {

	r := SetupRouter()
	//running
	r.Run()
}
