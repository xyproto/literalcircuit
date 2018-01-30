package literalcircuit

import (
	"testing"
)

func TestSelftest(t *testing.T) {
	circuit := New()
	circuit.RegisterTruthTable("and", and)
	circuit.RegisterTruthTable("xor", xor)
	circuit.RegisterGateTable("main", mainGateTable)
	circuit.RegisterTruthTable("test", testTruthTable)
	//if !circuit.SelfTest() {
	//	t.Error("CIRCUIT SELF TEST FAILED")
	//}
}
