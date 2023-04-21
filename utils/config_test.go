package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadBaseConfig(t *testing.T) {
	assert.NotPanics(t, func() {
		config := LoadBaseConfig("../", "test")
		assert.NotNil(t, config)
	})
}
