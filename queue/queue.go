package queue

import (
)

// guess you gotta implement linked list yourself
// a nice fifo queue. AddStart does turn it into FILO tho

type queue[T any] struct{
	head *node[T]
	tail *node[T]
}

type node[T any] struct{
	value T
	next *node[T]
}

func NewQueue[T any](...T) *queue[T]{
	return &queue[T]{nil,nil}
}

func (q *queue[T])init(val T){
	node := &node[T]{val, nil}
	q.head = node
	q.tail = node
}
func (q *queue[T]) Add(val T){
	if q.head == nil{
		q.init(val) 
		return
	}
	node := node[T]{val, q.tail}
	q.tail.next = &node
	q.tail = &node
}

func (q *queue[T]) AddStart(val T){
	if q.head == nil{
		q.init(val) 
		return
	}
	node := &node[T]{val, q.head.next}
	q.head = node
}

func (n *node[T]) Next() *node[T]{
	return n.next
}

func (q *queue[T]) Peek() T{
	return q.head.next.value
}

func (q *queue[T]) HasNext() bool{
	if q.head == nil{
		return false
	}
	return true
}

func (q *queue[T]) Next() T{
	val := &q.head.value
	if(q.head.next == nil){
		q.head = nil
		return *val
	}
	q.head = q.head.next
	return *val
}
