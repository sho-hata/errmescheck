package main

import (
	"github.com/sho-hata/errmescheck"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(errmescheck.Analyzer) }
