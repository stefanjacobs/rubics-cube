
package main

import "math"
import "strconv"

// SimpleState is a trivial search problem:
// search on one-dimensional grid for value 10
// actions are +1 and -1
type SimpleState struct {
	state int
	previous *SimpleState
	cost int
	estimateOverall int
}

// check, if final state is reached
func (s SimpleState) isFinal() bool {
	if s.state == 10 {
		return true
	}
	return false
}

// generate all follow-up states and return list of new states
func (s SimpleState) getChildren() []State {
	low := SimpleState{
		state: s.state-1,
		previous: &s,
		cost: s.cost+1,
	}
	low.estimateOverall = low.getCost() + low.getEstimate()
	high := SimpleState{
		state: s.state+1,
		previous: &s,
		cost: s.cost+1,
	}
	high.estimateOverall = high.getCost() + high.getEstimate()
	res := make([]State, 2)
	res[0]=low
	res[1]=high
	return res
}

// return an estimate to the final state -> no overestimation!
func (s SimpleState) getEstimate() int {
	return int(math.Abs(10.0-float64(s.state)))
}

// return cost of current state
func (s SimpleState) getCost() int {
	return s.cost
}

// return the previous state, null, if it is the initial state
func (s SimpleState) getPrevious() State {
	return s.previous
}

// return a string hash representation of the simple state
func (s SimpleState) getHash() string {
	return strconv.Itoa(s.state)
}

    // return Cost + Estimate
func (s SimpleState) getEstimateOverall() int {
	if s.estimateOverall == -1 {
		s.estimateOverall = s.cost + s.getEstimate()
	}
	return s.estimateOverall
}