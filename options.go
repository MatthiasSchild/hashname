package main

import (
	"flag"
)

var (
	optionMethod    string
	optionExtension bool
	optionDry       bool
)

func init() {
	flag.StringVar(&optionMethod, "method", "sha1", "Hashing method")
	flag.BoolVar(&optionExtension, "ext", false, "Keep the extension of the file")
	flag.BoolVar(&optionDry, "dry", false, "Execute a dry run")
	flag.Parse()
}
