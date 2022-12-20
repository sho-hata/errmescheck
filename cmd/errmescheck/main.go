package main

import (
	"github.com/sho-hata/errmescheck"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(errmescheck.Analyzer) }
