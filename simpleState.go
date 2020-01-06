package main

import "math"
// import "strconv"

// SimpleState is a trivial search problem:
// search on one-dimensional grid for value 10
// actions are +1 and -1
type SimpleState struct {
	state int
	previous *SimpleState
	cost int
	estimateOverall int
}

const goal = 10 

// check, if final state is reached
func (s SimpleState) isFinal() bool {
	if s.state == goal {
		return true
	}
	return false
}

// generate all follow-up states and return list of new states
func (s SimpleState) getChildren() []State {
	var res []State
	if s.state > -20 {	
		low := SimpleState{
			state: s.state-1,
			previous: &s,
			cost: s.cost+1,
			estimateOverall: -1,
		}
		res = append(res, low)
	}
	if s.state < 20 {
		high := SimpleState{
			state: s.state+1,
			previous: &s,
			cost: s.cost+1,
			estimateOverall: -1,
		}
		res = append(res, high)
	}
	return res
}

// return an estimate to the final state -> no overestimation!
func (s SimpleState) getEstimate() int {
	return int(math.Abs(float64(goal)-float64(s.state)))
}

// return cost of current state
func (s SimpleState) getCost() int {
	return s.cost
}

// return the previous state, nil, if it is the initial state
func (s SimpleState) getPrevious() State {
	if s.previous == nil {
		return nil
	}
	return s.previous
}

// return a string hash representation of the simple state
func (s SimpleState) getHash() uint64 {
	return uint64(s.state)
}

    // return Cost + Estimate, if estimate Overall is not yet calculated
func (s SimpleState) getEstimateOverall() int {
	if s.estimateOverall == -1 {
		s.estimateOverall = s.cost + s.getEstimate()
	}
	return s.estimateOverall
}