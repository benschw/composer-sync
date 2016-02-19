package loader

import (
	"fmt"
	"strings"

	"github.com/benschw/composer-sync/git"
	"github.com/benschw/composer-sync/packagist"
	"github.com/benschw/composer-sync/stash"
	satis "github.com/benschw/satis-go/satis/client"
	satisapi "github.com/benschw/satis-go/satis/satisphp/api"
)

type Loader struct {
	Stash     *stash.Client
	Packagist *packagist.Client
	Satis     *satis.SatisClient
	DestTpl   string
	StashProj string
}

func (l *Loader) Load(name string, update bool, dryrun bool, recursive bool) error {
	if recursive {
		if err := l.loadRecursive(name, update, dryrun); err != nil {
			return err
		}
	} else {
		if err := l.loadOne(name, update, dryrun); err != nil {
			return err
		}
	}
	if !dryrun {
		if err := l.Satis.GenerateStaticWeb(); err != nil {
			return err
		}
	}
	return nil
}

func (l *Loader) loadOne(name string, update bool, dryrun bool) error {
	p, err := l.Packagist.Get(name)
	if err != nil {
		return err
	}
	return l.loadPackage(p, update, dryrun)
}
func (l *Loader) loadRecursive(name string, update bool, dryrun bool) error {
	packages, err := l.Packagist.GetRecursive(name)
	if err != nil {
		return err
	}

	for _, p := range packages {
		if err := l.loadPackage(p, update, dryrun); err != nil {
			return err
		}
	}

	return nil
}
func (l *Loader) loadPackage(p packagist.Package, update bool, dryrun bool) error {
	repoName := strings.Replace(p.Name, "/", "_", -1)
	destRepo := fmt.Sprintf(l.DestTpl, repoName)

	repo, err := l.Stash.GetRepo(l.StashProj, repoName)
	if err != nil {
		return err
	}

	if repo == nil {
		fmt.Printf("add:  %s \n", repoName)
		if !dryrun {
			if _, err = l.Stash.CreateRepo(l.StashProj, repoName); err != nil {
				return err
			}
			git.MirrorRepo(p.Repository, destRepo)
			if _, err = l.Satis.AddRepo(satisapi.NewRepo("vcs", destRepo)); err != nil {
				return err
			}
		}
	} else if update {

		fmt.Printf("sync: %s \n", repoName)

		if !dryrun {
			git.MirrorRepo(p.Repository, destRepo)
		}
	} else {
		fmt.Printf("skip: %s\n", repoName)
	}

	return nil
}
