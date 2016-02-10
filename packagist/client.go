package packagist

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/benschw/opin-go/rest"
)

func New() *Client {
	return &Client{}
}

type Client struct {
}

func (c *Client) Get(name string) (Package, error) {
	var p PackageWrapper
	url := fmt.Sprintf("https://packagist.org/packages/%s.json", name)

	r, err := rest.MakeRequest("GET", url, nil)
	if err != nil {
		return p.Package, err
	}
	err = rest.ProcessResponseEntity(r, &p, http.StatusOK)
	return p.Package, err
}

func (c *Client) GetRecursive(name string) ([]Package, error) {
	var packages []Package

	p, err := c.Get(name)
	if err != nil {
		return packages, err
	}

	packages = append(packages, p)

	if _, ok := p.Versions["dev-master"]; !ok {
		return packages, fmt.Errorf("version dev-master not found for package %s", name)
	}

	for name, _ := range p.Versions["dev-master"].Require {
		if strings.Index(name, "/") >= 0 {
			deps, err := c.GetRecursive(name)
			if err != nil {
				return packages, err
			}
			packages = append(packages, deps...)
		}
	}
	return packages, nil
}
