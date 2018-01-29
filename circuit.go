package literalcircuit

import (
	"github.com/xyproto/bits"

	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
)

// Circuit is a collection of components (truth tables that act as functions, such as "xor"),
// a collection of connections between components (gate table expressions)
// and the name of the main list of gate table expressions, if there are several disconnected circuits.
type Circuit struct {
	gateMap        map[string]*bits.TruthTable
	connMap        map[string]*GateTable
	mainGateTable  *GateTable       // pointer to the main GateTable of the circuit
	testTruthTable *bits.TruthTable // pointer to the TruthTable that is used for testing the circuit
}

// New creates a new circuit, which can have several available components and several named collections of connections (circuits / gate table collections)
func New() *Circuit {
	var c Circuit
	c.gateMap = make(map[string]*bits.TruthTable)
	c.connMap = make(map[string]*GateTable)
	return &c
}

// RegisterTruthTable registers a gate, like "and" or "or", by using a name and a truth table
func (c *Circuit) RegisterTruthTable(name string, tt *bits.TruthTable) {
	// If this panics, the circuit must be made with NewCircuit instead of &Circuit{}
	c.gateMap[name] = tt
	if c.testTruthTable == nil || name == "test" {
		c.testTruthTable = tt
	}
}

// RegisterGateTable registers connections between gates, with a name and a gate table
func (c *Circuit) RegisterGateTable(name string, gt *GateTable) {
	// If this panics, the circuit must be made with NewCircuit instead of &Circuit{}
	c.connMap[name] = gt
	// Use this as the main circuit if no others are defined, or if this name is "main"
	if c.mainGateTable == nil || name == "main" {
		c.mainGateTable = gt
	}
}

// SetMain selects one of the registered GateTable names as the main GateTable of the circuit
func (c *Circuit) SetMain(name string) {
	if gt, ok := c.connMap[name]; ok {
		c.mainGateTable = gt
	} else {
		panic(name + " does not exist in the list of connections/GateTables")
	}
}

// Load a circuit file (literal circuit file: both Markdown and a circuit)
func Load(filename string, verbose bool) (*Circuit, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var (
		inGateTable  bool
		inTruthTable bool
		name         string
	)
	truthTable := &bits.TruthTable{}
	gateTable := &GateTable{}
	c := New()
	// Appending "# done" to the data is a sneaky way of not having an
	// additional check after the for loop for unprocessed items that were
	// gathered but not registered.
	data = append(data, []byte("# done")...)
	for _, byteLine := range bytes.Split(data, []byte("\n")) {
		line := string(byteLine)
		if strings.HasPrefix(line, "    ") && strings.Contains(line, "->") {
			// This is a line in a truth table or gate table
			if !inGateTable && !inTruthTable && strings.Contains(line, ":") {
				inGateTable = true
			} else if !inGateTable && !inTruthTable {
				inTruthTable = true
			}
			if inGateTable {
				*gateTable = append(*gateTable, strings.TrimSpace(line))
			} else if inTruthTable {
				*truthTable = append(*truthTable, strings.TrimSpace(line))
			}
		} else if strings.HasPrefix(line, "# ") {
			// Use the results that has been gathered so far
			if inTruthTable {
				// Register all the names for the truthTable
				if strings.Contains(name, ":") {
					firstName := strings.Split(name, ":")[0]
					if verbose {
						fmt.Println("Registering a truth table for " + firstName)
					}
					c.RegisterTruthTable(strings.TrimSpace(firstName), truthTable)
					aliases := strings.Split(strings.Split(name, ":")[1], ",")
					for _, alias := range aliases {
						if verbose {
							fmt.Println("Registering a truth table for " + strings.TrimSpace(alias))
						}
						c.RegisterTruthTable(strings.TrimSpace(alias), truthTable)
					}
				} else {
					if verbose {
						fmt.Println("Registering a truth table for " + name)
					}
					c.RegisterTruthTable(name, truthTable)
				}
			} else if inGateTable {
				if verbose {
					fmt.Println("Registering a gate table for " + name)
				}
				c.RegisterGateTable(name, gateTable)
			}
			// Reset the fields
			if inTruthTable || inGateTable {
				inTruthTable = false
				inGateTable = false
				name = ""
				truthTable = &bits.TruthTable{}
				gateTable = &GateTable{}
			}
			// Assign a new name
			name = strings.TrimSpace(line[2:])
		}
	}
	return &Circuit{}, nil
}
