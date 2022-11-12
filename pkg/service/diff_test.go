package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGoJsonDiffer_Diff(t *testing.T) {
	differ := NewGoJsonDiffer()

	for _, tc := range []struct {
		left  []byte
		right []byte
	}{
		{left: []byte(`{"a": "b"}`), right: []byte(`{"a": "c"}`)},
		{left: []byte(`[{"a": "ba"}]`), right: []byte(`[{"a": "c"}]`)},
	} {
		diff, err := differ.Diff(tc.left, tc.right)
		assert.Nil(t, err)
		assert.NotNil(t, diff)
	}
}
