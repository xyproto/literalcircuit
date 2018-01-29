package main

import (
	"fmt"
	lc "github.com/xyproto/literalcircuit"
)

func main() {
	fn := "circuit1.md"
	fmt.Printf("Loading literal circuit: %s...", fn)
	_, err := lc.Load("circuit1.md")
	if err != nil {
		panic(err)
	}
	//fmt.Println("Retrieving \"or\"-gate...")
	//orGate := cf.Gate("or")
	//fmt.Println("Running tests...")
	//fmt.Println(cf.Test())
	fmt.Println("done")
}
