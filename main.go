package main

import (
	"os"

	"github.com/userosettadev/rosetta-cli/internal"
)

func main() {

	os.Exit(internal.Run(os.Args, os.Stdout, os.Stderr))
}
