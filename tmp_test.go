package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTmpDirDelete(t *testing.T) {

	// given
	path, err := ioutil.TempDir("/tmp", "vendor-sync")
	assert.Nil(t, err)

	_, err = os.Stat(path)
	assert.Nil(t, err)

	// when
	os.RemoveAll(path)

	//then
	_, err = os.Stat(path)

	assert.NotNil(t, err)
}
