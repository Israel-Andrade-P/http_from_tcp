package headers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	// Test: Valid single header
	h := NewHeaders()
	data := []byte("Host: localhost:42069\r\nFooFoo:     barbar      \r\n\r\n")
	n, done, err := h.Parse(data)
	require.NoError(t, err)
	require.NotNil(t, h.headers)
	assert.Equal(t, "localhost:42069", h.Get("Host"))
	assert.Equal(t, "barbar", h.Get("FooFoo"))
	assert.Equal(t, "", h.Get("MissingKey"))
	assert.Equal(t, 51, n)
	assert.True(t, done)

	// Test: Valid single header
	h = NewHeaders()
	data = []byte("Content-Type: application/json\r\nContent-Type: text\r\n\r\n")
	n, done, err = h.Parse(data)
	require.NoError(t, err)
	require.NotNil(t, h.headers)
	assert.Equal(t, "application/json, text", h.Get("CONTENT-TYPE"))

	// Test: Invalid spacing header
	h = NewHeaders()
	data = []byte("       Host : localhost:42069       \r\n\r\n")
	n, done, err = h.Parse(data)
	require.Error(t, err)
	assert.Equal(t, 0, n)
	assert.False(t, done)

	// Test: Invalid token header key
	h = NewHeaders()
	data = []byte("H©st: localhost:42069\r\n\r\n")
	n, done, err = h.Parse(data)
	require.Error(t, err)
	assert.Equal(t, 0, n)
	assert.False(t, done)
}
