package main

import (
	"fmt"

	"github.com/xyproto/bits"
)

type ChanGate func(input chan bits.Bits, output chan bits.Bits)

func WrapOneToManyGate(gate bits.OneToManyGate) ChanGate {
	return func(input chan bits.Bits, output chan bits.Bits) {
		input_bits := <-input
		output <- bits.Bits{gate(input_bits)}
	}
}

func WrapTruthTable(tt *bits.TruthTable) ChanGate {
	return WrapOneToManyGate(tt.Gate())
}

var and = &bits.TruthTable{
	"0 0 -> 0",
	"0 1 -> 0",
	"1 0 -> 0",
	"1 1 -> 1",
}

func main() {

	// Start a worker goroutine, giving it the channel to
	// notify on.
	i := make(chan bits.Bits, 1) // size
	o := make(chan bits.Bits, 1) // size

	i <- bits.Bits{bits.B1, bits.B1}

	go WrapTruthTable(and)(i, o)

	// Block until we receive an output on and_o0
	result := <-o

	fmt.Println("Result:", result)
}
