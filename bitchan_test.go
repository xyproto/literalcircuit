package literalcircuit

import (
	"github.com/xyproto/bits"
	"log"
	"testing"
)

func TestWrapAnd(t *testing.T) {
	log.Println("AND")

	i0 := make(BitChan, 1)
	i1 := make(BitChan, 1)
	o := make(BitChan, 1)

	// Set up a stopping mechanism
	stop := make(StopChan, 32) // buffered

	i := BitChans{i0, i1}
	go WrapTruthTable(and)(i, o, stop)

	i0 <- bits.B1
	i1 <- bits.B1

	// Block until we receive an output bit on o
	result := <-o

	log.Println("Got output bit", result)

	// Stop the and gate from processing
	stop <- true

	if result != bits.B1 {
		t.Error("and 1 1 should return 1")
	}
}

func TestWrapXor(t *testing.T) {
	log.Println("XOR")

	i0 := make(BitChan, 1)
	i1 := make(BitChan, 1)
	o := make(BitChan, 1)

	stop := make(StopChan, 32) // buffered

	i0 <- bits.B0
	i1 <- bits.B0

	go WrapTruthTable(xor)(BitChans{i0, i1}, o, stop)

	// Block until we receive an output bit on o
	result := <-o

	log.Println("Got output bit", result)

	// Then stop the gate
	stop <- true

	if result != bits.B0 {
		t.Error("xor 0 0 should return 0")
	}
}

//func TestSpew(t *testing.T) {
//	// Set up circuit input bits
//	I0 := make(BitChan) // unbuffered, only one at the time
//	I1 := make(BitChan) // unbuffered, only one at the time
//
//	stop := make(StopChan) // unbuffered, only one at the time
//
//	go SpewBitsFromString("1 0", BitChans{I0, I1}, stop)
//
//	var a, b bits.Bit
//	for i := 0; i < 10; i++ {
//		a = <-I0
//		b = <-I1
//		log.Println("A B", a, b)
//	}
//
//	stop <- true
//}

//func TestWrapCombine(t *testing.T) {
//	// Set up circuit input bits
//	I0 := make(BitChan, 1) // size
//	I1 := make(BitChan, 1) // size
//
//	// Stopping mechanism
//	stop := make(StopChan, 1) // size
//	stopConsumers := 0        // gate counter, used when stopping all of them
//
//	// ----------
//
//	// Set up input/output bits and run the xor gate as a goroutine
//	xorI0 := I0
//	xorI1 := I1
//	xorO0 := make(BitChan, 1) // size
//	go WrapTruthTable(xor)(BitChans{xorI0, xorI1}, xorO0, stop)
//	stopConsumers++
//
//	// Set up input/output bits and run the xor gate as a goroutine
//	andI0 := xorO0
//	andI1 := I0               // Duplicate input bit 0 as and input bit 1 (will be fed B1 in a loop)
//	andO0 := make(BitChan, 1) // size
//	go WrapTruthTable(and)(BitChans{andI0, andI1}, andO0, stop)
//	stopConsumers++
//
//	// Input the input bits into the circuit, until stopped
//	go SpewBitsFromString("1 0", BitChans{I0, I1}, stop)
//	stopConsumers++
//
//	// Set up the circuit output bit
//	O0 := andO0
//
//	// ----------
//
//	// Block until we receive an output bit on o0
//	result := <-O0
//
//	log.Println("Got output bit", result)
//
//	// Then stop the gates
//	for x := 0; x < stopConsumers; x++ {
//		stop <- true
//	}
//
//	if result != bits.B1 {
//		t.Error("and(xor(1, 0), 1) should return 1")
//	}
//}
