package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnectDB(t *testing.T) {
	config := LoadBaseConfig("../", "test")
	assert.NotPanics(t, func() {
		db := ConnectDB(config)
		assert.NotNil(t, db)
	})
}
