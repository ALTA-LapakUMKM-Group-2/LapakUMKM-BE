package delivery

import (
	"lapakUmkm/app/middlewares"
	"lapakUmkm/features/auth"
	"lapakUmkm/features/users/delivery"
	"lapakUmkm/utils/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	Service auth.AuthServiceInterface
}

func New(s auth.AuthServiceInterface) *AuthHandler {
	return &AuthHandler{
		Service: s,
	}
}

func (u *AuthHandler) Login(c echo.Context) error {
	loginRequest := LoginRequest{}
	if err := c.Bind(&loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail("error bind data"))
	}

	_, user, err := u.Service.Login(loginRequest.Email, loginRequest.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, helpers.ResponseFail(err.Error()))
	}

	token, _ := middlewares.CreateToken(int(user.Id), user.Role)

	tokesResponse := map[string]any{
		"token": token,
		"user":  delivery.UserEntityToUserResponse(user),
	}

	return c.JSON(http.StatusOK, helpers.ResponseSuccess("login success", tokesResponse))
}

func (h *AuthHandler) Register(c echo.Context) error {
	registerRequest := delivery.UserRequest{}
	if err := c.Bind(&registerRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail("error bind data"))
	}

	registerRequest.Role = "user"

	user := delivery.UserRequestToUserEntity(registerRequest)
	if err := h.Service.Register(user); err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFail(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.ResponseSuccess("register success", delivery.UserEntityToUserResponse(user)))
}

func (u *AuthHandler) ChangePassword(c echo.Context) error {
	user_id := middlewares.ClaimsToken(c).Id

	r := ChangePasswordRequest{}
	if err := c.Bind(&r); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail("rrror bind data"))
	}

	hash, _ := helpers.HashPassword(r.NewPassword)
	if err := u.Service.ChangePassword(uint(user_id), r.OldPassword, r.NewPassword, r.ConfirmPassword, hash); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.ResponseSuccess("change password Success", nil))
}

func (u *AuthHandler) LoginSSOGoogle(c echo.Context) error {
	callbackSSORequest := CallbackSSORequest{}
	if err := c.Bind(&callbackSSORequest); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail("error bind data"))
	}

	if !callbackSSORequest.VerifiedEmail {
		return c.JSON(http.StatusUnauthorized, helpers.ResponseFail("email not verified yet at google"))
	}

	maping := CallbackSSORequestToUserEntity(callbackSSORequest)
	_, user, err := u.Service.LoginSSOGoogle(maping)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, helpers.ResponseFail(err.Error()))
	}

	tokens, _ := middlewares.CreateToken(int(user.Id), user.Role)

	tokesResponse := map[string]any{
		"token": tokens,
		"user":  delivery.UserEntityToUserResponse(user),
	}

	return c.JSON(http.StatusOK, helpers.ResponseSuccess("login success", tokesResponse))
}

func (u *AuthHandler) ForgetPassword(c echo.Context) error {
	email := c.QueryParam("email")

	if err := u.Service.ForgetPassword(email); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail(err.Error()))
	}

	token := helpers.EncryptText(email)
	urlLink := "https://lapakumkm.netlify.app/new-password?token=" + token

	//send URL to Email
	if errSendmail := helpers.SendMail("Forget Password", email, urlLink); errSendmail != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFail(errSendmail.Error()))
	}

	return c.JSON(http.StatusOK, helpers.ResponseSuccess("email was sended", nil))
}

func (u *AuthHandler) IsUserExist(c echo.Context) error {
	email := c.QueryParam("email")

	if err := u.Service.IsUserExist(email); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.ResponseSuccess("email found", email))
}

func (u *AuthHandler) NewPassword(c echo.Context) error {
	req := NewPasswordRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail("error bind data"))
	}

	if err := u.Service.NewPassword(req.Token, req.NewPassword, req.ConfirmPassword); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.ResponseSuccess("new password success", nil))
}
