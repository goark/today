package main

import (
	"os"

	"github.com/goark/today/internal/facade"
)

// Version of the application for build info, will be replaced by build script
var Version = ""

func main() {
	os.Exit(facade.Execute(
		os.Stdin,
		os.Stdout,
		os.Stderr,
		Version,
		os.Args[1:],
	))
}
