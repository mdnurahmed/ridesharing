package main

import (
	"fmt"
	"shopup/app/controllers"
	"shopup/app/models"
	"shopup/app/repositories"
	"strconv"
)

func main() {
	//initialize
	rideRepo := repositories.RideRepo{}
	userRepo := repositories.UserRepo{}
	userController := controllers.NewInstanceOfUserController(&userRepo)
	rideController := controllers.NewInstanceOfRideController(&rideRepo)
	var name, gender, age, rideid, license, seats, command string
	var ageInt, rideIdINT, seatINT, commandID int
	var user, origin, destination, vechile string
	var err error
	for {
		fmt.Println("\n\nMAIN MENU")
		fmt.Println("Type 1 for add_user")
		fmt.Println("Type 2 for add_vechile")
		fmt.Println("Type 3 for offer_ride")
		fmt.Println("Type 4 for select_ride")
		fmt.Println("Type 5 for end_ride")
		fmt.Println("Type 6 for ride status")
		fmt.Printf("Type command ID : ")
		fmt.Scanln(&command)
		if command == "main_menu" {
			continue
		}
		commandID, err = strconv.Atoi(command)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if commandID == 1 {
			fmt.Println("\nadd_user")
			fmt.Println("\n\nType main_menu for main_menu at any stage of the input\n\n")

			fmt.Printf("Name  : ")
			fmt.Scanln(&name)
			if name == "main_menu" {
				continue
			}
			fmt.Printf("Gender  : ")
			fmt.Scanln(&gender)
			if gender == "main_menu" {
				continue
			}
			fmt.Printf("Age  : ")
			fmt.Scanln(&age)
			if age == "main_menu" {
				continue
			}
			ageInt, err = strconv.Atoi(age)
			if err != nil {
				fmt.Println(err)
				continue
			}
			//------
			err = userController.AddUser(models.User{
				Name:   name,
				Gender: gender,
				Age:    ageInt,
			})
			if err == nil {
				fmt.Println("Added user")
			} else {
				fmt.Println(err)
			}
			continue
		}

		if commandID == 2 {
			fmt.Println("\nadd_vechile")
			fmt.Println("\n\nType main_menu for main_menu at any stage of the input\n\n")
			fmt.Printf("User  : ")
			fmt.Scanln(&user)
			if user == "main_menu" {
				continue
			}

			fmt.Printf("Model  : ")
			fmt.Scanln(&vechile)
			if vechile == "main_menu" {
				continue
			}

			fmt.Printf("License  : ")
			fmt.Scanln(&license)
			if license == "main_menu" {
				continue
			}
			//------
			err = userController.AddVechile(models.Vechile{
				UserName:  user,
				Model:     vechile,
				LicenseNo: license,
			})
			if err == nil {
				fmt.Println("Added vechile")
			} else {
				fmt.Println(err)
			}

			continue
		}

		if commandID == 3 {
			fmt.Println("\noffer_ride")
			fmt.Println("\n\nType main_menu for main_menu at any stage of the input\n\n")
			fmt.Printf("User  : ")
			fmt.Scanln(&user)
			if user == "main_menu" {
				continue
			}

			fmt.Printf("origin  : ")
			fmt.Scanln(&origin)
			if origin == "main_menu" {
				continue
			}

			fmt.Printf("destination  : ")
			fmt.Scanln(&destination)
			if destination == "main_menu" {
				continue
			}

			fmt.Printf("vechile  : ")
			fmt.Scanln(&vechile)
			if vechile == "main_menu" {
				continue
			}

			fmt.Printf("Seats  : ")
			fmt.Scanln(&seats)
			if seats == "main_menu" {
				continue
			}
			seatINT, err = strconv.Atoi(seats)
			if err != nil {
				fmt.Println(err)
				continue
			}

			//------
			var id int
			id, err = rideController.OfferRide(models.RideRepo{
				User:           user,
				Origin:         origin,
				Destination:    destination,
				VechileModel:   vechile,
				AvailableSeats: seatINT,
			})
			if err == nil {
				fmt.Printf("Added ride id no = %d\n", id)
			} else {
				fmt.Println(err)
			}

			continue
		}

		if commandID == 4 {
			fmt.Println("\nselect_ride")
			fmt.Println("\n\nType main_menu for main_menu at any stage of the input\n\n")
			fmt.Printf("User  : ")
			fmt.Scanln(&user)
			if user == "main_menu" {
				continue
			}

			fmt.Printf("origin  : ")
			fmt.Scanln(&origin)
			if origin == "main_menu" {
				continue
			}

			fmt.Printf("destination  : ")
			fmt.Scanln(&destination)
			if destination == "main_menu" {
				continue
			}

			fmt.Printf("preferred vechile  : ")
			fmt.Scanln(&vechile)
			if vechile == "main_menu" {
				continue
			}

			fmt.Printf("Seats  : ")
			fmt.Scanln(&seats)
			if seats == "main_menu" {
				continue
			}
			seatINT, err = strconv.Atoi(seats)
			if err != nil {
				fmt.Println(err)
				continue
			}
			//----

			res, er := rideController.FindRide(models.RideRepo{
				User:           user,
				Origin:         origin,
				Destination:    destination,
				VechileModel:   vechile,
				AvailableSeats: seatINT,
			})
			if er == nil {
				fmt.Printf("found ride \n %+v\n", res)
			} else {
				fmt.Println(er)
			}

			continue

		}

		if commandID == 5 {
			fmt.Println("\nend_ride")
			fmt.Println("\n\nType main_menu for main_menu at any stage of the input\n\n")
			fmt.Printf("Ride ID  : ")
			fmt.Scanln(&rideid)
			if destination == "main_menu" {
				continue
			}
			rideIdINT, err = strconv.Atoi(rideid)
			if err != nil {
				fmt.Println(err)
				continue
			}
			err = rideController.EndRide(rideIdINT)
			if err == nil {
				fmt.Println("RIDE ENDED")
			} else {
				fmt.Println(err)
			}
		}

		if commandID == 6 {
			fmt.Println("\nride_status")
			for _, element := range repositories.UserCollection {
				fmt.Printf("%s : taken : %d , offered : %d \n", element.Name, element.RideTaken, element.RideOffered)
			}

		}
	}

}
