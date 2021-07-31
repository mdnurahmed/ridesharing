package repositories

import (
	"shopup/app/models"
	"testing"
)

func TestAddRide_HappyCase(t *testing.T) {
	riderepo := RideRepo{}
	_, err := riderepo.AddRide(models.RideRepo{
		User:             "john",
		Origin:           "Delhi",
		Destination:      "Dhaka",
		VechileModel:     "Corolla",
		VechileLicenseNo: "A123",
		AvailableSeats:   4,
	})
	if err == nil {
		t.Error("Expected user deosnt exist error but got nil ")
	}
	if len(RideCollection) != 0 {
		t.Errorf("Expected RideCollection length to be 1 but got - %d", len(RideCollection))
	}

}
