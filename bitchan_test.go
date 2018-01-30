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
	I0 := make(BitChan, 1) // size
	I1 := make(BitChan, 1) // size

	// ----------

	// Set up input/output bits and run the xor gate as a goroutine
	xorI0 := I0
	xorI1 := I1
	xorO0 := make(BitChan, 1) // size
	go WrapTruthTable(xor)(BitChans{xorI0, xorI1}, xorO0)

	// Set up input/output bits and run the xor gate as a goroutine
	andI0 := xorO0
	andI1 := I0               // Duplicate input bit 0 as and input bit 1 (will be fed B1 in a loop)
	andO0 := make(BitChan, 1) // size
	go WrapTruthTable(and)(BitChans{andI0, andI1}, andO0)

	// Input the input bits into the circuit, for N cycles
	go func(cycles int) {
		for n := 0; n < cycles; n++ {
			I0 <- bits.B1
			I1 <- bits.B0
		}
	}(100)

	// Set up the circuit output bit
	O0 := andO0

	// ----------

	// Block until we receive an output bit on o0
	result := <-O0

	if result != bits.B1 {
		t.Error("and(xor(1, 0), 1) should return 1")
	}
}
