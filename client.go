package zkc

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

type ClientConfig struct {
	Url     string
	Timeout time.Duration
}

type Client struct {
	config     *ClientConfig
	httpClient *http.Client
}

func NewClient(config *ClientConfig) *Client {
	return &Client{
		config:     config,
		httpClient: &http.Client{Timeout: config.Timeout},
	}
}

func (c *Client) CDKErigon() CdkErigonApi {
	return &cdkErigonApiImpl{
		client: c,
	}
}

func (c *Client) handleRequest(r *request) (*response, error) {
	payload, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, c.config.Url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Add("content-type", "application/json")
	req.Header.Add("accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var res response
	if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}

	return &res, nil
}
