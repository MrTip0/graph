package containers

import "container/list"

type Stack struct {
	v *list.List
}

func NewStack() Stack {
	return Stack{
		v: list.New(),
	}
}

func (s *Stack) Push(v interface{}) {
	s.v.PushBack(v)
}

func (s *Stack) Pop() interface{} {
	if s.v.Len() > 0 {
		ele := s.v.Back()
		val := ele.Value
		s.v.Remove(ele)
		return val
	}
	return nil
}

func (s *Stack) Len() int {
	return s.v.Len()
}

func (s Stack) Contains(val interface{}) bool {
	l := s.v
	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value == val {
			return true
		}
	}
	return false
}
