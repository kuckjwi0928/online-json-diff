package service

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"time"
)

type HttpClient struct {
	headers map[string]string
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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return body, nil
}

func (c *HttpClient) Post(url string, requestBody []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	c.bindHeaders(req)

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	resp, err := c.worker.Do(req)
	if err != nil {
		return nil, err
	}

	log.Println(url)
	log.Printf("headers: %s", req.Header)
	log.Printf("body: %s", string(requestBody))

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return body, nil
}
