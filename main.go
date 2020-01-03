package main

import (
	"fmt"
	"container/heap"
)

func main() {
	fmt.Println("Hello Rubik's Cube!")

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

	// a star algorithm as written here: https://de.wikipedia.org/wiki/A*-Algorithmus
	// Create Openlist with initial state
	s0 := SimpleState{
		state: 0,
		previous: nil,
		cost: 0,
		estimateOverall: -1,
	}
	openList := make(priorityQueue, 1)
	openList[0] = (*SimpleState)(&s0)
	heap.Init(&openList)

	// Create an empty ClosedList
	closedList := map[string]bool{} // check for existence with _, ok := s[6]
	for ;; {
		// if the openlist is empty, there is no solution
		if openList.Len() == 0 {
			fmt.Println("No solution found!")
			break
		}

		// the current state is the first of the openlist
		currentState := heap.Pop(&openList).(State)

		// check, if the current state is the solution -> Done!
		if currentState.isFinal() {
			fmt.Println("Found solution!")
			solution := currentState
			for ;; {
				if solution == nil {
					fmt.Println("Breaking...")
					break
				}
				fmt.Printf("state: %s\n", solution.getHash())
				solution = solution.getPrevious()
			}
			break
		}

		// set current state to closed list.
		closedList[currentState.getHash()] = true

		// generate all children of current state
		children := currentState.getChildren()
		for _, child := range children {
			if closedList[child.getHash()] {
				// child is already on closed list -> nothing todo
				continue
			}
			pos, alreadyOnOpenList := openList.Contains(child)
			if alreadyOnOpenList != nil && alreadyOnOpenList.getEstimateOverall() <= child.getEstimate() {
				// child is already on openlist AND
				// the element on the openlist is less or equal expensive overall
				continue
			}
			if alreadyOnOpenList != nil {
				// child is cheaper than element on openlist -> replace and fix priorityqueue
				openList[pos] = child
				heap.Fix(&openList, pos)
			} else {
				// child is not yet on openlist, do it
				heap.Push(&openList, child)
			}

		}
	}
}