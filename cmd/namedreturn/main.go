package main

import (
	"github.com/thara/namedreturn"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(namedreturn.Analyzer) }
