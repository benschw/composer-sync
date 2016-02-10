package git

import "os/exec"

//git clone --bare $ADDR ./tmp/
//cd ./tmp/
//git push --mirror http://foo:asdf@localhost:7990/scm/phpv/${NAME}.git

func CloneBare(url string, path string) error {
	_, err := exec.
		Command("git", "clone", "--bare", url, path).
		Output()

	return err
}

func PushMirror(url string, path string) error {

	cmd := "cd " + path + " && git push --mirror " + url

	_, err := exec.
		Command("bash", "-c", cmd).
		Output()

	return err
}
