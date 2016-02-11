package main

import (
	"testing"

	"github.com/benschw/opin-go/config"
	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	// given
	cfgPath := "./cfg.yaml"

	// when
	cfg := &Config{}
	err := config.Bind(cfgPath, cfg)

	//then
	assert.Nil(t, err)
	assert.Equal(t, "PHPV", cfg.Stash.ProjKey, "project key looks wrong")
}
