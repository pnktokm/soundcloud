package soundcloud

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Api struct {
	ClientId     string
	ClientSecret string
	AccessToken  string
	RefreshToken string
}

func NewClient(clientId, clientSecret string) (*Api, error) {
	credentials, err := authorize(clientId, clientSecret)
	if err != nil {
		log.Printf("authenticate error: %s", err.Error())
		return nil, err
	}

	return &Api{
		clientId,
		clientSecret,
		credentials.AccessToken,
		credentials.RefreshToken,
	}, err
}

func (api Api) GetTracks(query string) ([]*Track, error) {
	url := fmt.Sprintf("%s/tracks?client_id=%s&q=%s", apiBaseUrl, api.ClientId, query)
	resp, err := http.DefaultClient.Get(url)
	defer resp.Body.Close()

	if err != nil {
		log.Print(err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)
	response := make([]*Track, 0)
	if err = json.Unmarshal(body, &response); err != nil {
		log.Println("unmarshal error: " + err.Error())
	}

	return response, err
}

func (api Api) GetFavorites(userId int) ([]*Track, error) {
	url := fmt.Sprintf("%s/users/%d/favorites?client_id=%s", apiBaseUrl, userId, api.ClientId)
	resp, err := http.DefaultClient.Get(url)
	defer resp.Body.Close()

	if err != nil {
		log.Print(err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)
	response := make([]*Track, 0)
	if err = json.Unmarshal(body, &response); err != nil {
		log.Println("unmarshal error: " + err.Error())
	}

	return response, err
}
