package exporter_24x7

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type authenticator struct {
	ClientID     string
	ClientSecret string
	RefreshToken string
}

type tscan struct {
	Token string `json:"access_token"`
}

var Auth authenticator

func getToken() (string, error) {
	ret := initAuth()
	if ret {
		return "", errors.New("access keys are not well formated")
	}
	token, err := getTokenID()
	if err != nil {
		return "", err
	}
	return token, nil
}

func getTokenID() (string, error) {

	var tscan tscan

	apiURL := fmt.Sprintf("https://accounts.zoho.com/oauth/v2/token?client_id=%s&client_secret=%s&refresh_token=%s&grant_type=refresh_token", Auth.ClientID, Auth.ClientSecret, Auth.RefreshToken)
	resp, err := http.Post(apiURL, "", nil)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New(http.StatusText(resp.StatusCode))
	}

	buf, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(buf, &tscan)
	if len(tscan.Token) <= 0 {
		return "", errors.New("token_id wasn't generated")
	}
	return tscan.Token, nil
}

func initAuth() bool {
	Auth.ClientID = os.Getenv("CLIENT_ID_24x7")
	Auth.ClientSecret = os.Getenv("CLIENT_SECRET_24x7")
	Auth.RefreshToken = os.Getenv("REFRESH_TOKEN_24x7")
	if len(Auth.ClientID) < 0 || len(Auth.ClientSecret) < 0 || len(Auth.RefreshToken) < 0 {
		return true
	}
	return false
}
