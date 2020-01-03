// A solver for problems that are characterized as described in problem.go

package main

import (
//	"container/heap"
//	"fmt"
)

type priorityQueue []State

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	// We want Pop to give us the lowest based on expiration number as the priority
	// The lower the expiry, the higher the priority
	lower := pq[i]
	upper := pq[j]
	// return lower.getCost()+lower.getEstimate() < upper.getCost()+upper.getEstimate()
	return lower.getEstimateOverall() < upper.getEstimateOverall()
}

// We just implement the pre-defined function in interface of heap.

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (pq *priorityQueue) Push(x interface{}) {
	item := x.(State)
	*pq = append(*pq, item)
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Contains(s State) (int, State) {
	sHash := s.getHash()
	for i, element := range *pq {
		if element.getHash() == sHash {
			return i, element
		}
	}
	return -1, nil
}