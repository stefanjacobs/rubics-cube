package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

func main() {
	fmt.Println("Hello Rubik's Cube!")

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	// var c Color = red
	// fmt.Println(c)
	// fmt.Println("Hello Color: " + c.String())

	// cube := Cube{top: [][]Color{{red, white}, {white, white}},
	// 			bottom: [][]Color{{yellow, yellow}, {green, yellow}},
	// 			left: [][]Color{{orange, orange}, {orange, orange}},
	// 			right: [][]Color{{red, red}, {red, red}},
	// 			front: [][]Color{{green, green}, {green, green}},
	// 			back: [][]Color{{blue, blue}, {blue, blue}}}
	// fmt.Println("Cube looks like that: ", cube.showDetails())

	// // newCube := cube.actTopCCW()
	// fmt.Println("NewCube looks like that: ", cube.turnTopCW().showDetails())

	// fmt.Println("Turning back: ", cube.turnTopCW().turnTopCCW().showDetails() )

	// fmt.Println("Turning bottom ccw: ", cube.turnBottomCCW().showDetails())

	// WORKING SIMPLESTATE
	// s0 := SimpleState{
	// 	state: 0,
	// 	previous: nil,
	// 	cost: 0,
	// 	estimateOverall: -1,
	// }

	// solution := aStarSolve(s0)
	// if solution == nil {
	// 	fmt.Println("No solution found")
	// 	return
	// }
	// for ;; {
	// 	if solution == nil {
	// 		fmt.Println("Breaking...")
	// 		break
	// 	}
	// 	fmt.Printf("state: %s\n", solution.getHash())
	// 	solution = solution.getPrevious()
	// }

	// RUBIKs CUBE
	cube := Cube{
		top:    [][]Color{{white, white, white}, {white, white, white}, {white, white, white}},
		bottom: [][]Color{{yellow, yellow, yellow}, {yellow, yellow, yellow}, {yellow, yellow, yellow}},
		front:  [][]Color{{orange, orange, orange}, {orange, orange, orange}, {orange, orange, orange}},
		right:  [][]Color{{green, green, green}, {green, green, green}, {green, green, green}},
		back:   [][]Color{{red, red, red}, {red, red, red}, {red, red, red}},
		left:   [][]Color{{blue, blue, blue}, {blue, blue, blue}, {blue, blue, blue}}}
	// cube := Cube{
	// 	top:    [][]Color{{blue, orange, yellow}, {green, white, orange}, {orange, orange, yellow}},
	// 	bottom: [][]Color{{green, blue, white}, {orange, yellow, red}, {blue, blue, red}},
	// 	left:   [][]Color{{white, red, green}, {red, blue, blue}, {yellow, green, white}},
	// 	right:  [][]Color{{orange, white, blue}, {green, green, green}, {red, yellow, yellow}},
	// 	front:  [][]Color{{white, blue, green}, {red, orange, yellow}, {red, yellow, blue}},
	// 	back:   [][]Color{{red, yellow, orange}, {white, red, white}, {green, white, orange}}}

	s0 := CubeState{
		// state: cube.actionTopLayerCW(2).actionRightLayerCCW(1).actionFrontLayerCCW(1).actionTopLayerCCW(0).actionRightLayerCCW(2).actionTopLayerCW(1).actionRightLayerCCW(0),
		state: cube.actionFrontLayerCCW(1).actionTopLayerCCW(0).actionRightLayerCCW(2).actionTopLayerCW(1).actionRightLayerCCW(0),
		// state: cube,
		previous:        nil,
		cost:            0,
		estimateOverall: -1,
		action:          "none",
	}

	solution := aStarSolve(s0)
	if solution == nil {
		fmt.Println("No solution found")
		return
	}
	for {
		if solution == nil {
			fmt.Println("Breaking out of solution loop...")
			break
		}
		// fmt.Printf("state: %s, Action: %s\n", solution.getHash(), solution.(CubeState).action)
		fmt.Printf("state: %v\n", solution.getHash())
		solution = solution.getPrevious()
	}

	fmt.Println("Done!")

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close()
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
}
