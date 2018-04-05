package player

import (
	"errors"
	"io"
	"reflect"
	"time"

	"github.com/slemgrim/jsonapi"
)

// Player structure represents a player entry
type Player struct {
	ID           string    `json:"primary,player"`
	Name         string    `jsonapi:"attr,name"`
	ShardID      string    `jsonapi:"attr,shardId"`
	CreatedAt    time.Time `jsonapi:"attr,createdAt,iso8601"`
	UpdatedAt    time.Time `jsonapi:"attr,updatedAt,iso8601"`
	PatchVersion string    `jsonapi:"attr,patchVersion"`
	TitleID      string    `jsonapi:"attr,titleId"`
	Matches      []*Match  `jsonapi:"relation,matches"`
}

// Match structure represent data related to a PUBG match
type Match struct {
	ID     string `jsonapi:"primary,match"`
	GameID string `jsonapi:"attr,id"`
}

// ParsePlayers parses a json response containing players information
func ParsePlayers(in io.Reader) ([]*Player, error) {
	result, err := jsonapi.UnmarshalManyPayload(in, reflect.TypeOf(new(Player)))
	if err != nil {
		return nil, err
	}

	players := make([]*Player, len(result))
	for idx, elt := range result {
		player, ok := elt.(*Player)
		if !ok {
			return nil, errors.New("Failed to convert players")
		}
		players[idx] = player
	}
	return players, nil
}
