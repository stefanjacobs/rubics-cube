# Rubiks Cube solver

## Base Assumptions

- The cube has six colors
- The cube has six sides
- The dimensions are the same on all the sides
- The dimension of the cube is arbitrary, I know of 2x2x2, 3x3x3, 4x4x4 cubes that fit those assumptions

## Problem class

- State S_n
- Previous State S_n-1
- Actions {A_i}: S_n+1 := A_i(S_n)
- isFinal(S_n): bool
- cost: int
- heuristic: int (has to be always lower than estimate)

## Heuristic in Rubis Cube

- Max Number of colors on one layer is minimal amount of moves (is that so?)
- 