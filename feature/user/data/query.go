package data

import (
	"errors"
	"log"

	"github.com/ALTA-BE10-AthaDhaiffathin/Deployment-3.git/domain"
	"gorm.io/gorm"
)

type userData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.UserData {
	return &userData{
		db: db,
	}
}

func (ud *userData) Insert(newUser domain.User) domain.User {
	var cnv = FromModel(newUser)
	err := ud.db.Create(&cnv).Error
	if err != nil {
		log.Println("Cannot create object", err.Error())
		return domain.User{}
	}

	return cnv.ToModel()
}

func (ud *userData) Update(userID uint, updatedData domain.User) domain.User {
	var cnv = FromModel(updatedData)
	err := ud.db.Model(&User{}).Where("ID = ?", userID).Updates(updatedData).Error
	if err != nil {
		log.Println("Cannot update data", err.Error())
		return domain.User{}
	}
	cnv.ID = userID
	return cnv.ToModel()
}

func (ud *userData) Delete(userID uint) (bool, error) {
	var userData User
	res := ud.db.Where("ID = ?", userID).First(&userData)
	if res.RowsAffected < 1 {
		log.Println("No data deleted", res.Error.Error())
		return false, res.Error
	}
	ud.db.Delete(&userData)
	return true, nil
}

func (ud *userData) GetAll() []domain.User {
	var tmp []User
	err := ud.db.Find(&tmp).Error

	if err != nil {
		log.Println("Cannot retrive object", err.Error())
		return nil
	}
	return ParseToArr(tmp)
}
func (ud *userData) GetSpecific(userID uint) domain.User {
	var tmp User
	err := ud.db.Where("ID = ?", userID).First(&tmp).Error
	if err != nil {
		log.Println("There is a problem with data", err.Error())
		return domain.User{}
	}

	return tmp.ToModel()
}

func (ud *userData) Login(userData domain.User) (domain.User, error) {
	var userLoginData = User{}
	err := ud.db.Where("email = ?", userData.Email).First(&userLoginData).Error
	if err != nil {
		return domain.User{}, errors.New("email not registered")
	}

	err = ud.db.Where("email = ? and password = ?", userData.Email, userData.Password).First(&userLoginData).Error
	if err != nil {
		return domain.User{}, errors.New("wrong password")
	}

	return userLoginData.ToModel(), nil
}
