# Artificial Intelligence Puzzle

Source code for my AI class to solve sliding-pieces puzzles.

## Input, with examples

For any given puzzle, the pieces will be ordered positive numbers starting at 1. Zero will be considered the "empty" piece,
and we move it to solve the puzzle.

### 3x3 matrix, input

```
[1, 2, 3]
[4, 5, 6]
[7, 0, 8]
```

### 3x3 matrix, solution

```
[1, 2, 3]
[4, 5, 6]
[7, 8, 0] # move 0 to the right
```
