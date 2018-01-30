package literalcircuit

import (
	"testing"

	"github.com/xyproto/bits"
)

func TestTest(t *testing.T) {
	circuit := New()
	circuit.RegisterTruthTable("and", and)
	circuit.RegisterTruthTable("xor", xor)
	circuit.RegisterGateTable("main", mainGateTable)
	circuit.RegisterTruthTable("test", testTruthTable)
	if !circuit.SelfTest() {
		t.Error("CIRCUIT SELF TEST FAILED")
	}
}

func TestRun(t *testing.T) {
	circuit := New()
	circuit.RegisterTruthTable("xor", xor)
	circuit.RegisterTruthTable("and", and)
	circuit.RegisterGateTable("main", mainGateTable)
	b, err := String2Bits("1 0")
	if err != nil {
		t.Error(err)
	}
	output := circuit.Run(b, 100)
	if output != bits.B0 {
		t.Error("CIRCUIT RUN FAILED")
	}
}
