package main

import (
	"github.com/lucasloureiror/slh/internal/cli"
	"os"
)

func main() {

	version:= os.Getenv("VERSION")
	if version == "" {
		version = "local dev build"
	}
	cli.Start(version)

}
