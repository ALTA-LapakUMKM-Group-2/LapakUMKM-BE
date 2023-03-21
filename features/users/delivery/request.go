package delivery

import "lapakUmkm/features/users"

type UserRequest struct {
	FullName    string `json:"full_name" form:"full_name"`
	Email       string `json:"email" form:"email"`
	Password    string `json:"password" form:"password"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
	Address     string `json:"address" form:"address"`
	Role        string `json:"role" form:"role"`
}

func UserRequestToUserEntity(userRequest UserRequest) users.UserEntity {
	return users.UserEntity{
		FullName:    userRequest.FullName,
		Email:       userRequest.Email,
		Password:    userRequest.Password,
		PhoneNumber: userRequest.PhoneNumber,
		Address:     userRequest.Address,
		Role:        userRequest.Role,
	}
}
