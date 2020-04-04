package ds

import (
	"errors"
	"reflect"
)

// ListNode hold all nodes in a linkedlist
type ListNode struct {
	Index int
	Data  interface{}
	Next  *ListNode
}

// List every linkedlist have head and tail node
type List struct {
	Head *ListNode
	Tail *ListNode
}

// Insert insert new node at the tail of the list and return index of the inserted element with err if any
func (l *List) Insert(node interface{}) (int, error) {
	if !reflect.TypeOf(node).Comparable() {
		return 0, errors.New("can not insert un-comparable value")
	}
	temp := &ListNode{Data: node}
	if l.Head == nil {
		temp.Index = 1
		l.Head = temp
	} else {
		temp.Index = l.Tail.Index + 1
		l.Tail.Next = temp
	}
	l.Tail = temp
	return temp.Index, nil
}

// GetAt return data of list node at this index param
func (l *List) GetAt(index int) interface{} {
	for n := l.Head; n != nil; n = n.Next {
		if n.Index == index {
			return n.Data
		}
	}
	return -1
}

// Get return list node which value is equal to this node param
func (l *List) Get(node interface{}) *ListNode {
	for n := l.Head; n != nil; n = n.Next {
		if reflect.TypeOf(node) == reflect.TypeOf(n.Data) && n.Data == node {
			return n
		}
	}
	return nil
}

// LinkedList return new LinkedList
func LinkedList() *List { return new(List) }

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
