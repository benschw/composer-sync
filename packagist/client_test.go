package packagist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPackage(t *testing.T) {
	// given
	c := New()

	// when
	p, err := c.Get("fliglio/web")

	//then
	assert.Nil(t, err)
	assert.Equal(t, "fliglio/web", p.Name, "name doesn't match")
}

func TestGetPackageRecursive(t *testing.T) {
	// given
	c := New()

	// when
	ps, err := c.GetRecursive("fliglio/web")

	//fmt.Printf("%+v", ps)

	//then
	assert.Nil(t, err)

	assert.True(t, len(ps) > 1, "fliglio/web should have some dependencies")
}
