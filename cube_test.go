package main

import (
//	"fmt"
	"testing"
)

// Tests cube actions with some exemplary samples
func TestCube(t *testing.T) {
	// fmt.Println("Running 3x3x3 Testsuite")
	cube := Cube{
		Top:    [][]Color{{blue, orange, yellow}, {green, white, orange}, {orange, orange, yellow}},
		Bottom: [][]Color{{green, blue, white}, {orange, yellow, red}, {blue, blue, red}},
		Left:   [][]Color{{white, red, green}, {red, blue, blue}, {yellow, green, white}},
		Right:  [][]Color{{orange, white, blue}, {green, green, green}, {red, yellow, yellow}},
		Front:  [][]Color{{white, blue, green}, {red, orange, yellow}, {red, yellow, blue}},
		Back:   [][]Color{{red, yellow, orange}, {white, red, white}, {green, white, orange}}}

	TopCW := cube.turnTopCW()
	// fmt.Println(TopCW.showDetails())
	if TopCW.Top[0][0] != orange || TopCW.Front[0][2] != blue || TopCW.Left[0][1] != blue {
		t.Errorf("Turning Top clockwise did not work properly")
	}

	TopCCW := cube.turnTopCCW()
	// fmt.Println(TopCCW.showDetails())
	if TopCCW.Top[0][0] != yellow || TopCCW.Front[0][2] != green || TopCCW.Left[0][1] != yellow {
		t.Errorf("Turning Top clockwise did not work properly")
	}

	botCW := cube.turnBottomCW()
	// fmt.Println(botCW.showDetails())
	if botCW.Bottom[0][0] != blue || botCW.Front[2][1] != green || botCW.Right[2][0] != red {
		t.Errorf("Turning Bottom clockwise did not work properly")
	}

	botCCW := cube.turnBottomCCW()
	// fmt.Println(botCCW.showDetails())
	if botCCW.Bottom[0][0] != white || botCCW.Front[2][0] != red || botCCW.Right[2][0] != green {
		t.Errorf("Turning Bottom counter clockwise did not work properly")
	}

	FrontCW := cube.turnFrontCW()
	// fmt.Println(FrontCW.showDetails())
	if FrontCW.Top[2][2] != green || FrontCW.Front[0][0] != red || FrontCW.Right[1][0] != orange {
		t.Errorf("Turning Front clockwise did not work properly")
	}

	FrontCCW := cube.turnFrontCCW()
	// fmt.Println(FrontCCW.showDetails())
	if FrontCCW.Top[2][1] != green || FrontCCW.Bottom[0][0] != green || FrontCCW.Right[2][0] != green {
		t.Errorf("Turning Front counterclockwise did not work properly")
	}

	BackCW := cube.turnBackCW()
	// fmt.Println(BackCW.showDetails())
	if BackCW.Top[0][0] != blue || BackCW.Bottom[2][2] != yellow || BackCW.Left[1][0] != orange {
		t.Errorf("Turning Back clockwise did not work properly")
	}

	BackCCW := cube.turnBackCCW()
	// fmt.Println(BackCCW.showDetails())
	if BackCCW.Top[0][0] != yellow || BackCCW.Bottom[2][1] != green || BackCCW.Right[2][2] != yellow {
		t.Errorf("Turning Back counterclockwise did not work properly")
	}

	RightCW := cube.turnRightCW()
	// fmt.Println(RightCW.showDetails())
	if RightCW.Top[2][2] != blue || RightCW.Front[0][2] != white || RightCW.Back[0][0] != yellow {
		t.Errorf("Turning Right clockwise did not work properly")
	}

	RightCCW := cube.turnRightCCW()
	// fmt.Println(RightCCW.showDetails())
	if RightCCW.Top[2][2] != red || RightCCW.Front[0][2] != yellow || RightCCW.Back[0][0] != red {
		t.Errorf("Turning Right counter clockwise did not work properly")
	}

	LeftCW := cube.turnLeftCW()
	// fmt.Println(LeftCW.showDetails())
	if LeftCW.Top[0][0] != orange || LeftCW.Front[2][0] != orange || LeftCW.Back[1][2] != orange {
		t.Errorf("Turning Left clockwise did not work properly")
	}

	LeftCCW := cube.turnLeftCCW()
	// fmt.Println(LeftCCW.showDetails())
	if LeftCCW.Top[0][0] != white || LeftCCW.Front[2][0] != blue || LeftCCW.Back[1][2] != green {
		t.Errorf("Turning Left counter clockwise did not work properly")
	}
}

