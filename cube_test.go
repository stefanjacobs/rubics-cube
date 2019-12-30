package main

import (
	"fmt"
	"testing"
)

func TestCube(t *testing.T) {
	fmt.Println("Running 3x3x3 Testsuite")
	cube := Cube{
		top:    [][]Color{{blue, orange, yellow}, {green, white, orange}, {orange, orange, yellow}},
		bottom: [][]Color{{green, blue, white}, {orange, yellow, red}, {blue, blue, red}},
		left:   [][]Color{{white, red, green}, {red, blue, blue}, {yellow, green, white}},
		right:  [][]Color{{orange, white, blue}, {green, green, green}, {red, yellow, yellow}},
		front:  [][]Color{{white, blue, green}, {red, orange, yellow}, {red, yellow, blue}},
		back:   [][]Color{{red, yellow, orange}, {white, red, white}, {green, white, orange}}}

	topCW := cube.turnTopCW()
	// fmt.Println(topCW.showDetails())
	if topCW.top[0][0] != orange || topCW.front[0][2] != blue || topCW.left[0][1] != blue {
		t.Errorf("Turning top clockwise did not work properly")
	}

	topCCW := cube.turnTopCCW()
	// fmt.Println(topCCW.showDetails())
	if topCCW.top[0][0] != yellow || topCCW.front[0][2] != green || topCCW.left[0][1] != yellow {
		t.Errorf("Turning top clockwise did not work properly")
	}

	botCW := cube.turnBottomCW()
	// fmt.Println(botCW.showDetails())
	if botCW.bottom[0][0] != blue || botCW.front[2][1] != green || botCW.right[2][0] != red {
		t.Errorf("Turning bottom clockwise did not work properly")
	}

	botCCW := cube.turnBottomCCW()
	// fmt.Println(botCCW.showDetails())
	if botCCW.bottom[0][0] != white || botCCW.front[2][0] != red || botCCW.right[2][0] != green {
		t.Errorf("Turning bottom counter clockwise did not work properly")
	}

// -- bis hier validiert - 
	frontCW := cube.turnFrontCW()
	// fmt.Println(frontCW.showDetails())
	if frontCW.top[2][2] != green || frontCW.front[0][0] != red || frontCW.right[1][0] != orange {
		t.Errorf("Turning front clockwise did not work properly")
	}

	frontCCW := cube.turnFrontCCW()
	// fmt.Println(frontCCW.showDetails())
	if frontCCW.top[2][1] != green || frontCCW.bottom[0][0] != green || frontCCW.right[2][0] != green {
		t.Errorf("Turning front counterclockwise did not work properly")
	}

	backCW := cube.turnBackCW()
	// fmt.Println(backCW.showDetails())
	if backCW.top[0][0] != blue || backCW.bottom[2][2] != yellow || backCW.left[1][0] != orange {
		t.Errorf("Turning back clockwise did not work properly")
	}

	backCCW := cube.turnBackCCW()
	//fmt.Println(backCCW.showDetails())
	if backCCW.top[0][0] != yellow || backCCW.bottom[2][1] != green || backCCW.right[2][2] != yellow {
		t.Errorf("Turning back counterclockwise did not work properly")
	}
}
