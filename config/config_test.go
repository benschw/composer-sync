package config

import (
	"testing"

	"github.com/benschw/opin-go/config"
	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	// given
	cfgPath := "../test.yaml"

	// when
	cfg := &Config{}
	err := config.Bind(cfgPath, cfg)

	//then
	assert.Nil(t, err)
	assert.Equal(t, "UT", cfg.Stash.ProjKey, "project key looks wrong")
}
