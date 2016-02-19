package git

import (
	"io/ioutil"
	"os"
	"os/exec"
)

//git clone --bare $ADDR ./tmp/
//cd ./tmp/
//git push --mirror http://foo:asdf@localhost:7990/scm/phpv/${NAME}.git

func MirrorRepo(srcRepo string, destRepo string) error {
	path, err := ioutil.TempDir("/tmp", "vendor-sync")
	if err != nil {
		return err
	}
	defer os.RemoveAll(path)

	if err = cloneBare(srcRepo, path); err != nil {
		return err
	}

	return pushMirror(destRepo, path)
}

func cloneBare(url string, path string) error {
	_, err := exec.
		Command("git", "clone", "--bare", url, path).
		Output()

	return err
}

func pushMirror(url string, path string) error {

	cmd := "cd " + path + " && git push --mirror " + url

	_, err := exec.
		Command("bash", "-c", cmd).
		Output()

	return err
}
