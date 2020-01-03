
package main
	
// CubeState is a search problem:
type CubeState struct {
	state Cube 
	previous *CubeState
	cost int
	estimateOverall int
	action string
}

// check, if final state is reached
func (s CubeState) isFinal() bool {
	return false
}

// generate all follow-up states and return list of new states
func (s CubeState) getChildren() []State {
	res := make([]State, 18)

	return res
}

// return an estimate to the final state -> no overestimation!
func (s CubeState) getEstimate() int {
	return 0
}

// return cost of current state
func (s CubeState) getCost() int {
	return s.cost
}

// return the previous state, null, if it is the initial state
func (s CubeState) getPrevious() *State {
	// return s.previous
	return nil
}

// return a string hash representation of the simple state
func (s CubeState) getHash() string {
	return ""
}

    // return Cost + Estimate, if estimate Overall is not yet calculated
func (s CubeState) getEstimateOverall() int {
	if s.estimateOverall == -1 {
		s.estimateOverall = s.cost + s.getEstimate()
	}
	return s.estimateOverall
}