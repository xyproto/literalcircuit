package main

import (
	"fmt"
	docopt "github.com/docopt/docopt-go"
)

func main() {
	usage := `Literal Circuit Parser

Usage:
  lcp [FILENAME]
  naval_fate -h | --help
  naval_fate --version

Options:
  -h --help     Show this screen.
  --version     Show version.`

	arguments, _ := docopt.ParseDoc(usage)
	fmt.Println(arguments)
}
