package literalcircuit

import (
	"fmt"
	"testing"

	"github.com/xyproto/bits"
)

var xor = &bits.TruthTable{
	"0 0 -> 0",
	"0 1 -> 1",
	"1 0 -> 1",
	"1 1 -> 0",
}

var and = &bits.TruthTable{
	"0 0 -> 0",
	"0 1 -> 0",
	"1 0 -> 0",
	"1 1 -> 1",
}

var halfAdder = &bits.TruthTable{
	"0 0 -> 0",
	"0 1 -> 1",
	"1 0 -> 1",
	"1 1 -> 0",
}

// i0 is the first input, i0 is the first output
// componentName.0 means the first input if on the left side of the arrow,
// and the first output if on the right side of the arrow.
// i means all the inputs
// o means all the outputs
// Expressions can be comma-separated
//
// "i0 -> and.0 -> o0" is a shortcut for:
//    "i0 -> and.0, and.0 -> o0"
//
// "i -> and -> o" is a shortcut for:
//    "i0 -> and.0, and.0 -> o0",
//    "i1 -> and.1, and.1 -> o1",
//    ...
//
var mainCircuit = &GateTable{
	"i0 -> and.0, and.0 -> o0",
	"i1 -> and.1, and.1 -> o1",
}

func TestAdder(t *testing.T) {
	circuit := NewCircuit()
	circuit.RegisterTruthTable("xor", xor)
	circuit.RegisterTruthTable("and", and)
	circuit.RegisterGateTable("main", mainCircuit)
	output := circuit.Run([]string{"1 -> main.0", "0 -> main.1"}, []string{"main.0", "main.1"})
	fmt.Println("The output after putting 1 into main.0 and 0 into main.1 is:", output)
}
