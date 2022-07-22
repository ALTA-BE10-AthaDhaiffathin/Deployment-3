package data

import (
	"github.com/ALTA-BE10-AthaDhaiffathin/Deployment-3.git/domain"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Judul    string `json:"judul" form:"judul"`
	Penerbit string `json:"penerbit" form:"penerbit"`
	ISBN     string
	User_ID  int `json:"user_id" form:"user_id"`
}

func (b *Book) ToDomain() domain.Book {
	return domain.Book{
		ID:       int(b.ID),
		Judul:    b.Judul,
		Penerbit: b.Penerbit,
		ISBN:     b.ISBN,
		User_ID:  b.User_ID,
	}
}

func ToEntity(newBook domain.Book) Book {
	return Book{
		Judul:    newBook.Judul,
		Penerbit: newBook.Penerbit,
		ISBN:     newBook.ISBN,
		User_ID:  newBook.User_ID,
	}
}
