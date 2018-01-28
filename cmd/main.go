package main

import (
	"fmt"
	"github.com/xyproto/circuit"
)

func main() {
	fn := "circuit1.md"
	fmt.Printf("Loading literal circuit: %s...", fn)
	_, err := circuit.Load("circuit1.md")
	if err != nil {
		panic(err)
	}
	//fmt.Println("Retrieving \"or\"-gate...")
	//orGate := cf.Gate("or")
	//fmt.Println("Running tests...")
	//fmt.Println(cf.Test())
	fmt.Println("done")
}
