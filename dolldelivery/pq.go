package dolldelivery

//this is an example that demonstrates a Priority queue built using the heap interface.
//this is from the golang documentation here: https://golang.org/pkg/container/heap/

import (
	"container/heap"
)

// An Item is something we manage in a Priority queue.
type Item struct {
	Value    string // The value of the item; arbitrary.
	Priority int    // The Priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	Index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the lowest priority = lowest cost
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

//Push adds a new item on to the queue
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

//Pop removes the top most item off of the queue
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.Index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

//Update modifies the Priority and value of an Item in the queue.
func (pq *PriorityQueue) Update(item *Item, value string, Priority int) {
	item.Value = value
	item.Priority = Priority
	heap.Fix(pq, item.Index)
}
