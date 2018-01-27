<!-- literate circuit -->

# Circuit 1

* Regular text is ignored.
* Logic gates are defined with "# name:", a comma separated list of aliases (names) and then a truth table (indented 4 spaces).
* `-`, `|`, `/` and `\` connects gates.
* `i0`, `i1`, `i2` etc. are the input bits.
* `o0`, `o1`, `o2` etc. are the output bits.
* Gate aliases can be used in the circuit definition.

```none
    i0 ----\
            A------B-------- o0
    i1 ----/      /
                 /
    i2 ----C----/
          /
    i3 --/
```

Alternative syntax:

    i0 -> Ai0:Ao0 -> Bi0:Bo0 -> o0
    i1 -> Ai1:
    i2 -> Ci0:Co0 -> Bi1:
    i3 -> Ci1:

# or: A, C

    0 0 -> 0
    0 1 -> 1
    1 0 -> 1
    1 1 -> 1

# and: B

    0 0 -> 0
    0 1 -> 0
    1 0 -> 0
    1 1 -> 1

# test

    0 0 0 0 -> 0
    1 0 0 1 -> 1
    0 0 1 0 -> 0
    1 1 1 1 -> 1
