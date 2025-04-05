package zkc

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

type clientConfig struct {
	Url        string
	Timeout    time.Duration
	MaxRetries int
	RetryDelay time.Duration
}

type ClientOption interface {
	applyOption(*clientConfig)
}

type optionFunc func(*clientConfig)

func (fn optionFunc) applyOption(opt *clientConfig) {
	fn(opt)
}

func WithTimeout(timeout time.Duration) ClientOption {
	return optionFunc(func(opt *clientConfig) {
		opt.Timeout = timeout
	})
}

func WithMaxRetries(maxRetries int, interval time.Duration) ClientOption {
	return optionFunc(func(opt *clientConfig) {
		opt.MaxRetries = maxRetries
		opt.RetryDelay = interval
	})
}

type Client struct {
	config     *clientConfig
	httpClient *http.Client
}

func newClient(config *clientConfig) *Client {
	return &Client{
		config:     config,
		httpClient: &http.Client{Timeout: config.Timeout},
	}
}

func CDKErigonClient(url string, options ...ClientOption) CdkErigonApi {
	cfg := new(clientConfig)
	for _, opt := range options {
		opt.applyOption(cfg)
	}
	cfg.Url = url
	c := newClient(cfg)
	return &cdkErigonApiImpl{
		client: c,
	}
}

func (c *Client) handleRequest(r *request) (*response, error) {
	var err error

	payload, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	var lastErr error
	attempts := c.config.MaxRetries + 1

	for attempt := 0; attempt < attempts; attempt++ {
		var req *http.Request
		req, err = http.NewRequest(http.MethodPost, c.config.Url, bytes.NewBuffer(payload))
		if err != nil {
			return nil, err
		}

		req.Header.Add("content-type", "application/json")
		req.Header.Add("accept", "application/json")

		var resp *http.Response
		resp, err = c.httpClient.Do(req)
		if err != nil {
			lastErr = err
			if attempt < attempts-1 {
				time.Sleep(c.config.RetryDelay)
				continue
			}
			return nil, err
		}
		defer resp.Body.Close()

		var res response
		if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
			return nil, err
		}

		return &res, nil
	}

	return nil, lastErr
}
