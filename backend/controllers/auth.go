package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"delta.nitt.edu/dion/config"
	"delta.nitt.edu/dion/models"
	"delta.nitt.edu/dion/repository"
	"golang.org/x/oauth2"
)

type DAuthUserResponse struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

func GetUser(email string) (models.User, error) {
	return repository.GetUser(email)
}

func HandleCallBack(code string) (string, error) {
	oauthClient := oauth2.Config{
		RedirectURL:  config.C.OauthConfig.RedirectURL,
		ClientSecret: config.C.OauthConfig.ClientSecret,
		ClientID:     config.C.OauthConfig.ClientId,
		Scopes:       []string{"email"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://auth.delta.nitt.edu/authorize",
			TokenURL: "https://auth.delta.nitt.edu/api/oauth/token",
		},
	}
	token, err := oauthClient.Exchange(oauth2.NoContext, code)
	if err != nil {
		return "", fmt.Errorf("code exchange failed: %s", err.Error())
	}
	req, err := http.NewRequest("POST", "https://auth.delta.nitt.edu/api/resources/user", nil)
	if err != nil {
		return "", fmt.Errorf("Unexpected error")
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("failed reading response body: %s", err.Error())
	}
	var user DAuthUserResponse
	err = json.Unmarshal(contents, &user)
	if err != nil {
		return "", err
	}
	err = repository.UpsertUser(user.Name, user.Email)
	if err != nil {
		return "", err
	}
	return user.Email, nil
}
