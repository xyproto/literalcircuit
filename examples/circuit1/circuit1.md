<!-- literate circuit -->

# Circuit 1

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

    i0 -> A.i0:A.o0 -> B.i0:B.o0 -> o0
    i1 -> A.i1:
    i2 -> C.i0:C.o0 -> B.i1:
    i3 -> C.i1:

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
