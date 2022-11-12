package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHttpClientOptions(t *testing.T) {
	_, err := NewHttpClient(
		WithURL(""),
		WithHeaders(map[string]string{
			"Content-Type": "application/json",
		}),
		WithBody(map[string]string{
			"key": "value",
		}),
	)

	_, err2 := NewHttpClient(
		WithURL("tcp://"),
	)

	_, err3 := NewHttpClient(
		WithURL("http://localhost:8080"),
		WithHeaders(map[string]string{
			"Content-Type": "application/json",
		}),
		WithBody(map[string]string{
			"key": "value",
		}),
	)

	assert.Error(t, err)
	assert.Error(t, err2)
	assert.Nil(t, err3)
}
