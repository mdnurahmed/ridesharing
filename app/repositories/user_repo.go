package repositories

import (
	"fmt"
	"shopup/app/models"
)

var UserCollection []models.UserRepo

type IUserRepo interface {
	AddUser(models.User) error
	AddVechile(models.Vechile) error
}

type UserRepo struct {
}

func (u *UserRepo ) AddUser(user models.User) error {
	var err error 
for _,element := range UserCollection{
	if element.Name == user.Name {
		err = fmt.Errorf("username already exists")
		break
	}
}
if err != nil {
	return err
}
newuser := models.UserRepo{
	Name : user.Name,
	Age : user.Age,
	Gender : user.Gender,
}
UserCollection = append(UserCollection,newuser)
return err
}



func (u *UserRepo ) AddVechile(vechile models.Vechile)error {
	var err error 
userIndex := -1

for index,element := range UserCollection{
	if element.Name == vechile.UserName {
		userIndex = index
		break
	}
}
if userIndex == -1 {
	err = fmt.Errorf("user Not found")
	return err
}
newvechile := models.VechileDescriptions{
	Model : vechile.Model,
	LicenseNo : vechile.LicenseNo,
}
UserCollection[userIndex].Vechiles = append(UserCollection[userIndex].Vechiles,newvechile)
return err
}