package domain

type Book struct {
	ID       int    `json:"id" form:"id"`
	Judul    string `json:"judul" form:"judul"`
	Penerbit string `json:"penerbit" form:"penerbit"`
	ISBN     string
	User_ID  int
}

type BookUseCase interface {
	AddBook(newBook Book) (Book, error)
	GetAll() ([]Book, error)
	GetSpecificBook(id uint) (Book, error)
	UpdateBook(userID uint, bookID uint, updatedBook Book) (Book, error)
	DeleteBook(userID uint, bookID uint) error
}

type BookData interface {
	Insert(newBook Book) (Book, error)
	GetAll() ([]Book, error)
	GetSpecific(bookID uint) (Book, error)
	Update(userID uint, bookID uint, updatedData Book) (Book, error)
	Delete(userID uint, bookID uint) (bool, error)
}
