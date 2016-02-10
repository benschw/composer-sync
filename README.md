## composer-sync

cli tool to sync php composer packages from the vcs repo registered in packagist to 
a private stash repo


update `cfg.yaml` (and optionally copy it to its default location: `~/.composer-sync.yaml`)


composer-sync load fliglio/web
add:  fliglio_web 
add:  doctrine_cache 
add:  symfony_validator 
add:  symfony_translation 
add:  symfony_polyfill-mbstring 
add:  fliglio_http 
add:  doctrine_annotations 
add:  doctrine_lexer

composer-sync load fliglio/web
skip: fliglio_web 
skip: doctrine_cache 
skip: symfony_validator 
skip: symfony_translation 
skip: symfony_polyfill-mbstring 
skip: fliglio_http 
skip: doctrine_annotations 
skip: doctrine_lexer

composer-sync -u load fliglio/web
sync: fliglio_web 
sync: doctrine_cache 
sync: symfony_validator 
sync: symfony_translation 
sync: symfony_polyfill-mbstring 
sync: fliglio_http 
sync: doctrine_annotations 
sync: doctrine_lexer
