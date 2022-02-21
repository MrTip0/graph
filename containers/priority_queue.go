package containers

import "container/heap"

type Item struct {
	value    interface{}
	priority int
	index    int
}

type priorityQueue []*Item

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq priorityQueue) update(item *Item, priority int) {
	item.priority = priority
	heap.Fix(&pq, item.index)
}

type PriorityQueue struct {
	queue priorityQueue
}

func NewPriorityQueue() PriorityQueue {
	pr := make(priorityQueue, 0)
	qu := PriorityQueue{queue: pr}
	heap.Init(&pr)
	return qu
}

func (Q *PriorityQueue) Enqueue(val interface{}, priority int) {
	item := new(Item)
	item.priority = priority
	item.value = val
	heap.Push(&Q.queue, item)
}

func (pq PriorityQueue) Len() int { return pq.queue.Len() }

func (Q *PriorityQueue) Dequeue() interface{} {
	item := Q.queue.Pop().(*Item)
	return item.value
}

func (Q PriorityQueue) Contains(x interface{}) bool {
	for _, el := range Q.queue {
		if el.value == x {
			return true
		}
	}
	return false
}

func (Q *PriorityQueue) UpdatePriority(x interface{}, p int) {
	for _, el := range Q.queue {
		if el.value == x {
			Q.queue.update(el, p)
		}
	}
}
