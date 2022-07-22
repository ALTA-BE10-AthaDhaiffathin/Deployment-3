package usecase

import (
	"errors"

	"github.com/ALTA-BE10-AthaDhaiffathin/Deployment-3.git/domain"
	"github.com/ALTA-BE10-AthaDhaiffathin/Deployment-3.git/feature/common"
)

type userUseCase struct {
	userData domain.UserData
}

func New(ud domain.UserData) domain.UserUseCase {
	return &userUseCase{
		userData: ud,
	}
}

func (ud *userUseCase) AddUser(newUser domain.User) (domain.User, error) {
	inserted := ud.userData.Insert(newUser)
	if inserted.ID == 0 {
		return domain.User{}, errors.New("cannot insert data")
	}

	return inserted, nil
}
func (ud *userUseCase) GetAll() ([]domain.User, error) {
	data := ud.userData.GetAll()

	if len(data) == 0 {
		return nil, errors.New("no data")
	}

	return data, nil
}

func (ud *userUseCase) GetSpecificUser(id uint) (domain.User, error) {
	data := ud.userData.GetSpecific(id)

	if data.ID == 0 {
		return data, errors.New("data record not found")
	}

	return data, nil
}

func (ud *userUseCase) UpdateUser(userId uint, updatedUser domain.User) (domain.User, error) {
	updatedData := ud.userData.Update(userId, updatedUser)
	if updatedData.ID == 0 {
		return domain.User{}, errors.New("cannot update user")
	}

	return updatedData, nil
}

func (ud *userUseCase) DeleteUser(userID uint) error {
	res, err := ud.userData.Delete(userID)
	if !res {
		return err
	}
	return err
}

func (ud *userUseCase) LoginUser(userData domain.User) (domain.User, string, error) {
	var token string
	data, err := ud.userData.Login(userData)
	if err != nil {
		return domain.User{}, token, err
	}
	token = common.GenerateToken(int(data.ID))
	return data, token, nil
}

func (ud *userUseCase) UserProfile(userID uint) (domain.User, error) {
	data := ud.userData.GetSpecific(userID)
	if data.ID == 0 {
		return data, errors.New("data record not found")
	}
	return data, nil
}
