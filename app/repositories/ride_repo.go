package repositories

import (
	"fmt"
	"shopup/app/models"
)

var RideCollection []models.RideRepo

type IRideRepo interface {
	AddRide(models.RideRepo) (int,error)
	FindRide(models.RideRepo) (models.RideRepo,error)
	EndRide(RideID int) error 
}

type RideRepo struct {
}

func (r *RideRepo  ) AddRide(ride models.RideRepo) (int,error) {
	var err error 
	userExists := false
	carExists := false 
	userIndex := -1
	for index,element := range UserCollection{
		if element.Name == ride.User {
			userExists = true 
			userIndex = index
			break
		}
	}
	if !userExists {
		err =fmt.Errorf("no such user exists")
		return -1,err
	}
	for _,car := range UserCollection[userIndex].Vechiles {
		if car.Model == ride.VechileModel{
			carExists = true 
			break
		}
	}
	if !carExists {
		err =fmt.Errorf("user doesnt have this car")
		return -1,err
	}
	alreadyOnaRide := false 
	for _,element := range RideCollection{
		if element.User == ride.User {
			alreadyOnaRide = true 
			break
		}
	}
	if alreadyOnaRide {
		err =fmt.Errorf("user alrready on a active ride")
		return -1,err
	}
	UserCollection[userIndex].RideOffered++
	RideCollection = append(RideCollection, ride)
	return len(RideCollection)-1,err 

}



func (r *RideRepo)FindRide(ride models.RideRepo) (models.RideRepo,error){
	var err error 
	userExists := false
	userIndex := -1
	for index,element := range UserCollection{
		if element.Name == ride.User {
			userExists = true 
			userIndex = index
			break
		}
	}
	if !userExists {
		err =fmt.Errorf("no such user exists")
		return models.RideRepo{},err
	}


	if ride.AvailableSeats >2 {
		return  models.RideRepo{},fmt.Errorf("cannot offer more than 2 seats ")
	}
	UserCollection[userIndex].RideTaken++
	for _,element := range RideCollection{
		if element.Origin == ride.Origin && 
		element.Destination == ride.Destination && 
		element.AvailableSeats >= ride.AvailableSeats && 
		element.VechileModel == ride.VechileModel{
			element.AvailableSeats = element.AvailableSeats-ride.AvailableSeats
			return element,nil
		}
	}
	return models.RideRepo{},fmt.Errorf("no ride found")
}



func (r *RideRepo)EndRide(RideId int ) error{
	if RideId > len(RideCollection) || RideCollection[RideId].User==""{
		return fmt.Errorf("ride Doesnt exists")
	} 
	RideCollection[RideId] = models.RideRepo{}
	return nil 
}