package helpers

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type RequestCallbackSSO struct {
	Id            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail string `json:"verified_email"`
	Picture       string `json:"picture"`
}

var OauthConfig = &oauth2.Config{
	RedirectURL: "http://localhost:5173/auth/sso-response-callback",
	Scopes: []string{
		"https://www.googleapis.com/auth/userinfo.email",
	},
	Endpoint: google.Endpoint,
}

func GetSSOGoogleUrl() string {
	url := OauthConfig.AuthCodeURL("state")
	return url
}

func CallbackSSO(code string) (RequestCallbackSSO, error) {
	var dataSSO RequestCallbackSSO
	token, err := OauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return dataSSO, err
	}

	client := OauthConfig.Client(context.Background(), token)
	response, err := client.Get("https://www.googleapis.com/userinfo/v2/me")
	if err != nil {
		return dataSSO, err
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return dataSSO, err
	}

	err1 := json.Unmarshal(data, &dataSSO)
	if err1 != nil {
		return dataSSO, err1
	}

	return dataSSO, nil
}
