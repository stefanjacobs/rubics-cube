package main

import (
//	"fmt"
	"testing"
)

// Tests cube actions with some exemplary samples
func TestCube(t *testing.T) {
	// fmt.Println("Running 3x3x3 Testsuite")
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
	// fmt.Println(backCCW.showDetails())
	if backCCW.top[0][0] != yellow || backCCW.bottom[2][1] != green || backCCW.right[2][2] != yellow {
		t.Errorf("Turning back counterclockwise did not work properly")
	}

	rightCW := cube.turnRightCW()
	// fmt.Println(rightCW.showDetails())
	if rightCW.top[2][2] != blue || rightCW.front[0][2] != white || rightCW.back[0][0] != yellow {
		t.Errorf("Turning right clockwise did not work properly")
	}

	rightCCW := cube.turnRightCCW()
	// fmt.Println(rightCCW.showDetails())
	if rightCCW.top[2][2] != red || rightCCW.front[0][2] != yellow || rightCCW.back[0][0] != red {
		t.Errorf("Turning right counter clockwise did not work properly")
	}

	leftCW := cube.turnLeftCW()
	// fmt.Println(leftCW.showDetails())
	if leftCW.top[0][0] != orange || leftCW.front[2][0] != orange || leftCW.back[1][2] != orange {
		t.Errorf("Turning left clockwise did not work properly")
	}

	leftCCW := cube.turnLeftCCW()
	// fmt.Println(leftCCW.showDetails())
	if leftCCW.top[0][0] != white || leftCCW.front[2][0] != blue || leftCCW.back[1][2] != green {
		t.Errorf("Turning left counter clockwise did not work properly")
	}
}

