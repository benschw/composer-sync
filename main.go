package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/benschw/composer-sync/packagist"
	"github.com/benschw/composer-sync/stash"
	"github.com/benschw/opin-go/config"
)

type Config struct {
	StashRepoTpl string `json:"stashrepotpl"`
	StashAPIUrl  string `json:"stashapiurl"`
	StashProj    string `json:"stashproj"`
}

func main() {
	update := flag.Bool("u", false, "update existing repositories")
	dryrun := flag.Bool("dryrun", false, "perform dryrun")
	cfgPath := flag.String("config", "~/.composer-sync.yaml", "config path")

	flag.Parse()

	cfg := &Config{}
	if err := config.Bind(*cfgPath, cfg); err != nil {
		fmt.Printf("Bad Config: %s", err)
		os.Exit(1)
	}

	// Pull subcommand & package name from args
	if flag.NArg() != 2 {
		fmt.Printf("expected 2 arguments\n")
		flag.Usage()
		os.Exit(1)
	}

	cmd := flag.Arg(0)
	name := flag.Arg(1)

	if cmd != "load" {
		fmt.Printf("invalid subcommand %s\n", cmd)
		flag.Usage()
		os.Exit(1)
	}

	st := stash.New(cfg.StashAPIUrl)
	pa := packagist.New()

	loader := &Loader{
		SClient:   st,
		PClient:   pa,
		DestTpl:   cfg.StashRepoTpl,
		StashProj: cfg.StashProj,
	}

	if err := loader.Load(name, *update, *dryrun); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
