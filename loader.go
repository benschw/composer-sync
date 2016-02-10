package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/benschw/composer-sync/git"
	"github.com/benschw/composer-sync/packagist"
	"github.com/benschw/composer-sync/stash"
)

type Loader struct {
	SClient   *stash.Client
	PClient   *packagist.Client
	DestTpl   string
	StashProj string
}

func (l *Loader) Load(name string, update bool, dryrun bool) error {
	packages, err := l.PClient.GetRecursive(name)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, p := range packages {
		repoName := strings.Replace(p.Name, "/", "_", -1)
		destRepo := fmt.Sprintf(l.DestTpl, repoName)

		repo, err := l.SClient.GetRepo(l.StashProj, repoName)
		if err != nil {
			return err
		}

		if repo == nil {
			fmt.Printf("add:  %s \n", repoName)
			if !dryrun {
				if _, err = l.SClient.CreateRepo(l.StashProj, repoName); err != nil {
					return err
				}
				l.syncRepo(repoName, p.Repository, destRepo)
			}
		} else if update {

			fmt.Printf("sync: %s \n", repoName)

			if !dryrun {
				l.syncRepo(repoName, p.Repository, destRepo)
			}
		} else {
			fmt.Printf("skip: %s\n", repoName)
		}
	}

	return nil
}

func (l *Loader) syncRepo(repoName string, srcRepo string, destRepo string) error {
	path, err := ioutil.TempDir("/tmp", "vendor-sync")
	if err != nil {
		return err
	}
	defer os.Remove(path)

	if err = git.CloneBare(srcRepo, path); err != nil {
		return err
	}

	return git.PushMirror(destRepo, path)
}
