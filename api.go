package gopubg

import (
	"fmt"
	"net/url"

	"github.com/driquet/gopubg/models/match"
	"github.com/driquet/gopubg/models/player"
)

//API struct for holding the key
type API struct {
	Key string
}

//NewAPI creating a new API from a key
func NewAPI(key string) *API {
	return &API{
		Key: key,
	}
}

//RequestStatus A function that prints out the current status of the API Key
func (a *API) RequestStatus() error {
	endpointURL := "https://api.playbattlegrounds.com/status"

	buffer, err := httpRequest(endpointURL, a.Key)
	if err != nil {
		return err
	}

	fmt.Printf("data:\n%s\n", buffer)

	return nil
}

//RequestSinglePlayerByName A function that takes a shard string, and a player name, and returns either that players data, or an error
func (a *API) RequestSinglePlayerByName(shard, playerName string) (*player.Player, error) {
	parameters := url.Values{
		"filter[playerNames]": {playerName},
	}

	endpointURL := fmt.Sprintf("https://api.playbattlegrounds.com/shards/%s/players?%s", shard, parameters.Encode())

	buffer, err := httpRequest(endpointURL, a.Key)
	if err != nil {
		return nil, err
	}

	fmt.Printf("data:\n%s\n", buffer)

	// TODO parse player
	return nil, nil
}

//RequestMatch given a shard and a match_id string will print either match info, or a error
func (a *API) RequestMatch(shard, matchID string) (*match.Match, error) {

	endpointURL := fmt.Sprintf("https://api.playbattlegrounds.com/shards/%s/matches/%s", shard, matchID)
	buffer, err := httpRequest(endpointURL, a.Key)
	if err != nil {
		return nil, err
	}

	fmt.Printf("data:\n%s\n", buffer)
	return nil, nil
}
