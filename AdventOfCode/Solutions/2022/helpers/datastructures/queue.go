package datastructures

import "sync"

type NodeQueue struct {
	Items []Vertex
	Lock  sync.RWMutex
}

type Vertex struct {
	X        int
	Y        int
	Distance int
}

func (s *NodeQueue) Enqueue(t Vertex) {
	s.Lock.Lock()
	if len(s.Items) == 0 {
		s.Items = append(s.Items, t)
		s.Lock.Unlock()
		return
	}
	var insertFlag bool
	for k, v := range s.Items {
		if t.Distance < v.Distance {
			if k > 0 {
				s.Items = append(s.Items[:k+1], s.Items[k:]...)
				s.Items[k] = t
				insertFlag = true
			} else {
				s.Items = append([]Vertex{t}, s.Items...)
				insertFlag = true
			}
		}
		if insertFlag {
			break
		}
	}
	if !insertFlag {
		s.Items = append(s.Items, t)
	}
	//s.items = append(s.items, t)
	s.Lock.Unlock()
}

// Dequeue removes a Node from the start of the queue
func (s *NodeQueue) Dequeue() *Vertex {
	s.Lock.Lock()
	item := s.Items[0]
	s.Items = s.Items[1:len(s.Items)]
	s.Lock.Unlock()
	return &item
}

// NewQ Creates New Queue
func (s *NodeQueue) NewQ() *NodeQueue {
	s.Lock.Lock()
	s.Items = []Vertex{}
	s.Lock.Unlock()
	return s
}

// IsEmpty returns true if the queue is empty
func (s *NodeQueue) IsEmpty() bool {
	s.Lock.RLock()
	defer s.Lock.RUnlock()
	return len(s.Items) == 0
}

// Size returns the number of Nodes in the queue
func (s *NodeQueue) Size() int {
	s.Lock.RLock()
	defer s.Lock.RUnlock()
	return len(s.Items)
}

// Queue Generic queue
type Queue[T any] struct {
	Items []T
}

func (s *Queue[T]) Push(t T) {
	s.Items = append(s.Items, t)
}

func (s *Queue[T]) Pop() *T {
	item := s.Items[0]
	s.Items = s.Items[1:len(s.Items)]
	return &item
}

// NewQ Creates New Queue
func (s *Queue[T]) NewQ() *Queue[T] {
	s.Items = []T{}
	return s
}

// IsEmpty returns true if the queue is empty
func (s *Queue[T]) IsEmpty() bool {
	return len(s.Items) == 0
}

// Size returns the number of Nodes in the queue
func (s *Queue[T]) Size() int {
	return len(s.Items)
}
