package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHttpClientOptions(t *testing.T) {
	_, err := NewHttpClient(
		WithHeaders(map[string]string{
			"Content-Type": "application/json",
		}),
	)
	assert.Nil(t, err)
}
