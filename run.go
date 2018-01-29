package literalcircuit

import (
	"fmt"
	"github.com/xyproto/bits"
	"strings"
)

// Run the main GateTable, return the bits
func (c *Circuit) Run(inputRouting []string, outputNames []string) *bits.Bits {
	//mainGateTable := c.mainGateTable
	for _, inputRoute := range inputRouting {
		if !strings.Contains(inputRoute, "->") {
			panic("invalid input routing: " + inputRoute)
		} else {
			fmt.Println("ROUTE", inputRoute)

		}
	}
	return &bits.Bits{}
}
