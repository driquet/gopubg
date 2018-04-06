package gopubg

import (
	"fmt"
	"net/url"

	"github.com/driquet/gopubg/models/player"
)

type API struct {
	Key string
}

func NewAPI(key string) *API {
	return &API{
		Key: key,
	}
}

func (a *API) RequestStatus() error {
	endpoint_url := "https://api.playbattlegrounds.com/status"

	buffer, err := httpRequest(endpoint_url, a.Key)
	if err != nil {
		return err
	}

	fmt.Printf("data:\n%s\n", buffer)

	return nil
}

func (a *API) RequestSinglePlayerByName(shard, playerName string) (*player.Player, error) {
	parameters := url.Values{
		"filter[playerNames]": {playerName},
	}

	endpoint_url := fmt.Sprintf("https://api.playbattlegrounds.com/shards/%s/players?%s", shard, parameters.Encode())

	buffer, err := httpRequest(endpoint_url, a.Key)
	if err != nil {
		return nil, err
	}

	fmt.Printf("data:\n%s\n", buffer)

	return nil
}
