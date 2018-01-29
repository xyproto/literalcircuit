package literalcircuit

import (
	"testing"
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
	output := circuit.Run("1 0")
	if output != "0" {
		t.Error("CIRCUIT RUN FAILED")
	}
}
