package usecase

import (
	"github.com/ALTA-BE10-AthaDhaiffathin/Deployment-3.git/domain"
	"github.com/google/uuid"
)

type bookUseCase struct {
	bookData domain.BookData
}

func New(bd domain.BookData) domain.BookUseCase {
	return &bookUseCase{
		bookData: bd,
	}
}

func (bu *bookUseCase) AddBook(newBook domain.Book) (domain.Book, error) {
	newBook.ISBN = uuid.NewString()
	data, err := bu.bookData.Insert(newBook)
	if err != nil {
		return domain.Book{}, err
	}
	return data, nil
}

func (bu *bookUseCase) GetAll() ([]domain.Book, error) {
	data, err := bu.bookData.GetAll()
	if err != nil {
		return data, err
	}
	return data, nil
}

func (bu *bookUseCase) GetSpecificBook(id uint) (domain.Book, error) {
	data, err := bu.bookData.GetSpecific(id)
	if err != nil {
		return data, err
	}
	return data, nil
}

func (bu *bookUseCase) UpdateBook(userID uint, bookID uint, updatedData domain.Book) (domain.Book, error) {
	data, err := bu.bookData.Update(userID, bookID, updatedData)
	if err != nil {
		return data, err
	}
	return data, nil
}

func (bu *bookUseCase) DeleteBook(userID uint, bookID uint) error {
	_, err := bu.bookData.Delete(userID, bookID)
	if err != nil {
		return err
	}
	return nil
}
