package main

import "fmt"

func main() {
	fmt.Println("Hello Rubiks Cube!")

	// var c Color = red
	// fmt.Println(c)
	// fmt.Println("Hello Color: " + c.String())

	cube := Cube{top: [][]Color{{red, white}, {white, white}},
				bottom: [][]Color{{yellow, yellow}, {green, yellow}},
				left: [][]Color{{orange, orange}, {orange, orange}},
				right: [][]Color{{red, red}, {red, red}},
				front: [][]Color{{green, green}, {green, green}},
				back: [][]Color{{blue, blue}, {blue, blue}}}
	fmt.Println("Cube looks like that: ", cube.showDetails())

	// newCube := cube.actTopCCW()
	fmt.Println("NewCube looks like that: ", cube.turnTopCW().showDetails())

	fmt.Println("Turning back: ", cube.turnTopCW().turnTopCCW().showDetails() )
	
	fmt.Println("Turning bottom ccw: ", cube.turnBottomCCW().showDetails())
}