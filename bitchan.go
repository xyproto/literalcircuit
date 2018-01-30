package literalcircuit

import (
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
		// Continiously gather all needed input bits and output the result as an output bit
		for {
			// Gather input bits
			inputBits := make(bits.Bits, len(inputChans), len(inputChans))
			for i, inputChan := range inputChans {
				inputBits[i] = <-inputChan
			}
			// Process the input bits and output the result bit
			output <- gate(inputBits)
		}
	}
}

// WrapTruthTable takes a truth table that represents a function (like "and"
// or "xor") and returns a gate that accepts channels of bits instead of bits.
// This is useful for simulating circuits.
func WrapTruthTable(tt *bits.TruthTable) ChanGate {
	return WrapOneToManyGate(tt.Gate())
}
