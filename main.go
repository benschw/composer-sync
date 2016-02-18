package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/benschw/composer-sync/packagist"
	"github.com/benschw/composer-sync/stash"
	"github.com/benschw/opin-go/config"
	satis "github.com/benschw/satis-go/satis/client"
)

type Config struct {
	Stash StashConfig `yaml:"stash"`
	Satis SatisConfig `yaml:"satis"`
}
type StashConfig struct {
	RepoTpl string `yaml:"repo_tpl"`
	ApiUrl  string `yaml:"api_url"`
	ProjKey string `yaml:"proj_key"`
}
type SatisConfig struct {
	ApiUrl string `yaml:"api_url"`
}

func main() {
	update := flag.Bool("u", false, "update existing repositories")
	recursive := flag.Bool("r", false, "Recusively operate on all dependencies")
	dryrun := flag.Bool("dryrun", false, "perform dryrun")
	cfgPath := flag.String("config", "~/.composer-sync.yaml", "config path")

	flag.Parse()

	cfg := &Config{}
	if err := config.Bind(*cfgPath, cfg); err != nil {
		fmt.Printf("Bad Config: %s", err)
		os.Exit(1)
	}

	// Pull subcommand & package name from args
	if flag.NArg() != 1 {
		fmt.Printf("expected 1 arguments\n")
		flag.Usage()
		os.Exit(1)
	}

	name := flag.Arg(0)

	loader := &Loader{
		Stash:     stash.New(cfg.Stash.ApiUrl),
		Packagist: packagist.New(),
		Satis:     &satis.SatisClient{Host: cfg.Satis.ApiUrl},
		DestTpl:   cfg.Stash.RepoTpl,
		StashProj: cfg.Stash.ProjKey,
	}

	if err := loader.Load(name, *update, *dryrun, *recursive); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
