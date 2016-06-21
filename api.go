package soundcloud

import (
	"encoding/json"
	"fmt"
	"github.com/yanatan16/golang-soundcloud/soundcloud"
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

func (api Api) GetFavorites(userId int) []*soundcloud.Track {
	url := fmt.Sprintf("%s/users/%d/favorites?client_id=%s", apiBaseUrl, userId, api.ClientId)
	resp, err := http.DefaultClient.Get(url)
	defer resp.Body.Close()

	if err != nil {
		log.Print(err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)
	response := make([]*soundcloud.Track, 0)
	if err = json.Unmarshal(body, &response); err != nil {
		log.Println("unmarshal error: " + err.Error())
	}

	return response
}
