package service

import (
	"io"
	"log"
	"net/http"
	"time"
)

type HttpClient struct {
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
	return nil
}

func (c *HttpClient) bindHeaders(req *http.Request) {
	for key, value := range c.headers {
		req.Header.Add(key, value)
	}
}

func (c *HttpClient) Get(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	c.bindHeaders(req)

	resp, err := c.worker.Do(req)
	if err != nil {
		return nil, err
	}

	log.Println(url)
	log.Printf("headers {%s}", resp.Request.Header)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
