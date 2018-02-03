# Half Adder

### circuit

    i0 -> xor.i0:xor.o0 -> o0
    i1 -> xor.i1:
    i0 -> xor.i0:xor.o0 -> o1
    i1 -> xor.i1:

### xor

    0 0 -> 0
    0 1 -> 1
    1 0 -> 1
    1 1 -> 0

### and

    0 0 -> 0
    0 1 -> 0
    1 0 -> 0
    1 1 -> 1

### test

    0 0 -> 0 0
    0 1 -> 1 0
    1 0 -> 1 0
    1 1 -> 0 1
