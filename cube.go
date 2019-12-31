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


func (cube Cube) turnRightCW() Cube {
	newCube := Cube{
		top:    duplicate(cube.top),
		bottom: duplicate(cube.bottom),
		front:  duplicate(cube.front),
		left:   duplicate(cube.left),
		back:   duplicate(cube.back),
		right:  clockwise(cube.right)}

	// Surrounding layers have to be moved:
	len := len(newCube.back)
	bT := duplicate(newCube.top)
	bB := duplicate(newCube.bottom)
	for i := 0; i < len; i++ {
		newCube.top[i][len-1] = newCube.front[i][len-1]
		newCube.bottom[i][len-1] = newCube.back[len-1-i][0]
	}
	for i := 0; i < len; i++ {
		newCube.front[i][len-1] = bB[i][len-1]
		newCube.back[len-1-i][0] = bT[i][len-1]
	}

	return newCube	
}


func (cube Cube) turnRightCCW() Cube {
	newCube := Cube{
		top:    duplicate(cube.top),
		bottom: duplicate(cube.bottom),
		front:  duplicate(cube.front),
		left:   duplicate(cube.left),
		back:   duplicate(cube.back),
		right:  counterClockwise(cube.right)}

	// Surrounding layers have to be moved:
	len := len(newCube.back)
	bT := duplicate(newCube.top)
	bB := duplicate(newCube.bottom)
	for i := 0; i < len; i++ {
		newCube.top[i][len-1] = newCube.back[len-1-i][0]
		newCube.bottom[i][len-1] = newCube.front[i][len-1] 
	}
	for i := 0; i < len; i++ {
		newCube.front[i][len-1] = bT[i][len-1]
		newCube.back[len-1-i][0] = bB[i][len-1]
	}

	return newCube	
}


func (cube Cube) turnLeftCW() Cube {
	newCube := Cube{
		top:    duplicate(cube.top),
		bottom: duplicate(cube.bottom),
		front:  duplicate(cube.front),
		left:   clockwise(cube.left),
		back:   duplicate(cube.back),
		right:  duplicate(cube.right)}

	// Surrounding layers have to be moved:
	len := len(newCube.back)
	bT := duplicate(newCube.top)
	bB := duplicate(newCube.bottom)
	for i := 0; i < len; i++ {
		newCube.top[i][0] = newCube.back[len-1-i][len-1]
		newCube.bottom[i][0] = newCube.front[i][0]
	}
	for i := 0; i < len; i++ {
		newCube.front[i][0] = bT[i][0]
		newCube.back[len-1-i][len-1] = bB[i][0]
	}

	return newCube	
}


func (cube Cube) turnLeftCCW() Cube {
	newCube := Cube{
		top:    duplicate(cube.top),
		bottom: duplicate(cube.bottom),
		front:  duplicate(cube.front),
		left:   counterClockwise(cube.left),
		back:   duplicate(cube.back),
		right:  duplicate(cube.right)}

	// Surrounding layers have to be moved:
	len := len(newCube.back)
	bT := duplicate(newCube.top)
	bB := duplicate(newCube.bottom)
	for i := 0; i < len; i++ {
		newCube.top[i][0] = newCube.front[i][0]
		newCube.bottom[i][0] = newCube.back[len-1-i][len-1]
	}
	for i := 0; i < len; i++ {
		newCube.front[i][0] = bB[i][0]
		newCube.back[len-1-i][len-1] = bT[i][0]
	}

	return newCube		
}

func (cube Cube) actionTopLayerCW(layer int) Cube {
	if layer == 0 {
		return cube.turnTopCW()
	} else if layer == len(cube.top[0]) - 1 {
		return cube.turnBottomCCW()
	}
	newCube := Cube{
		top:    duplicate(cube.top),
		bottom: duplicate(cube.bottom),
		front:  duplicate(cube.front),
		left:   duplicate(cube.left),
		back:   duplicate(cube.back),
		right:  duplicate(cube.right)}

	// Surrounding layers have to be moved:
	buffer := newCube.front[layer]
	newCube.front[layer] = newCube.right[layer]
	newCube.right[layer] = newCube.back[layer]
	newCube.back[layer] = newCube.left[layer]
	newCube.left[layer] = buffer
    return newCube
}

