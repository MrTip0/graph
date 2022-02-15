package queue

import "container/list"

type Queue struct {
	v *list.List
}

func New() Queue {
	return Queue{
		v: list.New(),
	}
}

func (q *Queue) Enqueue(v interface{}) {
	q.v.PushBack(v)
}

func (q *Queue) Dequeue() interface{} {
	if q.v.Len() > 0 {
		ele := q.v.Front()
		val := ele.Value
		q.v.Remove(ele)
		return val
	}
	return nil
}

func (q *Queue) Len() int {
	return q.v.Len()
}

func (q Queue) Contains(val interface{}) bool {
	for	q.Len() > 0 {
		v := q.Dequeue()
		if v == val {
			return true
		}
	}
	return false
}