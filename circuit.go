package circuit

import (
	"github.com/xyproto/bits"

	"fmt"
	"strings"
)

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

type GateTable []string

type Circuit struct {
	components  map[string]*bits.TruthTable
	connections map[string]*GateTable
	mainCircuit string // the name of a GateTable in the connections map
}

type Inputs []string // Maps input bits to input ports, like this: "1 -> main.0" (send "1" to first port of "main")
type OutputNames []string

func NewCircuit() *Circuit {
	var c Circuit
	c.components = make(map[string]*bits.TruthTable)
	c.connections = make(map[string]*GateTable)
	return &c
}

func (c *Circuit) Register(name string, tt *bits.TruthTable) {
	c.components[name] = tt
}

func (c *Circuit) RegisterConnection(name string, gt *GateTable) {
	c.connections[name] = gt
	// Use this as the main circuit if no others are defined
	if c.mainCircuit == "" {
		c.mainCircuit = name
	}
}

// SetMain selects one of the registered GateTable names as the main circuit
func (c *Circuit) SetMain(name string) {
	if _, ok := c.connections[name]; ok {
		c.mainCircuit = name
	} else {
		panic(name + " does not exist in the list of connections/GateTables")
	}
}

// Run the main, return the bits
func (c *Circuit) Run(inputRouting []string, outputNames []string) *bits.Bits {
	name := c.mainCircuit
	for _, inputRoute := range inputRouting {
		if !strings.Contains(inputRoute, "->") {
			panic("invalid input routing: " + inputRoute)
		} else {
			fmt.Println(name, "ROUTE", inputRoute)
		}
	}
	return &bits.Bits{}
}
