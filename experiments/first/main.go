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

func ProcessAnd(inputBits bits.Bits) bits.Bit {
	for _, bit := range inputBits {
		if int(bit) != int(bits.B1) {
			// One bit is not 1
			return bits.B0
		}
	}
	return bits.B1
}

func main() {

	// Start a worker goroutine, giving it the channel to
	// notify on.
	i := make(chan bits.Bits, 1) // size
	o := make(chan bits.Bits, 1) // size

	i <- bits.Bits{bits.B1, bits.B1}

	go WrapOneToManyGate(ProcessAnd)(i, o)

	// Block until we receive an output on and_o0
	result := <-o

	fmt.Println("Result:", result)
}
