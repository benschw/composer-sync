## composer-sync

cli tool to sync php composer packages from the vcs repo registered in packagist to 
a private stash repo


update `cfg.yaml` (and optionally copy it to its default location: `~/.composer-sync.yaml`)

### Download
- [composer-sync (linux/amd64)](https://drone.io/github.com/benschw/composer-sync/files/composer-sync.linux-amd64.gz)
- [composer-sync (darwin/amd64)](https://drone.io/github.com/benschw/composer-sync/files/composer-sync.darwin-amd64.gz)

- [cfg.yaml](https://drone.io/github.com/benschw/composer-sync/files/cfg.yaml)

### Usage

	composer-sync [OPTIONS]... VENDOR/PACKAGE

	  -u       sync all packages, not just new one
	  -r       recursively load transitive dependencies
	  -dryrun  show add/sync/skip but don't do anything
	  -config  specify a config other that ~/.composer-sync.yaml


### Examples
	
	composer-sync fliglio/web
	add:  fliglio_web 

	composer-sync -r fliglio/web
	skip: fliglio_web 
	add:  doctrine_cache 
	add:  symfony_validator 
	add:  symfony_translation 
	add:  symfony_polyfill-mbstring 
	add:  fliglio_http 
	add:  doctrine_annotations 
	add:  doctrine_lexer

	composer-sync -r fliglio/web
	skip: fliglio_web 
	skip: doctrine_cache 
	skip: symfony_validator 
	skip: symfony_translation 
	skip: symfony_polyfill-mbstring 
	skip: fliglio_http 
	skip: doctrine_annotations 
	skip: doctrine_lexer

	composer-sync -ur fliglio/web
	sync: fliglio_web 
	sync: doctrine_cache 
	sync: symfony_validator 
	sync: symfony_translation 
	sync: symfony_polyfill-mbstring 
	sync: fliglio_http 
	sync: doctrine_annotations 
	sync: doctrine_lexer


