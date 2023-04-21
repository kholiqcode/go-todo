package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunMigration(t *testing.T) {
	config := LoadBaseConfig("../", "test")
	db := ConnectDB(config)

	assert.NotPanics(t, func() {
		config.MIGRATION_URL = "file://../database/migrations"
		RunMigration(db, config)
	})
}
