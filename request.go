package gopubg

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

func httpRequest(url, key string) (*bytes.Buffer, error) {
	logrus.WithField("url", url).Info("pubg api request")

	// Create request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Set request options
	req.Header.Set("Authorization", key)
	req.Header.Set("Accept", "application/vnd.api+json")
	req.Header.Set("Accept-Encoding", "gzip")

	// Execute request
	client := &http.Client{}
	response, err := client.Do(req)

	// Check http response code
	if response.StatusCode != 200 {
		response.Body.Close()
		return nil, fmt.Errorf("HTTP request failed: %s", response.Status)
	}

	// https://api.playbattlegrounds.com/shards/pc-eu?filter[playerNames]=dreuhdreuh
	// https://api.playbattlegrounds.com/shards/pc-eu/players?filter[playerNames]=dreuhdreuh

	// Retrieve response body
	var reader io.ReadCloser
	switch response.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(response.Body)
		if err != nil {
			return nil, err
		}
	default:
		reader = response.Body
	}
	defer reader.Close()

	var buffer bytes.Buffer
	buffer.ReadFrom(reader)

	return &buffer, nil
}
