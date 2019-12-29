package main

import "fmt"

// Cube is a representation of a Rubiks Cube
type Cube struct {
	top, bottom, left, right, front, back [][]Color
}

func (cube Cube) showDetails() string {
	desc := ""
	desc += "Top: " + fmt.Sprintf("%v", cube.top) + "\n"
	desc += "Fro: " + fmt.Sprintf("%v", cube.front) + " Right: " + fmt.Sprintf("%v", cube.right) + " "
	desc += "Back: " + fmt.Sprintf("%v", cube.back) + " Left: " + fmt.Sprintf("%v", cube.left) + "\n"
	desc += "Bot: " + fmt.Sprintf("%v", cube.bottom) + "\n"
	return desc
}

// Turn one layer clockwise. The layers next to it are not accounted for, so this is not a 
//   cube method!
func clockwise(layer [][]Color) [][]Color {
    il := len(layer)
    jl := len(layer[0])
    result := make([][]Color, il)
    for i := range result {
        result[i] = make([]Color, jl)
    }
    for i := 0; i < il; i++ {
        for j := 0; j < jl; j++ {
            result[j][il - i - 1] = layer[i][j]
        }
    }
    return result
}

// Turn one layer counter clockwise. The layers next to it are not accounted for, so this is not a 
//   cube method!
func counterClockwise(layer [][]Color) [][]Color {
    il := len(layer)
    jl := len(layer[0])
    result := make([][]Color, il)
    for i := range result {
        result[i] = make([]Color, jl)
    }
    for i := 0; i < il; i++ {
        for j := 0; j < jl; j++ {
            result[jl - j - 1][i] = layer[i][j]
        }
    }
    return result
}

// Duplicate a layer
func duplicate(layer [][]Color) [][]Color {
	duplicate := make([][]Color, len(layer))
	for i := range layer {
		duplicate[i] = make([]Color, len(layer[i]))
		copy(duplicate[i], layer[i])
	}
	return duplicate
}


// Action Top Clock Wise
func (cube Cube) actionTopCW() Cube {
	newCube := Cube{
		top: clockwise(cube.top),
		bottom: duplicate(cube.bottom),
		front: duplicate(cube.front),
		left: duplicate(cube.left),
		back: duplicate(cube.back),
		right: duplicate(cube.right)}
	
	// Surrounding layers have to be moved:
	buffer := newCube.front[0]
	newCube.front[0] = newCube.right[0]
	newCube.right[0] = newCube.back[0]
	newCube.back[0] = newCube.left[0]
	newCube.left[0] = buffer

    return newCube
}

// Action Top Counterclock Wise
func (cube Cube) actionTopCCW() Cube {
	newCube := Cube{
		top: counterClockwise(cube.top),
		bottom: duplicate(cube.bottom),
		front: duplicate(cube.front),
		left: duplicate(cube.left),
		back: duplicate(cube.back),
		right: duplicate(cube.right)}
	
	// Surrounding layers have to be moved:
	buffer := newCube.front[0]
	newCube.front[0] = newCube.left[0]
	newCube.left[0] = newCube.back[0]
	newCube.back[0] = newCube.right[0]
	newCube.right[0] = buffer

    return newCube
}