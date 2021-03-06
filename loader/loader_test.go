package loader

import (
	"testing"

	"github.com/benschw/composer-sync/config"
	"github.com/benschw/composer-sync/packagist"
	"github.com/benschw/composer-sync/stash"
	satis "github.com/benschw/satis-go/satis/client"
	"github.com/stretchr/testify/assert"
)

var cfg *config.Config
var stashClient *stash.Client
var satisClient *satis.SatisClient

func init() {
	c, err := config.Load("../test.yaml")
	if err != nil {
		panic(err)
	}
	cfg = c
	stashClient = stash.New(cfg.Stash.ApiUrl, cfg.Stash.Login, cfg.Stash.Password)
	satisClient = &satis.SatisClient{Host: cfg.Satis.ApiUrl}
}

func cleanup() {
	stashRepos, err := stashClient.GetAllReposPage(cfg.Stash.ProjKey)
	if err != nil {
		panic(err)
	}
	for _, repo := range stashRepos {
		stashClient.DeleteRepo(cfg.Stash.ProjKey, repo.Slug)
	}

	satisRepos, err := satisClient.FindAllRepos()
	if err != nil {
		panic(err)
	}
	for _, repo := range satisRepos {
		satisClient.DeleteRepo(repo.Id)
	}
}

func TestLoadOne(t *testing.T) {
	// given
	name := "fliglio/web"
	update := true
	dryrun := false
	recursive := false

	loader := &Loader{
		Stash:     stashClient,
		Packagist: packagist.New(),
		Satis:     satisClient,
		DestTpl:   cfg.Stash.RepoTpl,
		StashProj: cfg.Stash.ProjKey,
	}

	// when
	err := loader.Load(name, update, dryrun, recursive)

	//then
	assert.Nil(t, err)

	stashRepos, _ := stashClient.GetAllReposPage(cfg.Stash.ProjKey)
	assert.Equal(t, 1, len(stashRepos), "should be 1 repo")
	if len(stashRepos) > 0 {
		assert.Equal(t, "fliglio_web", stashRepos[0].Slug, "repo name unexpected")
	}
	satisRepos, _ := satisClient.FindAllRepos()
	assert.Equal(t, 1, len(satisRepos), "should be 1 repo")

	cleanup()
}
func TestLoadRecursive(t *testing.T) {
	// given
	name := "fliglio/web"
	update := true
	dryrun := false
	recursive := true

	loader := &Loader{
		Stash:     stashClient,
		Packagist: packagist.New(),
		Satis:     satisClient,
		DestTpl:   cfg.Stash.RepoTpl,
		StashProj: cfg.Stash.ProjKey,
	}

	// when
	err := loader.Load(name, update, dryrun, recursive)

	//then
	assert.Nil(t, err)

	stashRepos, _ := stashClient.GetAllReposPage(cfg.Stash.ProjKey)
	assert.Equal(t, 8, len(stashRepos), "should be 1 repo")

	satisRepos, _ := satisClient.FindAllRepos()
	assert.Equal(t, 8, len(satisRepos), "should be 1 repo")

	cleanup()
}
