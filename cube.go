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
			result[j][il-i-1] = layer[i][j]
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
			result[jl-j-1][i] = layer[i][j]
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

// turn Top Clock Wise
func (cube Cube) turnTopCW() Cube {
	newCube := Cube{
		top:    clockwise(cube.top),
		bottom: duplicate(cube.bottom),
		front:  duplicate(cube.front),
		left:   duplicate(cube.left),
		back:   duplicate(cube.back),
		right:  duplicate(cube.right)}

	// Surrounding layers have to be moved:
	buffer := newCube.front[0]
	newCube.front[0] = newCube.right[0]
	newCube.right[0] = newCube.back[0]
	newCube.back[0] = newCube.left[0]
	newCube.left[0] = buffer

	return newCube
}

// turn Top Counterclock Wise
func (cube Cube) turnTopCCW() Cube {
	newCube := Cube{
		top:    counterClockwise(cube.top),
		bottom: duplicate(cube.bottom),
		front:  duplicate(cube.front),
		left:   duplicate(cube.left),
		back:   duplicate(cube.back),
		right:  duplicate(cube.right)}

	// Surrounding layers have to be moved:
	buffer := newCube.front[0]
	newCube.front[0] = newCube.left[0]
	newCube.left[0] = newCube.back[0]
	newCube.back[0] = newCube.right[0]
	newCube.right[0] = buffer

	return newCube
}

// turn Bottom Clock Wise
func (cube Cube) turnBottomCW() Cube {
	newCube := Cube{
		top:    duplicate(cube.top),
		bottom: clockwise(cube.bottom),
		front:  duplicate(cube.front),
		left:   duplicate(cube.left),
		back:   duplicate(cube.back),
		right:  duplicate(cube.right)}

	// Surrounding layers have to be moved:
	lastLayer := len(newCube.front) - 1
	buffer := newCube.front[lastLayer]
	newCube.front[lastLayer] = newCube.left[lastLayer]
	newCube.left[lastLayer] = newCube.back[lastLayer]
	newCube.back[lastLayer] = newCube.right[lastLayer]
	newCube.right[lastLayer] = buffer

	return newCube
}

// turn Bottom Counterclock Wise
func (cube Cube) turnBottomCCW() Cube {
	newCube := Cube{
		top:    duplicate(cube.top),
		bottom: counterClockwise(cube.bottom),
		front:  duplicate(cube.front),
		left:   duplicate(cube.left),
		back:   duplicate(cube.back),
		right:  duplicate(cube.right)}

	// Surrounding layers have to be moved:
	lastLayer := len(newCube.front) - 1
	buffer := newCube.front[lastLayer]
	newCube.front[lastLayer] = newCube.right[lastLayer]
	newCube.right[lastLayer] = newCube.back[lastLayer]
	newCube.back[lastLayer] = newCube.left[lastLayer]
	newCube.left[lastLayer] = buffer

	return newCube
}

// turn Front clockwise
func (cube Cube) turnFrontCW() Cube {
	newCube := Cube{
		top:    duplicate(cube.top),
		bottom: duplicate(cube.bottom),
		front:  clockwise(cube.front),
		left:   duplicate(cube.left),
		back:   duplicate(cube.back),
		right:  duplicate(cube.right)}

	// Surrounding layers have to be moved:
	len := len(newCube.front)
	bT := duplicate(newCube.top)
	bB := duplicate(newCube.bottom)
	for i := 0; i < len; i++ {
		newCube.top[len-1][i] = newCube.left[len-1-i][len-1]
		newCube.bottom[0][i] = newCube.right[len-1-i][0]
	}
	for i := 0; i < len; i++ {
		newCube.right[i][0] = bT[len-1][i]
		newCube.left[i][len-1] = bB[0][i]
	}

	return newCube
}

// turn Front counterclockwise
func (cube Cube) turnFrontCCW() Cube {
	newCube := Cube{
		top:    duplicate(cube.top),
		bottom: duplicate(cube.bottom),
		front:  counterClockwise(cube.front),
		left:   duplicate(cube.left),
		back:   duplicate(cube.back),
		right:  duplicate(cube.right)}

	// Surrounding layers have to be moved:
	len := len(newCube.front)
	bT := duplicate(newCube.top)
	bB := duplicate(newCube.bottom)
	for i := 0; i < len; i++ {
		newCube.top[len-1][i] = newCube.right[i][0]
		newCube.bottom[0][i] = newCube.left[i][len-1]
	}
	for i := 0; i < len; i++ {
		newCube.right[i][0] = bB[0][len-1-i]
		newCube.left[i][len-1] = bT[len-1][len-1-i]
	}

	return newCube
}

// turn back clockwise
func (cube Cube) turnBackCW() Cube {
	newCube := Cube{
		top:    duplicate(cube.top),
		bottom: duplicate(cube.bottom),
		front:  duplicate(cube.front),
		left:   duplicate(cube.left),
		back:   clockwise(cube.back),
		right:  duplicate(cube.right)}

	// Surrounding layers have to be moved:
	len := len(newCube.back)
	bT := duplicate(newCube.top)
	bB := duplicate(newCube.bottom)
	for i := 0; i < len; i++ {
		newCube.top[0][i] = newCube.right[i][len-1]
		newCube.bottom[len-1][i] = newCube.left[i][0]
	}
	for i := 0; i < len; i++ {
		newCube.right[i][len-1] = bB[len-1][len-1-i]
		newCube.left[i][0] = bT[0][len-1-i]
	}

	return newCube	
}

// turn back counterclockwise
func (cube Cube) turnBackCCW() Cube {
	newCube := Cube{
		top:    duplicate(cube.top),
		bottom: duplicate(cube.bottom),
		front:  duplicate(cube.front),
		left:   duplicate(cube.left),
		back:   counterClockwise(cube.back),
		right:  duplicate(cube.right)}

	// Surrounding layers have to be moved:
	len := len(newCube.back)
	bT := duplicate(newCube.top)
	bB := duplicate(newCube.bottom)
	for i := 0; i < len; i++ {
		newCube.top[0][i] = newCube.left[len-1-i][0]
		newCube.bottom[len-1][i] = newCube.right[len-1-i][len-1]
	}
	for i := 0; i < len; i++ {
		newCube.right[i][len-1] = bT[0][i]
		newCube.left[i][0] = bB[len-1][i]
	}

	return newCube	
}



// func (cube Cube) turnRightCW() Cube {}
// func (cube Cube) turnRightCCW() Cube {}
// func (cube Cube) turnLeftCW() Cube {}
// func (cube Cube) turnLeftCCW() Cube {}

// func (cube Cube) actionMidXCW(layer int) Cube {}
// func (cube Cube) actionMidXCCW(layer int) Cube {}
// func (cube Cube) actionMidYCW(layer int) Cube {}
// func (cube Cube) actionMidYCCW(layer int) Cube {}
// func (cube Cube) actionMidZCW(layer int) Cube {}
// func (cube Cube) actionMidZCCW(layer int) Cube {}
