package main

import (
	"github.com/lyqscmy/pbcheck/pbcheck"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(pbcheck.Analyzer) }
