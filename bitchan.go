package literalcircuit

import (
	"fmt"
	"github.com/xyproto/bits"
	"log"
)

// BitChan is a channel that may return a bit
type BitChan chan bits.Bit

// BitChans is a slice of several channels that each may return a bit
type BitChans [](chan bits.Bit)

// StopChan is a boolean channel for making signal propagation stop
type StopChan chan bool

// ChanGate is a gate that deals with channels of bits instead of bits
type ChanGate func(input BitChans, output BitChan, stop StopChan)

// WrapOneToManyGate is a wrapper function that takes a OneToManyGate and
// returns a ChanGate instead. It deals with channels of bits instead of bits,
// and waits for bits to be received before using the given OneToManyGate to
// process the inputs and return an output.
func WrapOneToManyGate(gateName string, gate bits.OneToManyGate) ChanGate {
	log.Println("WrapOneToManyGate:", gateName)
	return func(inputChans BitChans, output BitChan, stop StopChan) {
		// Continuously gather all needed input bits and output the result as an output bit
		for {
			// Gather input bits
			inputBits := make(bits.Bits, len(inputChans))
			for i, inputChan := range inputChans {
				// New input bit, blocking
				inputBits[i] = <-inputChan
				log.Printf(gateName+": Input %d\t: %v\n", i, inputBits[i])
			}
			// Process the input bits and output the result bit
			outputBit := gate(inputBits)

			log.Printf(gateName+": Output\t\t: %v\n", outputBit)

			output <- outputBit
			// Stop if a stop signal was received
			select {
			case <-stop:
				log.Println(gateName + ": Got stop signal")
				// Stop looping
				return
			// Must have a default case for the select-case to be non-blocking!
			default:
				log.Println(gateName + ": No stop signal")
			}
		}
	}
}

// WrapTruthTable takes a truth table that represents a function (like "and"
// or "xor") and returns a gate that accepts channels of bits instead of bits.
// This is useful for simulating circuits.
func WrapTruthTable(gateName string, tt *bits.TruthTable) ChanGate {
	log.Println("WrapOneToManyGate:", gateName)
	return WrapOneToManyGate(gateName, tt.Gate())
}

// SpewBits continously outputs the given bits (as a string) to the
// given BitChans, until stopped by a "true" being sent to the stop channel.
func SpewBits(inputBits *bits.Bits, outputBitChans BitChans, stop StopChan) {
	log.Println("SpewBits")
	if len(*inputBits) != len(outputBitChans) {
		panic(fmt.Sprintf("Wrong number of input bits, should be %d: %v", len(outputBitChans), inputBits))
	}
	for {
		// Output all given bits
		for i, b := range *inputBits {
			outputBitChans[i] <- b
			log.Printf("SpewBits: Outputted %v\n", b)
		}
		// Check if we should stop
		select {
		case <-stop:
			log.Println("SpewBits: Got stop signal")

			// Close all output channels
			for _, outputChan := range outputBitChans {
				log.Println("SpewBits: Closing output channel")
				close(outputChan)
			}

			// Stop looping
			return
		// Must have a default case for the select-case to be non-blocking!
		default:
			log.Println("SpewBits: No stop signal")
		}
	}
}

// SpewBitsFromString continously outputs the given bits (as a string) to the
// given BitChans, until stopped by a "true" being sent to the stop channel.
func SpewBitsFromString(inputBitString string, outputBitChans BitChans, stop StopChan) {
	inputBits, err := bits.StringToBits(inputBitString)
	if err != nil {
		panic("Invalid input bit string: " + inputBitString)
	}
	if len(*inputBits) != len(outputBitChans) {
		panic(fmt.Sprintf("Wrong number of input bits, should be %d: %s", len(outputBitChans), inputBitString))
	}
	for {
		// Output all given bits
		for i, b := range *inputBits {
			outputBitChans[i] <- b
			log.Printf("SpewBitsFromString: Outputted %v\n", b)
		}
		// Check if we should stop
		select {
		case <-stop:
			log.Println("SpewBitsFromString: Got stop signal")
			// Close all output channels
			for _, outputChan := range outputBitChans {
				log.Println("SpewBitsFromString: Closing output channel")
				close(outputChan)
			}

			// Stop looping
			return
		// Must have a default case for the select-case to be non-blocking!
		default:
			log.Println("SpewBits: No stop signal")
		}
	}
}
