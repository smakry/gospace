package queue

import (
	//"fmt"
)

// Push Pop Front Size

type Node struct {
	Data		interface{}
	NextNode	*Node
}

func NewNode() *Node{
	return &Node{
		Data:		nil,
		NextNode:	nil,
	}
}

type Queue struct {
	FrontNode	*Node
	TailNode	*Node
	size		int
}

func NewQueue() *Queue {
	return &Queue{
		FrontNode:	nil,
		TailNode:	nil,
		size:		0,
	}
}

func (q *Queue) Size() int {
	return q.size
}

func (q *Queue) Push(data interface{}) interface{} {
	newNode := &Node{data, nil}

	if q.Size() <= 0 {
		q.FrontNode = newNode
		q.TailNode = newNode
	} else {
		q.TailNode.NextNode = newNode
		q.TailNode = newNode
	}

	q.size++

	return newNode
}

func (q *Queue) Pop() interface{} {
	if q.Size() <= 0 {
		return nil
	}

	tmpData := q.FrontNode.Data
	if q.Size() == 1 {
		q.FrontNode = nil
		q.TailNode = nil
	} else {
		q.FrontNode = q.FrontNode.NextNode
	}

	q.size--
	return tmpData
}

func (q *Queue) Front() interface{} {
	if q.Size() <= 0 {
		return nil
	}

	return q.FrontNode.Data
}

func (q *Queue) Back() interface{} {
	if q.Size() <= 0 {	
		return nil
	}

	return q.TailNode.Data
}

func (q *Queue) DelNode(cur *Node) (*Node, bool) {
	if cur == nil {
		return nil, false
	}

	if cur.NextNode == nil {
		//last node
		for ptr := q.FrontNode; ; ptr = ptr.NextNode {
			if ptr.NextNode == cur {
				ptr.NextNode = nil

				if cur == q.TailNode {
					q.TailNode = ptr
				}

				break
			}
		}
	} else {
		cur.Data = cur.NextNode.Data
		cur.NextNode = cur.NextNode.NextNode
	}

	q.size--
	return cur, true
}



















