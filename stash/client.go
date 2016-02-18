package stash

import (
	"fmt"
	"log"
	"net/http"

	"github.com/benschw/opin-go/rest"
)

type RepoPage struct {
	Size       int    `json:"size"`
	Limit      int    `json:"limit"`
	IsLastPage bool   `json:"isLiastPage"`
	Values     []Repo `json:"values"`
}

type RepoConfig struct {
	Name     string `json:"name"`
	ScmId    string `json:"scmId"`
	Forkable bool   `json:"forkable"`
}

type Repo struct {
	Slug  string `json:"slug"`
	State string `json:"state"`
}

func New(url string) *Client {
	return &Client{Url: url}
}

type Client struct {
	Url string
}

func (c *Client) CreateRepo(proj string, name string) (*Repo, error) {
	var repo Repo

	body := RepoConfig{Name: name, ScmId: "git", Forkable: true}

	url := fmt.Sprintf("%s/rest/api/1.0/projects/%s/repos", c.Url, proj)
	r, err := rest.NewRequestH("POST", url, map[string]interface{}{"X-Atlassian-Token": "no-check"}, body)
	if err != nil {
		return &repo, err
	}
	err = rest.ProcessResponseEntity(r, &repo, http.StatusCreated)
	if err != nil {
		b, _ := rest.ProcessResponseBytes(r, http.StatusBadRequest)
		log.Println(string(b[:]))
	}
	return &repo, err
}
func (c *Client) GetRepo(proj string, name string) (*Repo, error) {
	var repo Repo
	url := fmt.Sprintf("%s/rest/api/1.0/projects/%s/repos/%s", c.Url, proj, name)

	r, err := rest.NewRequestH("GET", url, map[string]interface{}{"X-Atlassian-Token": "no-check"}, nil)
	if err != nil {
		return &repo, err
	}
	err = rest.ProcessResponseEntity(r, &repo, http.StatusOK)
	if err != nil {
		_, err := rest.ProcessResponseBytes(r, http.StatusNotFound)
		return nil, err
	}
	return &repo, err
}
func (c *Client) GetAllReposPage(proj string) ([]Repo, error) {
	var page RepoPage
	url := fmt.Sprintf("%s/rest/api/1.0/projects/%s/repos", c.Url, proj)

	r, err := rest.NewRequestH("GET", url, map[string]interface{}{"X-Atlassian-Token": "no-check"}, nil)
	if err != nil {
		return page.Values, err
	}
	err = rest.ProcessResponseEntity(r, &page, http.StatusOK)
	if err != nil {
		_, err := rest.ProcessResponseBytes(r, http.StatusNotFound)
		return nil, err
	}
	return page.Values, err
}
func (c *Client) DeleteRepo(proj string, name string) error {
	url := fmt.Sprintf("%s/rest/api/1.0/projects/%s/repos/%s", c.Url, proj, name)

	r, err := rest.NewRequestH("DELETE", url, map[string]interface{}{"X-Atlassian-Token": "no-check"}, nil)
	if err != nil {
		return err
	}
	_, err = rest.ProcessResponseBytes(r, http.StatusAccepted)
	return err
}
