package models

type VechileDescriptions struct {
	Model     string
	LicenseNo string
}

type UserRepo struct {
	Name        string
	Gender      string
	Age         int
	Vechiles    []VechileDescriptions
	RideTaken   int
	RideOffered int
}
