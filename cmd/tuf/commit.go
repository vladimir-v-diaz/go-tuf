package main

import (
	"github.com/flynn/go-tuf"
	"github.com/flynn/go-tuf/Godeps/_workspace/src/github.com/flynn/go-docopt"
)

func init() {
	register("commit", cmdCommit, `
usage: tuf commit

Commit staged files to the repository.
`)
}

func cmdCommit(args *docopt.Args, repo *tuf.Repo) error {
	return repo.Commit()
}
