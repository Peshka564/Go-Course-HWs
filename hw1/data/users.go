package data

import (
	"github.com/Peshka564/Go-Course-HWs/hw1/httpclient"
)

type UserData struct {
	Login       string `json:"login"`
	PublicRepos int    `json:"public_repos"`
	Followers   int    `json:"followers"`
}

func GetUser(username string, httpClient *httpclient.HttpClient) (*UserData, error) {
	var user UserData
	userRoute := "/users/" + username
	err := httpClient.Get(userRoute, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
