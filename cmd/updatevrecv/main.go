package main

import (
	"updatevrecv"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(updatevrecv.Analyzer) }

