package soundcloud

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type authResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    uint64 `json:"expires_in"`
	Scope        string `json:"scope"`
}

func authorize(clientId, clientSecret string) (*authResponse, error) {
	resp, err := http.DefaultClient.PostForm(
		fmt.Sprintf("%s/oauth2/token", apiBaseUrl),
		url.Values{
			"client_id":     {clientId},
			"client_secret": {clientSecret},
			"grant_type":    {"client_credentials"},
		},
	)
	if err != nil {
		log.Printf("HTTP request error: %s", err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	response := new(authResponse)
	if err = json.Unmarshal(body, response); err != nil {
		log.Printf("unmarshal error: %s", err.Error())
		return response, err
	} else {
		return response, nil
	}
}
