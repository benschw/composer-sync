package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"

	"github.com/benschw/composer-sync/config"
	"github.com/benschw/composer-sync/packagist"
	"github.com/benschw/composer-sync/stash"
	satis "github.com/benschw/satis-go/satis/client"
)

func main() {
	update := flag.Bool("u", false, "update existing repositories")
	recursive := flag.Bool("r", false, "Recusively operate on all dependencies")
	dryrun := flag.Bool("dryrun", false, "perform dryrun")
	cfgPath := flag.String("config", "~/.composer-sync.yaml", "config path")

	flag.Parse()
	if *cfgPath == "~/.composer-sync.yaml" {
		usr, err := user.Current()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		cfgPathStr := (usr.HomeDir + "/.composer-sync.yaml")
		cfgPath = &cfgPathStr
	}

	cfg, err := config.Load(*cfgPath)
	if err != nil {
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

	loginCmd := name == "login"

	if (cfg.Stash.Login == "" || cfg.Stash.Password == "") || loginCmd {
		if err := config.Login(cfg, *cfgPath, loginCmd); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if loginCmd {
			os.Exit(0)
		}
	}

	loader := &Loader{
		Stash:     stash.New(cfg.Stash.ApiUrl, cfg.Stash.Login, cfg.Stash.Password),
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
