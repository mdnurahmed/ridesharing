package models

type RideRepo struct {
	User             string
	Origin           string
	Destination      string
	VechileModel     string
	VechileLicenseNo string
	AvailableSeats   int
}