// todo: automate tests, no output
func TestCubeActions(t *testing.T) {
	// fmt.Println("Running 3x3x3 Testsuite")
	cube := Cube{
		Top:    [][]Color{{white, white, white}, {white, white, white}, {white, white, white}},
		Bottom: [][]Color{{yellow, yellow, yellow}, {yellow, yellow, yellow}, {yellow, yellow, yellow}},
		Front:  [][]Color{{orange, orange, orange}, {orange, orange, orange}, {orange, orange, orange}},
		Right:  [][]Color{{green, green, green}, {green, green, green}, {green, green, green}},
		Back:   [][]Color{{red, red, red}, {red, red, red}, {red, red, red}},
		Left:   [][]Color{{blue, blue, blue}, {blue, blue, blue}, {blue, blue, blue}}}

    if isUniformColor(cube.Front) == false {
		t.Errorf("Uniform color check 1 did not work correctly.")
	}
	if heuristic(cube.Front) != 0 {
		t.Errorf("Heuristic calculation of Front did not return expected 0")
	}
    TopCW0 := cube.actionTopLayerCW(0)
	// fmt.Println(TopCW0.showDetails())
	if TopCW0.Front[0][0] != green || TopCW0.Left[0][0] != orange || TopCW0.Back[0][2] != blue {
		t.Errorf("Turning Top clockwise did not work properly")
	}
	if isUniformColor(TopCW0.Front) {
		t.Errorf("Uniform color check 2 did not work correctly.")
	}
	if heuristic(TopCW0.Front) != 1 {
		t.Errorf("Heuristic calculation of Front did not return expected 1")
	}

    TopCW1 := cube.actionTopLayerCW(1)
	// fmt.Println(TopCW1.showDetails())
	if TopCW1.Front[1][0] != green || TopCW1.Left[1][1] != orange || TopCW1.Back[1][2] != blue {
		t.Errorf("Turning Top clockwise did not work properly")
	}
	TopCW2 := cube.actionTopLayerCW(2)
	// fmt.Println(TopCW2.showDetails())
	if TopCW2.Front[2][0] != green || TopCW2.Left[2][2] != orange || TopCW2.Back[2][2] != blue {
		t.Errorf("Turning Top clockwise did not work properly")
	}

    TopCCW0 := cube.actionTopLayerCCW(0)
	// fmt.Println(TopCCW0.showDetails())
	if TopCCW0.Front[0][0] != blue || TopCCW0.Left[0][0] != red || TopCCW0.Back[0][0] != green {
		t.Errorf("Turning Top counter clockwise did not work properly")
	}
    TopCCW1 := cube.actionTopLayerCCW(1)
	// fmt.Println(TopCCW1.showDetails())
	if TopCCW1.Front[1][0] != blue || TopCCW1.Left[1][0] != red || TopCCW1.Back[1][1] != green {
		t.Errorf("Turning Top counter clockwise did not work properly")
	}
	TopCCW2 := cube.actionTopLayerCCW(2)
	// fmt.Println(TopCCW2.showDetails())
	if TopCCW2.Front[2][0] != blue || TopCCW2.Left[2][0] != red || TopCCW2.Back[2][2] != green {
		t.Errorf("Turning Top counter clockwise did not work properly")
	}

    FrontCW0 := cube.actionFrontLayerCW(0)
	// fmt.Println(FrontCW0.showDetails())
	if FrontCW0.Top[2][2] != blue || FrontCW0.Left[1][2] != yellow || FrontCW0.Bottom[0][0] != green {
		t.Errorf("Turning Front clockwise did not work properly")
	}
    FrontCW1 := cube.actionFrontLayerCW(1)
    // fmt.Println(FrontCW1.showDetails())
	if FrontCW1.Top[1][1] != blue || FrontCW1.Left[1][1] != yellow || FrontCW1.Bottom[1][1] != green {
		t.Errorf("Turning Front clockwise did not work properly")
	}
    FrontCW2 := cube.actionFrontLayerCW(2)
	// fmt.Println(FrontCW2.showDetails())
	if FrontCW2.Top[0][0] != blue || FrontCW2.Left[2][0] != yellow || FrontCW2.Bottom[2][2] != green {
		t.Errorf("Turning Front clockwise did not work properly")
	}

    FrontCCW0 := cube.actionFrontLayerCCW(0)
	// fmt.Println(FrontCCW0.showDetails())
	if FrontCCW0.Top[2][0] != green || FrontCCW0.Left[0][2] != white || FrontCCW0.Bottom[0][2] != blue {
		t.Errorf("Turning Front counter clockwise did not work properly")
	}
    FrontCCW1 := cube.actionFrontLayerCCW(1)
	// fmt.Println(FrontCCW1.showDetails())
	if FrontCCW1.Top[1][1] != green || FrontCCW1.Left[0][1] != white || FrontCCW1.Bottom[1][1] != blue {
		t.Errorf("Turning Front counter clockwise did not work properly")
	}
	FrontCCW2 := cube.actionFrontLayerCCW(2)
	// fmt.Println(FrontCCW2.showDetails())
	if FrontCCW2.Top[0][0] != green || FrontCCW2.Left[0][0] != white || FrontCCW2.Bottom[2][2] != blue {
		t.Errorf("Turning Front counter clockwise did not work properly")
	}

    RightCW0 := cube.actionRightLayerCW(0)
	// fmt.Println(RightCW0.showDetails())
	if RightCW0.Top[2][2] != orange || RightCW0.Front[0][2] != yellow || RightCW0.Back[0][0] != white {
		t.Errorf("Turning Right clockwise did not work properly")
	}
	RightCW1 := cube.actionRightLayerCW(1)
	// fmt.Println(RightCW1.showDetails())
	if RightCW1.Top[2][1] != orange || RightCW1.Front[0][1] != yellow || RightCW1.Back[0][1] != white {
		t.Errorf("Turning Right clockwise did not work properly")
	}
	RightCW2 := cube.actionRightLayerCW(2)
	// fmt.Println(RightCW2.showDetails())
	if RightCW2.Top[2][0] != orange || RightCW2.Front[0][0] != yellow || RightCW2.Back[0][2] != white {
		t.Errorf("Turning Right clockwise did not work properly")
	}

    RightCCW0 := cube.actionRightLayerCCW(0)
	// fmt.Println(RightCCW0.showDetails())
	if RightCCW0.Top[0][2] != red || RightCCW0.Front[2][2] != white || RightCCW0.Back[2][0] != yellow {
		t.Errorf("Turning Right counter clockwise did not work properly")
	}
	RightCCW1 := cube.actionRightLayerCCW(1)
	// fmt.Println(RightCCW1.showDetails())    
	if RightCCW1.Top[0][1] != red || RightCCW1.Front[2][1] != white || RightCCW1.Back[2][1] != yellow {
		t.Errorf("Turning Right counter clockwise did not work properly")
	}
	RightCCW2 := cube.actionRightLayerCCW(2)
	// fmt.Println(RightCCW2.showDetails())
	if RightCCW2.Top[0][0] != red || RightCCW2.Front[2][0] != white || RightCCW2.Back[2][2] != yellow {
		t.Errorf("Turning Right counter clockwise did not work properly")
	}	
}