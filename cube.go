package main

import (
	"fmt"
	// "crypto/sha256"
	"encoding/binary"
	//	"math"
)

// Cube is a representation of a Rubiks Cube
type Cube struct {
	Top, Bottom, Left, Right, Front, Back [][]Color
	hash uint64
}

func (cube Cube) showDetails() string {
	desc := ""
	desc += "Top: " + fmt.Sprintf("%v", cube.Top) + "\n"
	desc += "Fro: " + fmt.Sprintf("%v", cube.Front) + " Right: " + fmt.Sprintf("%v", cube.Right) + " "
	desc += "Back: " + fmt.Sprintf("%v", cube.Back) + " Left: " + fmt.Sprintf("%v", cube.Left) + "\n"
	desc += "Bot: " + fmt.Sprintf("%v", cube.Bottom) + "\n"
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

// Check, if a layer is only of one color
func isUniformColor(layer [][]Color) bool {
	initColor := layer[0][0]
	for _, line := range layer {
		for _, c := range line {
			if c != initColor {
				return false
			}
		}
	}
	return true
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func maxOfSlice(x []int) int {
	maxVal := 0
	for _, value := range x {
		if value > maxVal {
			maxVal = value
		}
	}
	return maxVal
}

var powTable = make([]int, 200)

func intPow(x, p int) int {
	val := powTable[p]
	if val > 0 {
		return val
	}
	val = 1
	for i := 0; i < p; i++ {
		val = val * x
	}
	powTable[p] = val
	return val
}

// Return ident nr of layer
func ident(layer [][]Color) int {
	val := 0
	length := len(layer)
	for i, v := range layer {
		for j, w := range v {
			curPos := i*length+j
			val += w.Int() * intPow(8, curPos)
		}
	}
	fmt.Printf("%v\n", val)
	return val
}

// Return a heuristic. Max value is 3...
func heuristic(layer [][]Color) int {
	colorMap := map[Color]int{}
	for _, line := range layer {
		for _, element := range line {
			colorMap[element] = 1
		}
	}
	colorCount := 0
	for _, c := range colorMap {
		colorCount += c
	}
	return min(3, colorCount-1) // middle stone always has correct color, so decrease one
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
		Top:    clockwise(cube.Top),
		Bottom: duplicate(cube.Bottom),
		Front:  duplicate(cube.Front),
		Left:   duplicate(cube.Left),
		Back:   duplicate(cube.Back),
		Right:  duplicate(cube.Right)}

	// Surrounding layers have to be moved:
	buffer := newCube.Front[0]
	newCube.Front[0] = newCube.Right[0]
	newCube.Right[0] = newCube.Back[0]
	newCube.Back[0] = newCube.Left[0]
	newCube.Left[0] = buffer

	return newCube.hashify()
}

func (cube Cube) hashify() Cube {
	sideLength := len(cube.Top)
	var tmp []byte
	for i:=0; i < sideLength; i++ {
		for j:=0; j < sideLength; j++ {
			tmp = append(tmp, cube.Top[i][j].Byte())
			tmp = append(tmp, cube.Bottom[i][j].Byte())
			tmp = append(tmp, cube.Front[i][j].Byte())
			tmp = append(tmp, cube.Back[i][j].Byte())
			tmp = append(tmp, cube.Left[i][j].Byte())
			tmp = append(tmp, cube.Right[i][j].Byte())
		}
	}
	// fmt.Printf("Tmp: %v\n", tmp)
	cube.hash = uint64(binary.LittleEndian.Uint64(tmp))
	// fmt.Printf("Hash: %v\n", cube.hash)
	return cube
}

// turn Top Counterclock Wise
func (cube Cube) turnTopCCW() Cube {
	newCube := Cube{
		Top:    counterClockwise(cube.Top),
		Bottom: duplicate(cube.Bottom),
		Front:  duplicate(cube.Front),
		Left:   duplicate(cube.Left),
		Back:   duplicate(cube.Back),
		Right:  duplicate(cube.Right)}

	// Surrounding layers have to be moved:
	buffer := newCube.Front[0]
	newCube.Front[0] = newCube.Left[0]
	newCube.Left[0] = newCube.Back[0]
	newCube.Back[0] = newCube.Right[0]
	newCube.Right[0] = buffer

	return newCube.hashify()
}

// turn Bottom Clock Wise
func (cube Cube) turnBottomCW() Cube {
	newCube := Cube{
		Top:    duplicate(cube.Top),
		Bottom: clockwise(cube.Bottom),
		Front:  duplicate(cube.Front),
		Left:   duplicate(cube.Left),
		Back:   duplicate(cube.Back),
		Right:  duplicate(cube.Right)}

	// Surrounding layers have to be moved:
	lastLayer := len(newCube.Front) - 1
	buffer := newCube.Front[lastLayer]
	newCube.Front[lastLayer] = newCube.Left[lastLayer]
	newCube.Left[lastLayer] = newCube.Back[lastLayer]
	newCube.Back[lastLayer] = newCube.Right[lastLayer]
	newCube.Right[lastLayer] = buffer

	return newCube.hashify()
}

// turn Bottom Counterclock Wise
func (cube Cube) turnBottomCCW() Cube {
	newCube := Cube{
		Top:    duplicate(cube.Top),
		Bottom: counterClockwise(cube.Bottom),
		Front:  duplicate(cube.Front),
		Left:   duplicate(cube.Left),
		Back:   duplicate(cube.Back),
		Right:  duplicate(cube.Right)}

	// Surrounding layers have to be moved:
	lastLayer := len(newCube.Front) - 1
	buffer := newCube.Front[lastLayer]
	newCube.Front[lastLayer] = newCube.Right[lastLayer]
	newCube.Right[lastLayer] = newCube.Back[lastLayer]
	newCube.Back[lastLayer] = newCube.Left[lastLayer]
	newCube.Left[lastLayer] = buffer

	return newCube.hashify()
}

// turn Front clockwise
func (cube Cube) turnFrontCW() Cube {
	newCube := Cube{
		Top:    duplicate(cube.Top),
		Bottom: duplicate(cube.Bottom),
		Front:  clockwise(cube.Front),
		Left:   duplicate(cube.Left),
		Back:   duplicate(cube.Back),
		Right:  duplicate(cube.Right)}

	// Surrounding layers have to be moved:
	len := len(newCube.Front)
	bT := duplicate(newCube.Top)
	bB := duplicate(newCube.Bottom)
	for i := 0; i < len; i++ {
		newCube.Top[len-1][i] = newCube.Left[len-1-i][len-1]
		newCube.Bottom[0][i] = newCube.Right[len-1-i][0]
	}
	for i := 0; i < len; i++ {
		newCube.Right[i][0] = bT[len-1][i]
		newCube.Left[i][len-1] = bB[0][i]
	}

	return newCube.hashify()
}

// turn Front counterclockwise
func (cube Cube) turnFrontCCW() Cube {
	newCube := Cube{
		Top:    duplicate(cube.Top),
		Bottom: duplicate(cube.Bottom),
		Front:  counterClockwise(cube.Front),
		Left:   duplicate(cube.Left),
		Back:   duplicate(cube.Back),
		Right:  duplicate(cube.Right)}

	// Surrounding layers have to be moved:
	len := len(newCube.Front)
	bT := duplicate(newCube.Top)
	bB := duplicate(newCube.Bottom)
	for i := 0; i < len; i++ {
		newCube.Top[len-1][i] = newCube.Right[i][0]
		newCube.Bottom[0][i] = newCube.Left[i][len-1]
	}
	for i := 0; i < len; i++ {
		newCube.Right[i][0] = bB[0][len-1-i]
		newCube.Left[i][len-1] = bT[len-1][len-1-i]
	}

	return newCube.hashify()
}

// turn Back clockwise
func (cube Cube) turnBackCW() Cube {
	newCube := Cube{
		Top:    duplicate(cube.Top),
		Bottom: duplicate(cube.Bottom),
		Front:  duplicate(cube.Front),
		Left:   duplicate(cube.Left),
		Back:   clockwise(cube.Back),
		Right:  duplicate(cube.Right)}

	// Surrounding layers have to be moved:
	len := len(newCube.Back)
	bT := duplicate(newCube.Top)
	bB := duplicate(newCube.Bottom)
	for i := 0; i < len; i++ {
		newCube.Top[0][i] = newCube.Right[i][len-1]
		newCube.Bottom[len-1][i] = newCube.Left[i][0]
	}
	for i := 0; i < len; i++ {
		newCube.Right[i][len-1] = bB[len-1][len-1-i]
		newCube.Left[i][0] = bT[0][len-1-i]
	}

	return newCube.hashify()
}

// turn Back counterclockwise
func (cube Cube) turnBackCCW() Cube {
	newCube := Cube{
		Top:    duplicate(cube.Top),
		Bottom: duplicate(cube.Bottom),
		Front:  duplicate(cube.Front),
		Left:   duplicate(cube.Left),
		Back:   counterClockwise(cube.Back),
		Right:  duplicate(cube.Right)}

	// Surrounding layers have to be moved:
	len := len(newCube.Back)
	bT := duplicate(newCube.Top)
	bB := duplicate(newCube.Bottom)
	for i := 0; i < len; i++ {
		newCube.Top[0][i] = newCube.Left[len-1-i][0]
		newCube.Bottom[len-1][i] = newCube.Right[len-1-i][len-1]
	}
	for i := 0; i < len; i++ {
		newCube.Right[i][len-1] = bT[0][i]
		newCube.Left[i][0] = bB[len-1][i]
	}

	return newCube.hashify()
}

func (cube Cube) turnRightCW() Cube {
	newCube := Cube{
		Top:    duplicate(cube.Top),
		Bottom: duplicate(cube.Bottom),
		Front:  duplicate(cube.Front),
		Left:   duplicate(cube.Left),
		Back:   duplicate(cube.Back),
		Right:  clockwise(cube.Right)}

	// Surrounding layers have to be moved:
	len := len(newCube.Back)
	bT := duplicate(newCube.Top)
	bB := duplicate(newCube.Bottom)
	for i := 0; i < len; i++ {
		newCube.Top[i][len-1] = newCube.Front[i][len-1]
		newCube.Bottom[i][len-1] = newCube.Back[len-1-i][0]
	}
	for i := 0; i < len; i++ {
		newCube.Front[i][len-1] = bB[i][len-1]
		newCube.Back[len-1-i][0] = bT[i][len-1]
	}

	return newCube.hashify()
}

func (cube Cube) turnRightCCW() Cube {
	newCube := Cube{
		Top:    duplicate(cube.Top),
		Bottom: duplicate(cube.Bottom),
		Front:  duplicate(cube.Front),
		Left:   duplicate(cube.Left),
		Back:   duplicate(cube.Back),
		Right:  counterClockwise(cube.Right)}

	// Surrounding layers have to be moved:
	len := len(newCube.Back)
	bT := duplicate(newCube.Top)
	bB := duplicate(newCube.Bottom)
	for i := 0; i < len; i++ {
		newCube.Top[i][len-1] = newCube.Back[len-1-i][0]
		newCube.Bottom[i][len-1] = newCube.Front[i][len-1]
	}
	for i := 0; i < len; i++ {
		newCube.Front[i][len-1] = bT[i][len-1]
		newCube.Back[len-1-i][0] = bB[i][len-1]
	}

	return newCube.hashify()
}

func (cube Cube) turnLeftCW() Cube {
	newCube := Cube{
		Top:    duplicate(cube.Top),
		Bottom: duplicate(cube.Bottom),
		Front:  duplicate(cube.Front),
		Left:   clockwise(cube.Left),
		Back:   duplicate(cube.Back),
		Right:  duplicate(cube.Right)}

	// Surrounding layers have to be moved:
	len := len(newCube.Back)
	bT := duplicate(newCube.Top)
	bB := duplicate(newCube.Bottom)
	for i := 0; i < len; i++ {
		newCube.Top[i][0] = newCube.Back[len-1-i][len-1]
		newCube.Bottom[i][0] = newCube.Front[i][0]
	}
	for i := 0; i < len; i++ {
		newCube.Front[i][0] = bT[i][0]
		newCube.Back[len-1-i][len-1] = bB[i][0]
	}

	return newCube.hashify()
}

func (cube Cube) turnLeftCCW() Cube {
	newCube := Cube{
		Top:    duplicate(cube.Top),
		Bottom: duplicate(cube.Bottom),
		Front:  duplicate(cube.Front),
		Left:   counterClockwise(cube.Left),
		Back:   duplicate(cube.Back),
		Right:  duplicate(cube.Right)}

	// Surrounding layers have to be moved:
	len := len(newCube.Back)
	bT := duplicate(newCube.Top)
	bB := duplicate(newCube.Bottom)
	for i := 0; i < len; i++ {
		newCube.Top[i][0] = newCube.Front[i][0]
		newCube.Bottom[i][0] = newCube.Back[len-1-i][len-1]
	}
	for i := 0; i < len; i++ {
		newCube.Front[i][0] = bB[i][0]
		newCube.Back[len-1-i][len-1] = bT[i][0]
	}

	return newCube.hashify()
}

func (cube Cube) actionTopLayerCW(layer int) Cube {
	if layer == 0 {
		return cube.turnTopCW()
	} else if layer == len(cube.Top[0])-1 {
		return cube.turnBottomCCW()
	}
	newCube := Cube{
		Top:    duplicate(cube.Top),
		Bottom: duplicate(cube.Bottom),
		Front:  duplicate(cube.Front),
		Left:   duplicate(cube.Left),
		Back:   duplicate(cube.Back),
		Right:  duplicate(cube.Right)}

	// Surrounding layers have to be moved:
	buffer := newCube.Front[layer]
	newCube.Front[layer] = newCube.Right[layer]
	newCube.Right[layer] = newCube.Back[layer]
	newCube.Back[layer] = newCube.Left[layer]
	newCube.Left[layer] = buffer
	return newCube
}

func (cube Cube) actionTopLayerCCW(layer int) Cube {
	if layer == 0 {
		return cube.turnTopCCW()
	} else if layer == len(cube.Top[0])-1 {
		return cube.turnBottomCW()
	}
	newCube := Cube{
		Top:    duplicate(cube.Top),
		Bottom: duplicate(cube.Bottom),
		Front:  duplicate(cube.Front),
		Left:   duplicate(cube.Left),
		Back:   duplicate(cube.Back),
		Right:  duplicate(cube.Right)}

	// Surrounding layers have to be moved:
	buffer := newCube.Front[layer]
	newCube.Front[layer] = newCube.Left[layer]
	newCube.Left[layer] = newCube.Back[layer]
	newCube.Back[layer] = newCube.Right[layer]
	newCube.Right[layer] = buffer
	return newCube.hashify()
}

func (cube Cube) actionFrontLayerCW(layer int) Cube {
	if layer == 0 {
		return cube.turnFrontCW()
	} else if layer == len(cube.Top[0])-1 {
		return cube.turnBackCCW()
	}
	newCube := Cube{
		Top:    duplicate(cube.Top),
		Bottom: duplicate(cube.Bottom),
		Front:  duplicate(cube.Front),
		Left:   duplicate(cube.Left),
		Back:   duplicate(cube.Back),
		Right:  duplicate(cube.Right)}

	len := len(newCube.Front)
	bT := duplicate(newCube.Top)
	bB := duplicate(newCube.Bottom)
	for i := 0; i < len; i++ {
		newCube.Top[len-1-layer][i] = newCube.Left[len-1-i][len-1-layer]
		newCube.Bottom[layer][i] = newCube.Right[len-1-i][layer]
	}
	for i := 0; i < len; i++ {
		newCube.Right[i][layer] = bT[len-1-layer][i]
		newCube.Left[i][len-1-layer] = bB[layer][i]
	}

	return newCube.hashify()
}

func (cube Cube) actionFrontLayerCCW(layer int) Cube {
	if layer == 0 {
		return cube.turnFrontCCW()
	} else if layer == len(cube.Top[0])-1 {
		return cube.turnBackCW()
	}
	newCube := Cube{
		Top:    duplicate(cube.Top),
		Bottom: duplicate(cube.Bottom),
		Front:  duplicate(cube.Front),
		Left:   duplicate(cube.Left),
		Back:   duplicate(cube.Back),
		Right:  duplicate(cube.Right)}

	len := len(newCube.Front)
	bT := duplicate(newCube.Top)
	bB := duplicate(newCube.Bottom)
	for i := 0; i < len; i++ {
		newCube.Top[len-1-layer][i] = newCube.Right[i][layer]
		newCube.Bottom[layer][i] = newCube.Left[i][len-1-layer]
	}
	for i := 0; i < len; i++ {
		newCube.Right[i][layer] = bB[layer][len-1-i]
		newCube.Left[i][len-1-layer] = bT[len-1-layer][len-1-i]
	}

	return newCube.hashify()
}

func (cube Cube) actionRightLayerCW(layer int) Cube {
	if layer == 0 {
		return cube.turnRightCW()
	} else if layer == len(cube.Top[0])-1 {
		return cube.turnLeftCCW()
	}
	newCube := Cube{
		Top:    duplicate(cube.Top),
		Bottom: duplicate(cube.Bottom),
		Front:  duplicate(cube.Front),
		Left:   duplicate(cube.Left),
		Back:   duplicate(cube.Back),
		Right:  clockwise(cube.Right)}
	// Surrounding layers have to be moved:
	len := len(newCube.Back)
	bT := duplicate(newCube.Top)
	bB := duplicate(newCube.Bottom)
	for i := 0; i < len; i++ {
		newCube.Top[i][len-1-layer] = newCube.Front[i][len-1-layer]
		newCube.Bottom[i][len-1-layer] = newCube.Back[len-1-i][layer]
	}
	for i := 0; i < len; i++ {
		newCube.Front[i][len-1-layer] = bB[i][len-1-layer]
		newCube.Back[len-1-i][layer] = bT[i][len-1-layer]
	}

	return newCube.hashify()
}

func (cube Cube) actionRightLayerCCW(layer int) Cube {
	if layer == 0 {
		return cube.turnRightCCW()
	} else if layer == len(cube.Top[0])-1 {
		return cube.turnLeftCW()
	}
	newCube := Cube{
		Top:    duplicate(cube.Top),
		Bottom: duplicate(cube.Bottom),
		Front:  duplicate(cube.Front),
		Left:   duplicate(cube.Left),
		Back:   duplicate(cube.Back),
		Right:  clockwise(cube.Right)}
	// Surrounding layers have to be moved
	len := len(newCube.Back)
	bT := duplicate(newCube.Top)
	bB := duplicate(newCube.Bottom)
	for i := 0; i < len; i++ {
		newCube.Top[i][len-1-layer] = newCube.Back[len-1-i][layer]
		newCube.Bottom[i][len-1-layer] = newCube.Front[i][len-1-layer]
	}
	for i := 0; i < len; i++ {
		newCube.Front[i][len-1-layer] = bT[i][len-1-layer]
		newCube.Back[len-1-i][layer] = bB[i][len-1-layer]
	}

	return newCube.hashify()

}
