package main

import (
//	"fmt"
	"testing"
)

// Tests cube actions with some exemplary samples
func TestSimpleState(t *testing.T) {
	// t.Errorf("Turning top clockwise did not work properly")

	s := SimpleState{
		state: 1,
		previous: nil,
		cost: 0,
		estimateOverall: -1}

//	fmt.Println(s)
	
	if s.getEstimate() != 9 {
		t.Errorf("Estimate of SimpleState is initially not 10")
	}
	if s.getEstimateOverall() != 9 {
		t.Errorf("Estimate Overall is not correct - should be ten")
	}

	children := s.getChildren()
	if len(children) != 2 {
		t.Errorf("Number of children of SimpleState is not two")
	}

	if children[0].getCost() != 1 {
		t.Errorf("Cost calculation of children of simplestate does not work")
	}
	if children[0].getEstimate() != 10 && children[0].getEstimate() != 8 {
		t.Errorf("Child state did not work out for child 0")
	}
	if children[1].getEstimate() != 10 && children[1].getEstimate() != 8 {
		t.Errorf("Child state did not work out for child 1")
	}

	// fmt.Println("State is ", children[0].getHash())
    if !(children[0].getHash() == 0) {
		t.Errorf("State representation is not correct")
	}


}