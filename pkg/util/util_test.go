package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandString(t *testing.T) {
	s1 := GenerateRandString(10)
	s2 := GenerateRandString(10)
	s3 := GenerateRandString(10)
	s4 := GenerateRandString(10)
	assert.Equal(t, 10, len(s1))
	assert.Equal(t, 10, len(s2))
	assert.Equal(t, 10, len(s3))
	assert.Equal(t, 10, len(s4))
	assert.NotEqual(t, s1, s2)
	assert.NotEqual(t, s3, s4)
}

func TestGenerateSubscribeEndpoint(t *testing.T) {
	s := GenerateSubscribeEndpoint()
	assert.NotEqual(t, make([]string, 16), s)
}
