package service

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"time"
)

type HttpClient struct {
	url     string
	headers map[string]string
	body    map[string]string
	worker  *http.Client
}

func NewHttpClient(options ...func(client *HttpClient)) (*HttpClient, error) {
	client := &HttpClient{
		worker: &http.Client{
			Transport: &http.Transport{
				MaxIdleConns:    20,
				IdleConnTimeout: 30 * time.Second,
			},
		},
	}
	for _, option := range options {
		option(client)
	}
	return client, client.validate()
}

func WithURL(url string) func(client *HttpClient) {
	return func(client *HttpClient) {
		client.url = url
	}
}

func WithHeaders(headers map[string]string) func(client *HttpClient) {
	return func(client *HttpClient) {
		client.headers = headers
	}
}

func WithBody(body map[string]string) func(client *HttpClient) {
	return func(client *HttpClient) {
		client.body = body
	}
}

func (c *HttpClient) validate() error {
	if c.url == "" {
		return errors.New("url is required")
	}
	if !strings.HasPrefix(c.url, "http") {
		return errors.New("url must be start with http")
	}
	return nil
}

func (c *HttpClient) toQueryString() string {
	queries := make([]string, len(c.body))
	for key, value := range c.body {
		queries = append(queries, key+"="+value)
	}
	return "?" + strings.Join(queries, "&")
}

func (c *HttpClient) bindHeaders(req *http.Request) {
	for key, value := range c.headers {
		req.Header.Add(key, value)
	}
}

func (c *HttpClient) Get() ([]byte, error) {
	req, err := http.NewRequest("GET", c.url+c.toQueryString(), nil)
	if err != nil {
		return nil, err
	}

	c.bindHeaders(req)

	resp, err := c.worker.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
