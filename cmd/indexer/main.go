package main

import (
	"os"

	"github.com/carlmjohnson/exitcode"
	"github.com/rugaviola/sdut-sourcesdb/pkg/indexer"
)

func main() {
	exitcode.Exit(indexer.CLI(os.Args[1:]))
}
