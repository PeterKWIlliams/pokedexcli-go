package api

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

const (
	BaseURL = "https://pokeapi.co/api/v2"
)

type Client struct {
	httpClient *http.Client
}

func NewClient(timeout time.Duration) *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: timeout,
		},
	}
}

func (client *Client) GetLocations(pageURL *string) (LocationsResponse, error) {
	baseURL := BaseURL + "/location"

	var res *http.Response
	var err error

	if pageURL != nil {
		baseURL = *pageURL
	}
	res, err = client.httpClient.Get(baseURL)
	if err != nil {
		return LocationsResponse{}, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationsResponse{}, err
	}
	locations := LocationsResponse{}
	err = json.Unmarshal(body, &locations)
	if err != nil {
		return LocationsResponse{}, err
	}
	return locations, nil
}
