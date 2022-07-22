package data

import (
	"errors"

	"github.com/ALTA-BE10-AthaDhaiffathin/Deployment-3.git/domain"
	"gorm.io/gorm"
)

type bookData struct {
	db *gorm.DB
}

func New(DB *gorm.DB) domain.BookData {
	return &bookData{
		db: DB,
	}
}

func (bd *bookData) Insert(newBook domain.Book) (domain.Book, error) {
	var bookData Book = ToEntity(newBook)
	err := bd.db.Create(&bookData).Error
	if err != nil {
		return domain.Book{}, errors.New("error create book")
	}

	return bookData.ToDomain(), nil
}

func (bd *bookData) GetAll() ([]domain.Book, error) {
	var books []Book
	err := bd.db.Find(&books).Error
	if err != nil {
		return []domain.Book{}, errors.New(err.Error())
	}

	var res []domain.Book
	for i := 0; i < len(books); i++ {
		res = append(res, books[i].ToDomain())
	}

	return res, nil
}

func (bd *bookData) GetSpecific(bookID uint) (domain.Book, error) {
	var bookData Book
	err := bd.db.Where("id = ?", bookID).First(&bookData).Error
	if err != nil {
		return domain.Book{}, errors.New(err.Error())
	}
	return bookData.ToDomain(), nil
}

func (bd *bookData) Update(userID uint, bookID uint, updatedData domain.Book) (domain.Book, error) {
	var currentBookData Book
	err := bd.db.Where("id = ? and user_id = ?", bookID, userID).First(&currentBookData).Error
	if err != nil {
		return domain.Book{}, errors.New("error " + err.Error() + " , you have no book with that id")
	}

	err = bd.db.Model(&Book{}).Where("ID = ?", bookID).Updates(ToEntity(updatedData)).Error
	if err != nil {
		return domain.Book{}, errors.New(err.Error())
	}

	return updatedData, nil
}

func (bd *bookData) Delete(userID uint, bookID uint) (bool, error) {
	var bookData Book
	err := bd.db.Where("id = ? and user_id = ?", bookID, userID).First(&bookData).Error
	if err != nil {
		return false, errors.New("you dont have book with that id")
	}
	err = bd.db.Delete(&bookData).Error
	if err != nil {
		return false, errors.New("error occured " + err.Error())
	}
	return true, nil
}
