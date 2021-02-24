package main

import (
	"os"

	"github.com/carlmjohnson/exitcode"
	"github.com/rugaviola/sdut-sourcesdb/pkg/autotweeter"
)

func main() {
	exitcode.Exit(autotweeter.CLI(os.Args[1:]))
}
