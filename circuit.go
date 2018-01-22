package circuit

import (
	"github.com/xyproto/bits"

	"fmt"
	"strings"
)

// GateTable is a collection of gate table expressions
//
// i0 is the first input, i0 is the first output
// componentName.0 means the first input if on the left side of the arrow,
// and the first output if on the right side of the arrow.
// i means all the inputs
// o means all the outputs
//
// "i0 -> and.0 -> o0" is a shortcut for:
//    "i0 -> and.0, and.0 -> o0"
//
// "i -> and -> o" is a shortcut for:
//    "i0 -> and.0, and.0 -> o0",
//    "i1 -> and.1, and.1 -> o1",
//    ...
//
// Gate table expressions can be comma-separated.
// There can be any number of arrows and commas in a GateTable statement.
type GateTable []string

// Circuit is a collection of components (truth tables that act as functions, such as "xor"),
// a collection of connections between components (gate table expressions)
// and the name of the main list of gate table expressions, if there are several disconnected circuits.
type Circuit struct {
	components  map[string]*bits.TruthTable
	connections map[string]*GateTable
	mainCircuit string // the name of a GateTable in the connections map
}

// Inputs maps input bits to input ports, like this: "1 -> main.0" (send "1" to first port of "main")
type Inputs []string

// NewCircuit creates a new circuit, which can have several available components and several named collections of connections (circuits / gate table collections)
func NewCircuit() *Circuit {
	var c Circuit
	c.components = make(map[string]*bits.TruthTable)
	c.connections = make(map[string]*GateTable)
	return &c
}

// Register a component with a name and a truth table
func (c *Circuit) Register(name string, tt *bits.TruthTable) {
	c.components[name] = tt
}

// Register a connection with a name and a gate table
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
