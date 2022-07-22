package data

import (
	"github.com/ALTA-BE10-AthaDhaiffathin/Deployment-3.git/domain"
	"github.com/ALTA-BE10-AthaDhaiffathin/Deployment-3.git/feature/book/data"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nama     string      `json:"nama" form:"nama"`
	Email    string      `json:"email" form:"email"`
	Password string      `json:"password" form:"password"`
	Book     []data.Book `gorm:"foreignKey:User_ID"`
}

func (u *User) ToModel() domain.User {
	return domain.User{
		ID:       u.ID,
		Nama:     u.Nama,
		Email:    u.Email,
		Password: u.Password,
	}
}

func ParseToArr(arr []User) []domain.User {
	var res []domain.User

	for _, val := range arr {
		res = append(res, val.ToModel())
	}

	return res
}

func FromModel(data domain.User) User {
	var res User
	res.Email = data.Email
	res.Nama = data.Nama
	res.Password = data.Password
	res.ID = uint(data.ID)
	return res
}
