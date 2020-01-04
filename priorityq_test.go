package main

import (
	"testing"
	"container/heap"
)

// Tests cube actions with some exemplary samples
func TestPriorityQ(t *testing.T) {
	listStates := []*SimpleState{
		{state: 2, previous: nil, cost: 2, estimateOverall: -1},
		{state: -2, previous: nil, cost: 2, estimateOverall: -1},
		{state: 5, previous: nil, cost: 4, estimateOverall: -1},
		{state: 3, previous: nil, cost: 3, estimateOverall: -1},
	}
	pQueue := make(priorityQueue, len(listStates))
	for i, item := range listStates {
		pQueue[i] = (*SimpleState)(item)
	}

	heap.Init(&pQueue)

	newItem := SimpleState{state: 4, previous: nil, cost: 3, estimateOverall: -1}
	heap.Push(&pQueue, &newItem)

	// Print the order by Priority
	// for pQueue.Len() > 0 {
	// 	item := heap.Pop(&pQueue).(*SimpleState)
	// 	fmt.Printf("Name: %s with estimatedcost: %v\n", item.getHash(), item.getEstimateOverall())
	// }

	item := heap.Pop(&pQueue).(*SimpleState)
	if item.getHash() != 5 {
		t.Errorf("First item in priority was not the expected one...")
	}

	_, check := pQueue.Contains(newItem)
	if check == nil {
		t.Errorf("Queue should contain the searched item")
	}

	item = heap.Pop(&pQueue).(*SimpleState)
	if item.getHash() != 4 {
		t.Errorf("First item in priority was not the expected one...")
	}
	_, check = pQueue.Contains(newItem)
	if check != nil {
		t.Errorf("Queue should NOT contain the searched item")
	}
	item = heap.Pop(&pQueue).(*SimpleState)
	if item.getHash() != 3 {
		t.Errorf("First item in priority was not the expected one...")
	}
	item = heap.Pop(&pQueue).(*SimpleState)
	if item.getHash() != 2 {
		t.Errorf("First item in priority was not the expected one...")
	}
	item = heap.Pop(&pQueue).(*SimpleState)
	if item.getHash() != -2 {
		t.Errorf("First item in priority was not the expected one...")
	}
}