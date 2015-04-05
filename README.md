pkgdex
====================

pkgdex generates a static website that serves as a canonical import path for your Go libraries per http://golang.org/cmd/go/#hdr-Remote_import_paths

## What
Turn a collection of json files that describe packages into HTML files & directory structure that you can "go get"

## Why
* Canonical import path that you can point to different source code repos. Github getting DDoSed? Repoint to self-hosted git repo and keep running your CI
* Wouldn't you rather users import your API library from my.company.com/myapi instead of github.com/company/myapi?

## How
1. Get the repo, compile, stash the binary somewhere in your path.
2. Look at example_src, maybe copy it somewhere and build from it. 
 * pkgdex-prefs.json holds global configs. Change the title, remove the template config line or customize the template
 * look at crowd.json & loginshare.json. Adapt these to your own packages. 
3. Run: pkgdex -dest output_dir source_dir
4. Upload output_dir somewhere usable. I rsync to a VM. You could push to S3.

## Future Work
This was hacked up late at night and I'm sure it has rough edges. Pull Requests are welcome. I really need to godoc it up and link to the docs for the json config structs to explain all the options.
