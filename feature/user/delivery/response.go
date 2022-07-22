package delivery

import "github.com/ALTA-BE10-AthaDhaiffathin/Deployment-3.git/domain"

type UserResponse struct {
	ID    uint
	Nama  string
	Email string
}

func ToUserResponse(data domain.User) UserResponse {
	return UserResponse{
		ID:    data.ID,
		Nama:  data.Nama,
		Email: data.Email,
	}
}

type LoginResponse struct {
	ID    int
	Nama  string
	Email string
	Token string
}
