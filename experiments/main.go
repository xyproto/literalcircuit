package main

import (
	"fmt"
)

type Bit int
type Bits []Bit

func ProcessAnd(bits Bits) Bit {
	B0 := Bit(0)
	B1 := Bit(1)
	for _, bit := range bits {
		if int(bit) != int(B1) {
			// One bit is not 1
			return B0
		}
	}
	return B1
}

type Gate func(Bits) Bit
type ChanGate func(input chan Bits, output chan Bits)

func WrapGate(f func(Bits) Bit) ChanGate {
	return func(input chan Bits, output chan Bits) {
		input_bits := <-input
		output <- Bits{f(input_bits)}
	}
}



func main() {

    // Start a worker goroutine, giving it the channel to
    // notify on.
    i := make(chan Bits, 1) // size
    o := make(chan Bits, 1) // size

    i <- Bits{Bit(1), Bit(1)}

    go WrapGate(ProcessAnd)(i, o)

    // Block until we receive an output on and_o0
    result := <-o

    fmt.Println("Result:", result)
}

