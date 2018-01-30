package literalcircuit

import (
	"github.com/xyproto/bits"
	"testing"
)

func TestWrapAnd(t *testing.T) {
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

	if result != bits.B1 {
		t.Error("and 1 1 should return 1")
	}
}

func TestWrapXor(t *testing.T) {
	i0 := make(BitChan, 1) // size
	i1 := make(BitChan, 1) // size
	o := make(BitChan, 1)  // size

	i0 <- bits.B0
	i1 <- bits.B0

	go WrapTruthTable(xor)(BitChans{i0, i1}, o)

	// Block until we receive an output bit on o
	result := <-o

	if result != bits.B0 {
		t.Error("xor 0 0 should return 0")
	}
}

func TestWrapCombine(t *testing.T) {
	// Set up circuit input bits
	i0 := make(BitChan, 1) // size
	i1 := make(BitChan, 1) // size

	// ----------

	// Set up input/output bits and run the xor gate as a goroutine
	xor_i0 := i0
	xor_i1 := i1
	xor_o0 := make(BitChan, 1) // size
	go WrapTruthTable(xor)(BitChans{xor_i0, xor_i1}, xor_o0)

	// Set up input/output bits and run the xor gate as a goroutine
	and_i0 := xor_o0
	and_i1 := i0               // Duplicate input bit 0 as and input bit 1 (will be fed B1 in a loop)
	and_o0 := make(BitChan, 1) // size
	go WrapTruthTable(and)(BitChans{and_i0, and_i1}, and_o0)

	// Input the input bits into the circuit, for N cycles
	go func(cycles int) {
		for n := 0; n < cycles; n++ {
			i0 <- bits.B1
			i1 <- bits.B0
		}
	}(100)

	// Set up the circuit output bit
	o0 := and_o0

	// ----------

	// Block until we receive an output bit on o0
	result := <-o0

	if result != bits.B1 {
		t.Error("and(xor(1, 0), 1) should return 1")
	}
}
