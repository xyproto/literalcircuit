# Literal Circuit

## Literal Circuit File Format Spec

* Spec version: 0.1
* File extension: `.md`
* The syntax is compatible with Markdown.

## Defining a truth table

The name comes first, after `# `:

    # name

Then comes the truth table, with one or more inputs, and arrow, and one or more outputs:

    0 -> 1
    1 -> 0

### Example 1: xor

    # xor

    0 0 -> 0
    0 1 -> 1
    1 0 -> 1
    1 1 -> 0

### Example 2: and

    # and

    0 0 -> 0
    0 1 -> 0
    1 0 -> 0
    1 1 -> 1

### Example 3: or, with aliases A and B

    # or: A, B

    0 0 -> 0
    0 1 -> 1
    1 0 -> 1
    1 1 -> 1

* Every line in a truth table must have an arrow (`->`), separating input and output.
* If the name of a truth table has a colon (`:`), the rest of the line is expected to be a comma-separated list of aliases.

## GateTables

A circuit, with connections between gates, can be defined by a GateTable.

* The arrow `->` is used to show what connects to what.
* `i0` is the 0th input bit. `i1` is the 1st input bit. etc.
* `o0` is the 0th output bit. `o1` is the 1st output bit. etc.
* A colon (`:`)  can be used to seperate multiple statements on one line.
* A name or alias of a gate (as defined by a truth table) can be used in the circuit.
* `.i0` and `.o0` can be used together with names or aliases of gates to select input and output bit posisions.

### Example 1: connecting inputs to "A" and then "A" to outputs

    # OR Circuit 1

    i0 -> A.i0:A.o0 -> o0
    i1 -> A.i1:A.o1 -> o1

This connects input bit 0 to A (an alias for or, as defined above), then the output bit 0 from A to output bit 0 of the circuit.
The same is done for input bit 1, through A, and to output bit 1.

## Special names

* If a truth table is named "test", it will be used for testing the main circuit.
* If a GateTable is named "main", then that is considered to be the main circuit.
* If there is only one GateTable, then that is considered to be the main circuit.

## Regular text

* Text that does not begin with either `# ` or four spaces is ignored.
* This means that it's completely acceptible (and encouraged) to include an ASCII-art diagram of the circuit in the circuit file. Example:

```none
    i0 ----\
            A------B-------- o0
    i1 ----/      /
                 /
    i2 ----C----/
          /
    i3 --/
```

* Explanations, comments, license and other information can also be included in the circuit.

## Expansion

A word that does not end with a number, between two arrows, in the definition of a LogicGate, like this:

`i0 -> and -> i1`

Will be expanded like this:

`i0 -> and.i0:and.i1 -> i1`

# Go package

This repository includes the `literalcircuit` package and an interpreter in `cmd`.

# General information

* Version: 0.1
* License: MIT
* Author: &lt;xyproto@archlinux.org&gt;
