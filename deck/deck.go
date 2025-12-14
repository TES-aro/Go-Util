package deck

import (
)

// guess you gotta implement linked list yourself
// so yeah, it's more of a deck than a deck. sue me.

type deck[T any] struct{
	head *dnode[T]
	tail *dnode[T]
	size int
}

type dnode[T any] struct{
	value T
	prev *dnode[T]
	next *dnode[T]
}
// a pseudo node
func NewDeck[T any]() *deck[T]{
	return &deck[T]{nil,nil, 0}
}

func (q *deck[T])init(val T){
	node := &dnode[T]{val, nil, nil}
	q.head = node
	q.tail = node
}
func (q *deck[T]) Add(val T){
	if q.head == nil{
		q.init(val) 
		return
	}
	node := dnode[T]{val, q.tail, nil}
	q.tail.next = &node
	q.tail = &node
}

func (q *deck[T]) AddStart(val T){
	if q.head == nil{
		q.init(val) 
		return
	}
	node := &dnode[T]{val, nil, q.head}
	q.head.prev = node
	q.head = node
}

func (n *dnode[T]) Next() *dnode[T]{
	return n.next
}

func (q *deck[T]) Peek() T{
	return q.head.next.value
}

func (q *deck[T]) HasNext() bool{
	if q.head == nil{
		return false
	}
	return true
}

func (q *deck[T]) HasPrev() bool{
	if q.tail == nil{
		return false
	}
	return true
}

func (q *deck[T]) Next() T{
	val := &q.head.value
	if(q.head.next == nil){
		q.head = nil
		return *val
	}
	q.head = q.head.next
	q.head.prev = nil
	return *val
}

func (q *deck[T]) Prev() T{
	val := &q.tail.value
	if(q.tail.prev == nil){
		q.tail = nil
		return *val
	}
	q.tail = q.tail.prev
	q.head.next = nil
	return *val
}

