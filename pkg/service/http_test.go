package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHttpServiceImpl_MultiGet(t *testing.T) {
	httpService := NewHttpService()

	res, err := httpService.MultiGet([]string{"https://636f1c4ebb9cf402c80fface.mockapi.io/kuckjwi", "https://636f1c4ebb9cf402c80fface.mockapi.io/kuckjwi2"}, nil)
	bytes := res.([][]byte)

	assert.Nil(t, err)
	assert.NotNil(t, bytes[0])
	assert.NotNil(t, bytes[1])
}
