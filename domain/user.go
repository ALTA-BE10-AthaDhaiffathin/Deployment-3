package domain

type User struct {
	ID       uint
	Nama     string
	Email    string
	Password string
}

type UserUseCase interface {
	AddUser(newUser User) (User, error)
	GetAll() ([]User, error)
	GetSpecificUser(id uint) (User, error)
	UpdateUser(userID uint, updatedUser User) (User, error)
	DeleteUser(userID uint) error
	LoginUser(userData User) (User, string, error)
	UserProfile(userID uint) (User, error)
}

type UserData interface {
	Insert(newUser User) User
	GetAll() []User
	GetSpecific(userID uint) User
	Update(userID uint, updatedData User) User
	Delete(userID uint) (bool, error)
	Login(userData User) (User, error)
}