func (cube Cube) actionTopLayerCCW(layer int) Cube {
	if layer == 0 {
		return cube.turnTopCCW()
	} else if layer == len(cube.top[0]) - 1 {
		return cube.turnBottomCW()
	}
	newCube := Cube{
		top:    duplicate(cube.top),
		bottom: duplicate(cube.bottom),
		front:  duplicate(cube.front),
		left:   duplicate(cube.left),
		back:   duplicate(cube.back),
		right:  duplicate(cube.right)}

	// Surrounding layers have to be moved:
	buffer := newCube.front[layer]
	newCube.front[layer] = newCube.left[layer]
	newCube.left[layer] = newCube.back[layer]
	newCube.back[layer] = newCube.right[layer]
	newCube.right[layer] = buffer
    return newCube
}

func (cube Cube) actionFrontLayerCW(layer int) Cube {
	if layer == 0 {
		return cube.turnFrontCW()
	} else if layer == len(cube.top[0]) - 1 {
		return cube.turnBackCCW()
	}
	newCube := Cube{
		top:    duplicate(cube.top),
		bottom: duplicate(cube.bottom),
		front:  duplicate(cube.front),
		left:   duplicate(cube.left),
		back:   duplicate(cube.back),
		right:  duplicate(cube.right)}

	len := len(newCube.front)
	bT := duplicate(newCube.top)
	bB := duplicate(newCube.bottom)
	for i := 0; i < len; i++ {
		newCube.top[len-1-layer][i] = newCube.left[len-1-i][len-1-layer]
		newCube.bottom[layer][i] = newCube.right[len-1-i][layer]
	}
	for i := 0; i < len; i++ {
		newCube.right[i][layer] = bT[len-1-layer][i]
		newCube.left[i][len-1-layer] = bB[layer][i]
	}

	return newCube
}

func (cube Cube) actionFrontLayerCCW(layer int) Cube {
	if layer == 0 {
		return cube.turnFrontCCW()
	} else if layer == len(cube.top[0]) - 1 {
		return cube.turnBackCW()
	}
	newCube := Cube{
		top:    duplicate(cube.top),
		bottom: duplicate(cube.bottom),
		front:  duplicate(cube.front),
		left:   duplicate(cube.left),
		back:   duplicate(cube.back),
		right:  duplicate(cube.right)}
	
	len := len(newCube.front)
	bT := duplicate(newCube.top)
	bB := duplicate(newCube.bottom)
	for i := 0; i < len; i++ {
		newCube.top[len-1-layer][i] = newCube.right[i][layer]
		newCube.bottom[layer][i] = newCube.left[i][len-1-layer]
	}
	for i := 0; i < len; i++ {
		newCube.right[i][layer] = bB[layer][len-1-i]
		newCube.left[i][len-1-layer] = bT[len-1-layer][len-1-i]
	}

	return newCube
}

func (cube Cube) actionRightLayerCW(layer int) Cube {
	if layer == 0 {
		return cube.turnRightCW()
	} else if layer == len(cube.top[0]) - 1 {
		return cube.turnLeftCCW()
	}
	newCube := Cube{
		top:    duplicate(cube.top),
		bottom: duplicate(cube.bottom),
		front:  duplicate(cube.front),
		left:   duplicate(cube.left),
		back:   duplicate(cube.back),
		right:  clockwise(cube.right)}
	// Surrounding layers have to be moved:
	len := len(newCube.back)
	bT := duplicate(newCube.top)
	bB := duplicate(newCube.bottom)
	for i := 0; i < len; i++ {
		newCube.top[i][len-1-layer] = newCube.front[i][len-1-layer]
		newCube.bottom[i][len-1-layer] = newCube.back[len-1-i][layer]
	}
	for i := 0; i < len; i++ {
		newCube.front[i][len-1-layer] = bB[i][len-1-layer]
		newCube.back[len-1-i][layer] = bT[i][len-1-layer]
	}

	return newCube	
}


func (cube Cube) actionRightLayerCCW(layer int) Cube {
	if layer == 0 {
		return cube.turnRightCCW()
	} else if layer == len(cube.top[0]) - 1 {
		return cube.turnLeftCW()
	}
	newCube := Cube{
		top:    duplicate(cube.top),
		bottom: duplicate(cube.bottom),
		front:  duplicate(cube.front),
		left:   duplicate(cube.left),
		back:   duplicate(cube.back),
		right:  clockwise(cube.right)}
	// Surrounding layers have to be moved
	len := len(newCube.back)
	bT := duplicate(newCube.top)
	bB := duplicate(newCube.bottom)
	for i := 0; i < len; i++ {
		newCube.top[i][len-1-layer] = newCube.back[len-1-i][layer]
		newCube.bottom[i][len-1-layer] = newCube.front[i][len-1-layer] 
	}
	for i := 0; i < len; i++ {
		newCube.front[i][len-1-layer] = bT[i][len-1-layer]
		newCube.back[len-1-i][layer] = bB[i][len-1-layer]
	}

	return newCube	

}
