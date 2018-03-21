package gopubg

import (
	"errors"
	"io"
	"reflect"
	"time"

	"github.com/slemgrim/jsonapi"
)

// Match structure represent data related to a PUBG match
type Match struct {
	ID           string    `jsonapi:"primary,match"`
	CreatedAt    time.Time `jsonapi:"attr,createdAt,iso8601"`
	Duration     int       `jsonapi:"attr,duration"`
	GameMode     string    `jsonapi:"attr,gameMode"`
	PatchVersion string    `jsonapi:"attr,patchVersion"`
	ShardID      string    `jsonapi:"attr,shardId"`
	TitleID      string    `jsonapi:"attr,titleId"`
	Rosters      []*Roster `jsonapi:"relation,rosters"`
	// Todo stats, tags, assets, rounds, spectators
}

// ParseMatch parses a json response containing matches information
func ParseMatch(in io.Reader) ([]*Match, error) {
	result, err := jsonapi.UnmarshalManyPayload(in, reflect.TypeOf(new(Match)))
	if err != nil {
		return nil, err
	}

	matches := make([]*Match, len(result))
	for idx, elt := range result {
		match, ok := elt.(*Match)
		if !ok {
			return nil, errors.New("Failed to convert matches")
		}
		matches[idx] = match
	}
	return matches, nil
}
