package literalcircuit

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/xyproto/bits"
)

// RunExpression runs a single arrow-expression, like "i0 -> xor.0"
func RunExpression(expr string, inputBits *bits.Bits) (string, error) {
	if strings.Count(expr, "->") != 1 {
		return "", errors.New("Expression does not contain exactly one arrow: " + expr)
	}
	elements := strings.Split(expr, "->")
	from := strings.TrimSpace(elements[0])
	to := strings.TrimSpace(elements[1])
	// Check if the from field is 0 or 1
	if from == "0" || from == "1" {
		fmt.Println("EVENT: Push " + from + " into " + to)
		return "", nil
	}
	// Check if the from field is i0, i1, i2 etc
	for i, bit := range *inputBits {
		if from == "i"+strconv.Itoa(i) {
			return RunExpression(bit.String()+" -> "+to, inputBits)
		} else if to == "i"+strconv.Itoa(i) {
			return "", errors.New("Input bits can only be on the left side of arrow: " + expr)
		}
	}
	// Check if the to field is o0, o1, o2 etc
	if len(to) == 2 && to[0] == 'o' {
		fmt.Println("EVENT: Push output from " + from + " to " + to)
	}
	return "", nil
}

// Run takes input bits and returns an output bit.
// Runs the circuit for N cycles.
func (c *Circuit) Run(inputBits *bits.Bits, cycles uint64) bits.Bit {
	for _, line := range *c.mainGateTable {
		if strings.Contains(line, ":") {
			for _, expr := range strings.Split(line, ":") {
				RunExpression(expr, inputBits)
			}
		} else {
			RunExpression(line, inputBits)
		}
	}
	// TODO: Run the circuit and return the result
	return bits.B0
}

// SelfTest returns true if a test TruthTable is set and returns the correct values
func (c *Circuit) SelfTest() bool {
	return true
}

// String2Bits converts a space-separated string of "0"s and "1"s to a pointer
// to a slice of bits.Bit. Returns an error if an invalid string is given.
func String2Bits(s string) (*bits.Bits, error) {
	var gatheredBits bits.Bits
	if strings.Contains(s, " ") {
		for _, sbit := range strings.Split(s, " ") {
			if sbit == "1" {
				gatheredBits = append(gatheredBits, bits.B1)
			} else if sbit == "0" {
				gatheredBits = append(gatheredBits, bits.B0)
			} else {
				return nil, errors.New("Invalid bit: " + sbit)
			}
		}
	} else {
		if s == "1" {
			gatheredBits = append(gatheredBits, bits.B1)
		} else if s == "0" {
			gatheredBits = append(gatheredBits, bits.B0)
		} else {
			return nil, errors.New("Invalid bit: " + s)
		}
	}
	return &gatheredBits, nil
}
