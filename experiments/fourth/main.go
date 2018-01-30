package main

import (
	"fmt"

	"github.com/xyproto/bits"
)

// A channel that may return a bit
type BitChan chan bits.Bit

// Several channels that may each return a bit
type BitChans [](chan bits.Bit)

// A gate that deals with channels of bits instead of bits
type ChanGate func(input BitChans, output BitChan)

// ChanGate is a wrapper function that takes a OneToManyGate and returns a
// ChanGate instead. It deals with channels of bits instead of bits, and waits
// for bits to be received before using the given OneToManyGate to process the
// inputs and return an output.
func WrapOneToManyGate(gate bits.OneToManyGate) ChanGate {
	return func(inputChans BitChans, output BitChan) {
		inputBits := make(bits.Bits, len(inputChans), len(inputChans))
		for i, inputChan := range inputChans {
			inputBits[i] = <-inputChan
		}
		output <- gate(inputBits)
	}
}

// WrapTruthTable takes a truth table that represents a function (like "and"
// or "xor") and returns a gate that accepts channels of bits instead of bits.
// This is useful for simulating circuits.
func WrapTruthTable(tt *bits.TruthTable) ChanGate {
	return WrapOneToManyGate(tt.Gate())
}

var and = &bits.TruthTable{
	"0 0 -> 0",
	"0 1 -> 0",
	"1 0 -> 0",
	"1 1 -> 1",
}

var or = &bits.TruthTable{
	"0 0 -> 0",
	"0 1 -> 1",
	"1 0 -> 1",
	"1 1 -> 1",
}

func main() {
	// Start a worker goroutine, giving it the channel to notify on

	i0 := make(BitChan, 1) // size
	i1 := make(BitChan, 1) // size
	o := make(BitChan, 1)  // size

	i0 <- bits.B1
	i1 <- bits.B1

	i := BitChans{i0, i1}

	go WrapTruthTable(and)(i, o)

	// Block until we receive an output bit on o
	result := <-o

	fmt.Println("Result:", result)
}
