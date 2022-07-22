package delivery

import "github.com/ALTA-BE10-AthaDhaiffathin/Deployment-3.git/domain"

type InsertFormat struct {
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (i *InsertFormat) ToModel() domain.User {
	return domain.User{
		Nama:     i.Nama,
		Email:    i.Email,
		Password: i.Password,
	}
}

type LoginFormat struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (l *LoginFormat) ToModel() domain.User {
	return domain.User{
		Email:    l.Email,
		Password: l.Password,
	}
}
