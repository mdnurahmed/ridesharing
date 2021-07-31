package controllers

import (
	"shopup/app/models"
	"shopup/app/repositories"
)

type IRideController interface {
	OfferRide(models.RideRepo) error 
	FindRide(models.RideRepo) (models.RideRepo,error) 
	EndRide(RideId int) error 
}


type RideController struct {
	rideRepo repositories.IRideRepo
}


func NewInstanceOfRideController(rideRepo repositories.IRideRepo)  RideController {
	return RideController{rideRepo : rideRepo }
}

func (r *RideController) OfferRide(ride models.RideRepo) (int,error)  {
	rideId,err := r.rideRepo.AddRide(ride)
	return rideId,err 
}


func (r *RideController) FindRide(ride models.RideRepo) (models.RideRepo,error){
	res,err := r.rideRepo.FindRide(ride)
	return res,err 
}

func (r *RideController) EndRide(RideId int) error {
	err := r.rideRepo.EndRide(RideId)
	return err 
}