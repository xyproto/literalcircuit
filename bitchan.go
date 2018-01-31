package literalcircuit

import (
	"fmt"
	"github.com/xyproto/bits"
	"log"
)

type (
	// BitChan is a channel that may return a bit
	BitChan chan bits.Bit

	// BitChans is a slice of several channels that each may return a bit
	BitChans [](chan bits.Bit)

	// StopChan is a boolean channel for making signal propagation stop
	StopChan chan bool

	// TickChan is for sending clock ticks, true for high, false for low
	TickChan chan bool

	// ChanGate is a gate that deals with channels of bits instead of bits
	ChanGate func(input BitChans, output BitChan, stop StopChan, tick TickChan)
)

// WrapOneToManyGate is a wrapper function that takes a OneToManyGate and
// returns a ChanGate instead. It deals with channels of bits instead of bits,
// and waits for bits to be received before using the given OneToManyGate to
// process the inputs and return an output.
func WrapOneToManyGate(gateName string, gate bits.OneToManyGate) ChanGate {
	return func(inputChans BitChans, output BitChan, stop StopChan, tick TickChan) {
		// Continuously gather all needed input bits and output the result as an output bit
		for {
			// Stop if a stop signal was received
			select {
			case <-stop:
				// Stop looping
				return
			// Must have a default case for the select-case to be non-blocking!
			default:
				// Gather input bits
				inputBits := make(bits.Bits, len(inputChans))
				for i, inputChan := range inputChans {
					// New input bit, blocking
					inputBits[i] = <-inputChan
					log.Printf(gateName+": Input %d\t: %v\n", i, inputBits[i])
				}
				// Process the input bits, wait for a clock tick, and output the result bit
				outputBit := gate(inputBits)
				<-tick
				log.Printf(gateName+": Output\t\t: %v\n", outputBit)
				output <- outputBit
			}
		}
	}
}

// WrapTruthTable takes a truth table that represents a function (like "and"
// or "xor") and returns a gate that accepts channels of bits instead of bits.
// This is useful for simulating circuits.
func WrapTruthTable(gateName string, tt *bits.TruthTable) ChanGate {
	return WrapOneToManyGate(gateName, tt.Gate())
}

// SpewBits continously outputs the given bits (as a string) to the
// given BitChans, until stopped by a "true" being sent to the stop channel.
func SpewBits(inputBits *bits.Bits, outputBitChans BitChans, stop StopChan, tick TickChan) {
	if len(*inputBits) != len(outputBitChans) {
		panic(fmt.Sprintf("Wrong number of input bits, should be %d: %v", len(outputBitChans), inputBits))
	}
	for {
		// Check if we should stop
		select {
		case <-stop:
			// Close all output channels
			for _, outputChan := range outputBitChans {
				close(outputChan)
			}

			// Stop looping
			return
		// Must have a default case for the select-case to be non-blocking!
		default:
			// Wait for a clock tick
			<-tick
			// Output all given bits
			for i, b := range *inputBits {
				outputBitChans[i] <- b
			}
		}
	}
}

// SpewBitsFromString continously outputs the given bits (as a string) to the
// given BitChans, until stopped by a "true" being sent to the stop channel.
func SpewBitsFromString(inputBitString string, outputBitChans BitChans, stop StopChan, tick TickChan) {
	inputBits, err := bits.StringToBits(inputBitString)
	if err != nil {
		panic("Invalid input bit string: " + inputBitString)
	}
	if len(*inputBits) != len(outputBitChans) {
		panic(fmt.Sprintf("Wrong number of input bits, should be %d: %s", len(outputBitChans), inputBitString))
	}
	for {
		// Check if we should stop
		select {
		case <-stop:
			// Close all output channels
			for _, outputChan := range outputBitChans {
				close(outputChan)
			}

			// Stop looping
			return
		// Must have a default case for the select-case to be non-blocking!
		default:
			// Wait for a clock tick
			<-tick
			// Output all given bits
			for i, b := range *inputBits {
				outputBitChans[i] <- b
			}
		}
	}
}
