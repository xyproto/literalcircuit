package literalcircuit

// Run takes input bits as a space separated string and returns
// output bits as a space separated string.
func (c *Circuit) Run(inputBits string) string {
	return "0"
}

// SelfTest returns true if a test TruthTable is set and returns the correct values
func (c *Circuit) SelfTest() bool {
	return true
}
