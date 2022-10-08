package config_test

import (
	"testing"

	"github.com/devdinu/simple-api/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConfigLoadSuccess(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			require.Fail(t, "failed to load config")
		}
	}()

	app := config.MustLoad()

	assert.Equal(t, config.AppAddress(), "localhost:8080")
	assert.Equal(t, app.DB.Host, "localhost")
	assert.Equal(t, app.DB.Port, 5432)
}
