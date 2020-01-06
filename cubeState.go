package main

import (
	"strconv"
//	"github.com/mitchellh/hashstructure"
//	"fmt"
)

// CubeState is a search problem:
type CubeState struct {
	state           Cube
	previous        *CubeState
	cost            int
	estimateOverall int
	action          string
	Hash			int
}

// check, if final state is reached
func (s CubeState) isFinal() bool {
	return isUniformColor(s.state.Top) && isUniformColor(s.state.Bottom) &&
		isUniformColor(s.state.Front) && isUniformColor(s.state.Back) &&
		isUniformColor(s.state.Left) && isUniformColor(s.state.Right)
}

// generate all follow-up CubeStates and return list of new States
func (s CubeState) getChildren() []State {
	var res []State
	for i := 0; i < len(s.state.Top); i++ {
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
	m = append(m, heuristic(s.state.Front))
	m = append(m, heuristic(s.state.Back))
	m = append(m, heuristic(s.state.Left))
	m = append(m, heuristic(s.state.Right))
	m = append(m, heuristic(s.state.Top))
	m = append(m, heuristic(s.state.Bottom))
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
func (s CubeState) getHash() uint64 {
	// hash, _ := hashstructure.Hash(s.state, nil)
	// fmt.Printf("Value %v\n", hash)
	return s.state.hash
// TODO: ggf. den getHash auf eine CubeState Var umlegen und bei
// generate Children jeweils den Hash mit generieren...

	// fmt.Println("Calc Hash")
	// Top := ident(s.state.Top)
	// bot := ident(s.state.Bottom)
	// fro := ident(s.state.Front)
	// bac := ident(s.state.Back)
	// lef := ident(s.state.Left)
	// rig := ident(s.state.Right)
	// return Top + bot*intPow(8, 1) + fro*intPow(8, 2) +
    //         bac*intPow(8,3) + lef*intPow(8,4) + rig*intPow(8,5)

	// return (ident(s.state.Top)) +
	// 	(ident(s.state.Bottom))*8 +
	// 	(ident(s.state.Back))*16 +
	// 	(ident(s.state.Front))*24 +
	// 	(ident(s.state.Left))*32+
	// 	(ident(s.state.Right))*40

	// return fmt.Sprintf("%v|%v|%v|%v|%v|%v", 
	// 	s.state.Front, s.state.Back, s.state.Top,
	//  	s.state.Bottom, s.state.Left, s.state.Right)
}

// return Cost + Estimate, if estimate Overall is not yet calculated
func (s CubeState) getEstimateOverall() int {
	if s.estimateOverall == -1 {
		s.estimateOverall = s.cost + s.getEstimate()
	}
	return s.estimateOverall
}
