// A search problem specification
// This is just a specification, not an implementation

package main

// State is the general interface for a search state
type State interface {

    // check, if final state is reached
    isFinal() bool

    // generate all follow-up states and return list of new states
    getChildren() []State

    // return an estimate to the final state -> no overestimation!
    getEstimate() int

    // return cost of current state
    getCost() int

    // return the previous state, nil, if it is the initial state
    getPrevious() State

    // return a hash value of the state
    getHash() uint64

    // return Cost + Estimate
    getEstimateOverall() int
}

