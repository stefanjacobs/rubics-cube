package main

import (
	"fmt"
	"container/heap"
)

func aStarSolve(s0 State) State {
	openList := make(priorityQueue, 1)
	openList[0] = s0
	heap.Init(&openList)

	// Create an empty ClosedList
	closedList := map[string]bool{} // check for existence with _, ok := s[6]
	for ;; {
		// if the openlist is empty, there is no solution
		if openList.Len() == 0 {
			fmt.Println("No solution found!")
			return nil
		}
		currentState := heap.Pop(&openList).(State) // the current state is the first of the openlist
		if currentState.isFinal() { // check, if the current state is the solution -> Done!
			return currentState
		}
		closedList[currentState.getHash()] = true // set current state to closed list.
		children := currentState.getChildren() // generate all children of current state
		for _, child := range children {
			if closedList[child.getHash()] { // child is already on closed list -> nothing todo
				continue
			}
			pos, alreadyOnOpenList := openList.Contains(child)
			if alreadyOnOpenList != nil && alreadyOnOpenList.getEstimateOverall() <= child.getEstimate() {
				// child is already on openlist AND
				// the element on the openlist is less or equal expensive overall
				continue
			}
			if alreadyOnOpenList != nil { // child is cheaper than element on openlist -> replace and fix priorityqueue
				openList[pos] = child
				heap.Fix(&openList, pos)
			} else { // child is not yet on openlist, do it
				heap.Push(&openList, child)
			}
		}
		fmt.Printf("OL: %v, CL: %v\n", len(openList), len(closedList))
	}
}