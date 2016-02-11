package stash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRepo(t *testing.T) {
	// given
	c := New("http://foo:asdf@localhost:7990")

	// when
	r, err := c.GetRepo("PHPV", "test")
	assert.Nil(t, err)
	assert.Nil(t, r)

	r2, err := c.CreateRepo("PHPV", "test")
	assert.Nil(t, err)
	assert.NotNil(t, r2)

	r3, err := c.GetRepo("PHPV", "test")
	assert.Nil(t, err)
	assert.NotNil(t, r3)

	//then
	assert.Equal(t, "test", r2.Slug, "name doesn't match")
	assert.Equal(t, "test", r3.Slug, "name doesn't match")

	c.DeleteRepo("PHPV", "test")
}

func TestCreateRepo(t *testing.T) {
	// given
	c := New("http://foo:asdf@localhost:7990")

	// when
	r, err := c.CreateRepo("PHPV", "web3")

	found, err2 := c.GetRepo("PHPV", "web3")

	//then
	assert.Nil(t, err)
	assert.Nil(t, err2)
	assert.Equal(t, r, found, "name doesn't match")
	assert.Equal(t, "web3", r.Slug, "slug doesn't match expected")

	c.DeleteRepo("PHPV", "web3")
}
