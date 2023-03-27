package delivery

import "lapakUmkm/features/users"

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type ChangePasswordRequest struct {
	OldPassword     string `json:"old_password" form:"old_password"`
	NewPassword     string `json:"new_password" form:"new_password"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password"`
}

type CallbackSSORequest struct {
	Id            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Picture       string `json:"picture"`
}

func CallbackSSORequestToUserEntity(c CallbackSSORequest) users.UserEntity {
	return users.UserEntity{
		Email:        c.Email,
		PhotoProfile: c.Picture,
	}
}
