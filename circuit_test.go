package literalcircuit

import (
	"testing"

	"github.com/xyproto/bits"
)

var and = &bits.TruthTable{
	"0 0 -> 0",
	"0 1 -> 0",
	"1 0 -> 0",
	"1 1 -> 1",
}

var xor = &bits.TruthTable{
	"0 0 -> 0",
	"0 1 -> 1",
	"1 0 -> 1",
	"1 1 -> 0",
}

var mainGateTable = &GateTable{
	"i0 -> and.0, and.0 -> o0",
	"i1 -> and.1, and.1 -> o1",
}

// For testing the mainGateTable
var testTruthTable = &bits.TruthTable{
	"0 0 -> 0",
	"0 1 -> 0",
	"1 0 -> 0",
	"1 1 -> 1",
}

func TestRegister(t *testing.T) {
	circuit := New()
	circuit.RegisterTruthTable("xor", xor)
	circuit.RegisterTruthTable("and", and)
	circuit.RegisterGateTable("main", mainGateTable)
}
