package oauth_utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"delivery_app_api.mmedic.com/m/v2/src/dto"
	"delivery_app_api.mmedic.com/m/v2/src/utils/env_utils"
)

type GithubAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

func GetLoginOAuthURL() string {
	clientId := env_utils.GetEnvVar("OAUTH_GITHUB_LOGIN_CLIENT_ID")
	clientSecret := env_utils.GetEnvVar("OAUTH_GITHUB_LOGIN_REDIRECTION_URL")

	urlTemplate := env_utils.GetEnvVar("OAUTH_GITHUB_URL_TEMPLATE")

	return fmt.Sprintf(urlTemplate, clientId, clientSecret)
}

func GetRegistrationOAuthURL() string {
	clientId := env_utils.GetEnvVar("OAUTH_GITHUB_REGISTRATION_CLIENT_ID")
	clientSecret := env_utils.GetEnvVar("OAUTH_GITHUB_REGISTRATION_REDIRECTION_URL")

	urlTemplate := env_utils.GetEnvVar("OAUTH_GITHUB_URL_TEMPLATE")

	return fmt.Sprintf(urlTemplate, clientId, clientSecret)
}

func GetCustomerGithubInformation(code, action string) (*dto.CustomerInputDto, error) {
	accessToken, err := GetOAuthAccessToken(code, action)
	if err != nil {
		return nil, err
	}

	customerData, err := GetCustomerOAuthProfile(accessToken)
	if err != nil {
		return nil, err
	}

	return customerData, nil
}

func GetOAuthAccessToken(code, action string) (string, error) {
	// TODO: REFACTOR AS REQUEST SERVICE
	bodyValues := map[string]string{
		"client_id":     env_utils.GetEnvVar(fmt.Sprintf("OAUTH_GITHUB_%s_CLIENT_ID", action)),
		"client_secret": env_utils.GetEnvVar(fmt.Sprintf("OAUTH_GITHUB_%s_CLIENT_SECRET", action)),
		"code":          code}
	json_data, err := json.Marshal(bodyValues)

	if err != nil {
		return "", err
	}

	client := &http.Client{}
	req, _ := http.NewRequest("POST", "https://github.com/login/oauth/access_token", bytes.NewBuffer(json_data))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	respbody, _ := ioutil.ReadAll(resp.Body)

	var ghresp GithubAccessTokenResponse
	json.Unmarshal(respbody, &ghresp)

	return ghresp.AccessToken, nil
}

func GetCustomerOAuthProfile(accessToken string) (*dto.CustomerInputDto, error) {
	// TODO: REFACTOR AS REQUEST SERVICE
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://api.github.com/user", nil)
	req.Header.Set("Authorization", fmt.Sprintf("token %s", accessToken))

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var res map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	user := new(dto.CustomerInputDto)
	addr := new(dto.AddressInputDto)

	user.Address = addr
	nameSurname := strings.Split(res["name"].(string), " ")
	user.Name = nameSurname[0]
	user.Surname = nameSurname[1]
	user.Username = res["login"].(string)
	user.Email = res["email"].(string)
	user.Address.City = res["location"].(string)

	return user, nil
}
