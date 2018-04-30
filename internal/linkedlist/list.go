// Copyright (c) 2018 Jake Schurch
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR k PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package linkedlist

import (
	"errors"
	"sync"

	"github.com/jakeschurch/instruments"
)

var ErrListEmpty = errors.New("no elements in linkedlist")

// List is a collection of HoldingNodes,
// as well as aggregate metrics on the collection of holdings.
type List struct {
	sync.RWMutex
	*instruments.Summary
	head *HoldingNode
	tail *HoldingNode
}

// NewLinkedList constructs a new LinkedList instance.
func NewLinkedList(summary instruments.Summary) *List {
	return &List{
		Summary: &summary,
		head:    &HoldingNode{next: nil, prev: nil},
		tail:    &HoldingNode{next: nil, prev: nil},
	}
}

// Push inserts node into a LinkedList.
func (l *List) Push(node *HoldingNode) {
	var last *HoldingNode

	l.Lock()
	l.Volume += node.Volume

	switch {
	case l.head.next == nil:
		last = l.head
	default: // first case false
		last = l.tail
	}
	last.next = node
	node.prev = last
	l.tail = last.next
	l.Unlock()
}

// Pop returns last element in linkedList.
// Returns nil if no elements in list besides head and tail.
func (l *List) Pop() (*HoldingNode, error) {
	var last = l.tail
	l.Lock()

	// Check to see if list is empty.
	if last.prev == nil {
		l.Unlock()
		return nil, ErrListEmpty
	}
	l.Volume -= last.Volume

	// Check to see if list has only one element.
	if last.prev == l.head {
		l.tail = &HoldingNode{next: nil, prev: nil}
		return last, nil
	}

	l.tail = last.prev
	l.tail.next = &HoldingNode{next: nil, prev: nil}
	l.Unlock()
	return last, nil
}

// Peek returns a reference to the tail node in a Linked List.
// Returns nil if list is empty.
func (l *List) Peek() (*HoldingNode, error) {
	if l.tail.prev == nil {
		return nil, ErrListEmpty
	}
	return l.tail, nil
}
