package oauth_utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"delivery_app_api.mmedic.com/m/v2/src/dto"
	"delivery_app_api.mmedic.com/m/v2/src/models"
	"delivery_app_api.mmedic.com/m/v2/src/services/request_service"
	"delivery_app_api.mmedic.com/m/v2/src/utils/env_utils"
)

type GithubAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

type GithubCustomerDetails struct {
	Name     string `json:"name"`
	Username string `json:"login"`
	Email    string `json:"email"`
	City     string `json:"location"`
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
	bodyValues := map[string]string{
		"client_id":     env_utils.GetEnvVar(fmt.Sprintf("OAUTH_GITHUB_%s_CLIENT_ID", action)),
		"client_secret": env_utils.GetEnvVar(fmt.Sprintf("OAUTH_GITHUB_%s_CLIENT_SECRET", action)),
		"code":          code}
	json_data, err := json.Marshal(bodyValues)

	if err != nil {
		return "", err
	}

	request := request_service.CreatePOSTRequest("https://github.com/login/oauth/access_token", bytes.NewBuffer(json_data))

	respbody, status, err := request.Send()
	if err != nil {
		return "", err
	}

	if status >= 400 && status <= 600 {
		return "", models.CreateCustomError(fmt.Sprintf("There was an error while sending the Request. Error status: %d", status))
	}

	var ghresp GithubAccessTokenResponse
	json.Unmarshal(respbody, &ghresp)

	return ghresp.AccessToken, nil
}

func GetCustomerOAuthProfile(accessToken string) (*dto.CustomerInputDto, error) {
	request := request_service.CreateGETRequest("https://api.github.com/user")
	request.SetHeader("Authorization", fmt.Sprintf("token %s", accessToken))

	respbody, status, err := request.Send()
	if err != nil {
		return nil, err
	}

	if status >= 400 && status <= 600 {
		return nil, models.CreateCustomError(fmt.Sprintf("There was an error while sending the Request. Error status: %d", status))
	}

	var gCustDetails GithubCustomerDetails
	json.Unmarshal(respbody, &gCustDetails)

	user := new(dto.CustomerInputDto)
	addr := new(dto.AddressInputDto)

	user.Address = addr
	nameSurname := strings.Split(gCustDetails.Name, " ")
	user.Name = nameSurname[0]
	user.Surname = nameSurname[1]
	user.Username = gCustDetails.Username
	user.Email = gCustDetails.Email
	user.Address.City = gCustDetails.City

	return user, nil
}
