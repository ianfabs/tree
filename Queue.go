package main

import (
	"container/list"
)

// Queue is an implementation of Queue
type Queue struct {
	queue *list.List
}

func (q Queue) new() {
	q.queue = list.New()
}
func (q *Queue) clear() {
	q.queue.Init()
}
func (q Queue) isEmpty() bool {
	if q.queue.Len() == 0 {
		return true
	}
	return false
}
func (q Queue) first() *list.Element {
	return q.queue.Front()
}
func (q Queue) dequeue() *list.Element {
	first := q.queue.Front()
	q.queue.Remove(first)
	return first
}
func (q Queue) enqueue(el *list.Element) {
	q.queue.PushBack(el.Value)
}
