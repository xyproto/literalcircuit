package literalcircuit

import (
	"github.com/xyproto/bits"
	"log"
	"testing"
)

func clock(tick TickChan) {
	for {
		tick <- true
	}
}

func TestWrapAnd(t *testing.T) {
	i0 := make(BitChan, 1)
	i1 := make(BitChan, 1)
	o := make(BitChan, 1)

	// Set up a stopping mechanism
	stop := make(StopChan, 1)

	// Set up a clock tick
	tick := make(TickChan, 1)

	// TICK!
	go clock(tick)

	i := BitChans{i0, i1}
	go WrapTruthTable("AND", and)(i, o, stop, tick)

	i0 <- bits.B1
	i1 <- bits.B1

	// Block until we receive an output bit on o
	result := <-o

	// Stop the and gate from processing
	stop <- true

	if result != bits.B1 {
		t.Error("and 1 1 should return 1")
	}
}

func TestWrapXor(t *testing.T) {
	i0 := make(BitChan, 1)
	i1 := make(BitChan, 1)
	o := make(BitChan, 1)

	stop := make(StopChan, 1)
	tick := make(TickChan, 1)

	i0 <- bits.B0
	i1 <- bits.B0

	// TICK!
	go clock(tick)

	go WrapTruthTable("XOR", xor)(BitChans{i0, i1}, o, stop, tick)

	// Block until we receive an output bit on o
	result := <-o

	// Then stop the gate
	stop <- true

	if result != bits.B0 {
		t.Error("xor 0 0 should return 0")
	}
}

func TestSpew(t *testing.T) {
	// Set up circuit input bits
	I0 := make(BitChan, 1)
	I1 := make(BitChan, 1)

	stop := make(StopChan, 32)
	tick := make(TickChan, 32)

	// TICK!
	go clock(tick)

	go SpewBitsFromString("1 0", BitChans{I0, I1}, stop, tick)

	// Try to consume the outputted bits
	go func() {
		var a, b bits.Bit
		for i := 0; i < 10; i++ {
			a = <-I0
			b = <-I1
			if a != bits.B1 {
				t.Error("bit 0 should be 1")
			}
			if b != bits.B0 {
				t.Error("bit 1 should be 0")
			}
			//log.Printf("TestSpew: i=%v a=%v b=%v\n", i, a, b)
		}
		<-stop
	}()
}

func TestWrapCombine(t *testing.T) {
	// Set up circuit input bits

	I0 := make(BitChan)
	I1 := make(BitChan)

	// Stopping mechanism
	stop := make(StopChan, 1)
	stopConsumers := 0 // gate counter, used when stopping all of them

	tick := make(TickChan, 1)

	// TICK!
	go clock(tick)

	// ----------

	// Set up input/output bits and run the xor gate as a goroutine
	xorI0 := I0
	xorI1 := I1
	xorO0 := make(BitChan) // size
	go WrapTruthTable("XOR", xor)(BitChans{xorI0, xorI1}, xorO0, stop, tick)
	stopConsumers++

	// Set up input/output bits and run the xor gate as a goroutine
	andI0 := xorO0
	andI1 := I0                // Duplicate input bit 0 as and input bit 1 (will be fed B1 in a loop)
	andO0 := make(BitChan) // size
	go WrapTruthTable("AND", and)(BitChans{andI0, andI1}, andO0, stop, tick)
	stopConsumers++

	// Input the input bits into the circuit, until stopped
	go SpewBitsFromString("1 0", BitChans{I0, I1}, stop, tick)
	stopConsumers++

	// Set up the circuit output bit
	O0 := andO0

	// ----------

	// Block until we receive an output bit on o0
	result := <-O0

	//log.Println("Got output bit", result)

	// Then stop the gates
	for x := 0; x < stopConsumers; x++ {
		stop <- true
	}

	log.Println("RESULT", result)

	if result != bits.B1 {
		t.Error("and(xor(1, 0), 1) should return 1")
	}
}
