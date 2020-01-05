package main

import (
	"strconv"
	"fmt"
)

// CubeState is a search problem:
type CubeState struct {
	state           Cube
	previous        *CubeState
	cost            int
	estimateOverall int
	action          string
}

// check, if final state is reached
func (s CubeState) isFinal() bool {
	return isUniformColor(s.state.top) && isUniformColor(s.state.bottom) &&
		isUniformColor(s.state.front) && isUniformColor(s.state.back) &&
		isUniformColor(s.state.left) && isUniformColor(s.state.right)
}

// generate all follow-up CubeStates and return list of new States
func (s CubeState) getChildren() []State {
	var res []State
	for i := 0; i < len(s.state.top); i++ {
		res = append(res, CubeState{
			state:    s.state.actionFrontLayerCCW(i),
			previous: &s, cost: s.cost + 1, estimateOverall: -1,
			action: "Front CCW Layer " + strconv.Itoa(i)})
		res = append(res, CubeState{
			state:    s.state.actionFrontLayerCW(i),
			previous: &s, cost: s.cost + 1, estimateOverall: -1,
			action: "Front CW Layer " + strconv.Itoa(i)})
		res = append(res, CubeState{
			state:    s.state.actionTopLayerCCW(i),
			previous: &s, cost: s.cost + 1, estimateOverall: -1,
			action: "Top CCW Layer " + strconv.Itoa(i)})
		res = append(res, CubeState{
			state:    s.state.actionTopLayerCW(i),
			previous: &s, cost: s.cost + 1, estimateOverall: -1,
			action: "Top CW Layer " + strconv.Itoa(i)})
		res = append(res, CubeState{
			state:    s.state.actionRightLayerCCW(i),
			previous: &s, cost: s.cost + 1, estimateOverall: -1,
			action: "Right CCW Layer " + strconv.Itoa(i)})
		res = append(res, CubeState{
			state:    s.state.actionRightLayerCW(i),
			previous: &s, cost: s.cost + 1, estimateOverall: -1,
			action: "Right CW Layer " + strconv.Itoa(i)})
	}
	return res
}

// return an estimate to the final state -> no overestimation for admissible heuristic
func (s CubeState) getEstimate() int {
	var m []int
	m = append(m, heuristic(s.state.front))
	m = append(m, heuristic(s.state.back))
	m = append(m, heuristic(s.state.left))
	m = append(m, heuristic(s.state.right))
	m = append(m, heuristic(s.state.top))
	m = append(m, heuristic(s.state.bottom))
	return maxOfSlice(m)
}

// return cost of current state
func (s CubeState) getCost() int {
	return s.cost
}

// return the previous state, null, if it is the initial state
func (s CubeState) getPrevious() State {
	if s.previous == nil {
		return nil
	}
	return *s.previous
}

// return a string hash representation of the simple state
func (s CubeState) getHash() int {

// TODO: ggf. den getHash auf eine CubeState Var umlegen und bei
// generate Children jeweils den Hash mit generieren...

	fmt.Println("Calc Hash")
	// top := ident(s.state.top)
	// bot := ident(s.state.bottom)
	// fro := ident(s.state.front)
	// bac := ident(s.state.back)
	// lef := ident(s.state.left)
	// rig := ident(s.state.right)
	// return top + bot*intPow(8, 1) + fro*intPow(8, 2) +
    //         bac*intPow(8,3) + lef*intPow(8,4) + rig*intPow(8,5)

	return (ident(s.state.top)) +
		(ident(s.state.bottom))*8 +
		(ident(s.state.back))*16 +
		(ident(s.state.front))*24 +
		(ident(s.state.left))*32+
		(ident(s.state.right))*40

	// return fmt.Sprintf("%v|%v|%v|%v|%v|%v", 
	// 	s.state.front, s.state.back, s.state.top,
	//  	s.state.bottom, s.state.left, s.state.right)
}

// return Cost + Estimate, if estimate Overall is not yet calculated
func (s CubeState) getEstimateOverall() int {
	if s.estimateOverall == -1 {
		s.estimateOverall = s.cost + s.getEstimate()
	}
	return s.estimateOverall
}
