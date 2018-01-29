package literalcircuit

import (
	"strings"
)

// GateTable is a collection of gate table expressions
// Gate table expressions can be comma-separated.
// There can be any number of arrows and commas in a GateTable statement.
type GateTable []string

// Output as a comma separated string
func (gt *GateTable) String() string {
	return strings.Join(*gt, ", ")
}
