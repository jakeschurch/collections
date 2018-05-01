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

package ordernode

import (
	"github.com/jakeschurch/instruments"
)

// OrderNode is an element associated within a LinkedList.
type OrderNode struct {
	Order      *instruments.Order
	next, prev *OrderNode
}

// NewNode returns a new reference to an OrderNode.
func NewNode(o *instruments.Order, prev, next *OrderNode) *OrderNode {
	var node = &OrderNode{
		Order: o, next: next, prev: prev,
	}
	if x := node.prev; x != nil {
		x.next = node
	}
	return node
}

// Next returns a reference to a HoldingNode's next Holdingnode pointer.
func (node *OrderNode) Next() *OrderNode {
	return node.next
}

// Prev returns a reference to a HoldingNode's prev Holdingnode pointer.
func (node *OrderNode) Prev() *OrderNode {
	return node.prev
}
