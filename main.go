package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/benschw/composer-sync/packagist"
	"github.com/benschw/composer-sync/stash"
)

func main() {
	update := flag.Bool("u", false, "update existing repositories")
	dryrun := flag.Bool("dryrun", false, "perform dryrun")

	destTpl := "http://foo:asdf@localhost:7990/scm/phpv/%s.git"
	flag.Parse()

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

	st := stash.New("http://foo:asdf@localhost:7990")
	pa := packagist.New()

	loader := &Loader{
		SClient:   st,
		PClient:   pa,
		DestTpl:   destTpl,
		StashProj: "PHPV",
	}

	if err := loader.Load(name, *update, *dryrun); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
