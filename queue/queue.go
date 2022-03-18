package queue

import (
	"errors"
	"fmt"
)

type Node struct {
	Val  interface{}
	Next *Node
	Prev *Node
}

type Queue struct {
	Head *Node
	Tail *Node
}

func MakeEmpty() *Queue {
	return &Queue{}
}

func Make(arr []interface{}) Queue {
	var queue Queue

	for _, e := range arr {
		queue.Push(e)
	}

	return queue
}

func (queue *Queue) Push(e interface{}) {
	newNode := &Node{
		Val: e,
	}

	if queue.Head == nil {
		queue.Head = newNode
		queue.Tail = queue.Head
	} else {
		queue.Head.Prev = newNode
		newNode.Next = queue.Head
		queue.Head = newNode
	}
}

func (queue *Queue) Poll() (interface{}, error) {
	if queue.Tail == nil {
		return nil, errors.New("can't poll an empty queue")
	}

	val := queue.Tail.Val
	queue.Tail = queue.Tail.Prev

	if queue.Tail != nil {
		queue.Tail.Next = nil
	} else {
		queue.Head = nil
	}

	return val, nil
}

func (queue *Queue) Format(f fmt.State, c rune) {
	curr := queue.Head
	for curr != nil {
		f.Write([]byte(fmt.Sprint(curr.Val)))
		curr = curr.Next
	}
}
