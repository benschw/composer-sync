package stash

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRepo(t *testing.T) {
	// given
	c := New("http://foo:asdf@localhost:7990", "foo", "asdf")

	// when
	r, err := c.GetRepo("UT", "test")
	assert.Nil(t, err)
	assert.Nil(t, r)

	r2, err := c.CreateRepo("UT", "test")
	assert.Nil(t, err)
	assert.NotNil(t, r2)

	r3, err := c.GetRepo("UT", "test")
	assert.Nil(t, err)
	assert.NotNil(t, r3)

	//then
	assert.Equal(t, "test", r2.Slug, "name doesn't match")
	assert.Equal(t, "test", r3.Slug, "name doesn't match")

	c.DeleteRepo("UT", "test")
}

func TestCreateRepo(t *testing.T) {
	// given
	c := New("http://localhost:7990", "foo", "asdf")

	// when
	r, err := c.CreateRepo("UT", "web3")

	found, err2 := c.GetRepo("UT", "web3")

	//then
	assert.Nil(t, err)
	assert.Nil(t, err2)
	assert.Equal(t, r, found, "name doesn't match")
	assert.Equal(t, "web3", r.Slug, "slug doesn't match expected")

	c.DeleteRepo("UT", "web3")
}

func TestGetPage(t *testing.T) {
	// given
	c := New("http://localhost:7990", "foo", "asdf")
	c.CreateRepo("UT", "web1")
	c.CreateRepo("UT", "web2")
	c.CreateRepo("UT", "web3")

	// when
	repos, err := c.GetAllReposPage("UT")
	log.Printf("%+v", repos)
	//then
	assert.Nil(t, err)

	assert.Equal(t, 3, len(repos), "should be 3 repos")

	c.DeleteRepo("UT", "web1")
	c.DeleteRepo("UT", "web2")
	c.DeleteRepo("UT", "web3")

}