// todo: automate tests, no output
func TestCubeActions(t *testing.T) {
	// fmt.Println("Running 3x3x3 Testsuite")
	cube := Cube{
		top:    [][]Color{{white, white, white}, {white, white, white}, {white, white, white}},
		bottom: [][]Color{{yellow, yellow, yellow}, {yellow, yellow, yellow}, {yellow, yellow, yellow}},
		front:  [][]Color{{orange, orange, orange}, {orange, orange, orange}, {orange, orange, orange}},
		right:  [][]Color{{green, green, green}, {green, green, green}, {green, green, green}},
		back:   [][]Color{{red, red, red}, {red, red, red}, {red, red, red}},
		left:   [][]Color{{blue, blue, blue}, {blue, blue, blue}, {blue, blue, blue}}}

    topCW0 := cube.actionTopLayerCW(0)
	// fmt.Println(topCW0.showDetails())
	if topCW0.front[0][0] != green || topCW0.left[0][0] != orange || topCW0.back[0][2] != blue {
		t.Errorf("Turning top clockwise did not work properly")
	}
    topCW1 := cube.actionTopLayerCW(1)
	// fmt.Println(topCW1.showDetails())
	if topCW1.front[1][0] != green || topCW1.left[1][1] != orange || topCW1.back[1][2] != blue {
		t.Errorf("Turning top clockwise did not work properly")
	}
	topCW2 := cube.actionTopLayerCW(2)
	// fmt.Println(topCW2.showDetails())
	if topCW2.front[2][0] != green || topCW2.left[2][2] != orange || topCW2.back[2][2] != blue {
		t.Errorf("Turning top clockwise did not work properly")
	}

    topCCW0 := cube.actionTopLayerCCW(0)
	// fmt.Println(topCCW0.showDetails())
	if topCCW0.front[0][0] != blue || topCCW0.left[0][0] != red || topCCW0.back[0][0] != green {
		t.Errorf("Turning top counter clockwise did not work properly")
	}
    topCCW1 := cube.actionTopLayerCCW(1)
	// fmt.Println(topCCW1.showDetails())
	if topCCW1.front[1][0] != blue || topCCW1.left[1][0] != red || topCCW1.back[1][1] != green {
		t.Errorf("Turning top counter clockwise did not work properly")
	}
	topCCW2 := cube.actionTopLayerCCW(2)
	// fmt.Println(topCCW2.showDetails())
	if topCCW2.front[2][0] != blue || topCCW2.left[2][0] != red || topCCW2.back[2][2] != green {
		t.Errorf("Turning top counter clockwise did not work properly")
	}

    frontCW0 := cube.actionFrontLayerCW(0)
	// fmt.Println(frontCW0.showDetails())
	if frontCW0.top[2][2] != blue || frontCW0.left[1][2] != yellow || frontCW0.bottom[0][0] != green {
		t.Errorf("Turning front clockwise did not work properly")
	}
    frontCW1 := cube.actionFrontLayerCW(1)
    // fmt.Println(frontCW1.showDetails())
	if frontCW1.top[1][1] != blue || frontCW1.left[1][1] != yellow || frontCW1.bottom[1][1] != green {
		t.Errorf("Turning front clockwise did not work properly")
	}
    frontCW2 := cube.actionFrontLayerCW(2)
	// fmt.Println(frontCW2.showDetails())
	if frontCW2.top[0][0] != blue || frontCW2.left[2][0] != yellow || frontCW2.bottom[2][2] != green {
		t.Errorf("Turning front clockwise did not work properly")
	}

    frontCCW0 := cube.actionFrontLayerCCW(0)
	// fmt.Println(frontCCW0.showDetails())
	if frontCCW0.top[2][0] != green || frontCCW0.left[0][2] != white || frontCCW0.bottom[0][2] != blue {
		t.Errorf("Turning front counter clockwise did not work properly")
	}
    frontCCW1 := cube.actionFrontLayerCCW(1)
	// fmt.Println(frontCCW1.showDetails())
	if frontCCW1.top[1][1] != green || frontCCW1.left[0][1] != white || frontCCW1.bottom[1][1] != blue {
		t.Errorf("Turning front counter clockwise did not work properly")
	}
	frontCCW2 := cube.actionFrontLayerCCW(2)
	// fmt.Println(frontCCW2.showDetails())
	if frontCCW2.top[0][0] != green || frontCCW2.left[0][0] != white || frontCCW2.bottom[2][2] != blue {
		t.Errorf("Turning front counter clockwise did not work properly")
	}

    rightCW0 := cube.actionRightLayerCW(0)
	// fmt.Println(rightCW0.showDetails())
	if rightCW0.top[2][2] != orange || rightCW0.front[0][2] != yellow || rightCW0.back[0][0] != white {
		t.Errorf("Turning right clockwise did not work properly")
	}
	rightCW1 := cube.actionRightLayerCW(1)
	// fmt.Println(rightCW1.showDetails())
	if rightCW1.top[2][1] != orange || rightCW1.front[0][1] != yellow || rightCW1.back[0][1] != white {
		t.Errorf("Turning right clockwise did not work properly")
	}
	rightCW2 := cube.actionRightLayerCW(2)
	// fmt.Println(rightCW2.showDetails())
	if rightCW2.top[2][0] != orange || rightCW2.front[0][0] != yellow || rightCW2.back[0][2] != white {
		t.Errorf("Turning right clockwise did not work properly")
	}

    rightCCW0 := cube.actionRightLayerCCW(0)
	// fmt.Println(rightCCW0.showDetails())
	if rightCCW0.top[0][2] != red || rightCCW0.front[2][2] != white || rightCCW0.back[2][0] != yellow {
		t.Errorf("Turning right counter clockwise did not work properly")
	}
	rightCCW1 := cube.actionRightLayerCCW(1)
	// fmt.Println(rightCCW1.showDetails())    
	if rightCCW1.top[0][1] != red || rightCCW1.front[2][1] != white || rightCCW1.back[2][1] != yellow {
		t.Errorf("Turning right counter clockwise did not work properly")
	}
	rightCCW2 := cube.actionRightLayerCCW(2)
	// fmt.Println(rightCCW2.showDetails())
	if rightCCW2.top[0][0] != red || rightCCW2.front[2][0] != white || rightCCW2.back[2][2] != yellow {
		t.Errorf("Turning right counter clockwise did not work properly")
	}	
}