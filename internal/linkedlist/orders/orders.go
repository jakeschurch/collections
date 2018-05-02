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

package orders

import (
	"errors"
	"sync"

	"github.com/jakeschurch/instruments"

	"github.com/jakeschurch/collections/internal/cache"
)

var ErrIndexRange = errors.New("index specified out of bounds")

// items is a container for multiple linked list instances.
type items struct {
	data []*list
	len  uint32
}

// newItems is a constructor for the items struct.
func newItems() *items {
	return &items{
		data: make([]*list, 0),
		len:  0,
	}
}

// get returns a linked list from data located at index i.
func (n *items) get(i uint32) (*list, error) {
	if i > n.len {
		return nil, ErrIndexRange
	}
	return n.data[i], nil
}

func (n *items) remove(i uint32) {
	n.data[i] = nil
}

func (n *items) insert(order instruments.Order, i uint32) {
	var node = NewNode(order, nil, nil)

	if i >= n.len {
		n.grow(i)
	}
	if n.data[i] == nil {
		n.data[i] = NewList()
	}
	n.data[i].Push(node)
}

func (n *items) grow(i uint32) {
	for ; n.len <= i; n.len = (1 + n.len) * 2 {
	}
	n.data = append(n.data, make([]*list, n.len)...)
}

// ------------------------------------------------------------------

// Orders is a collection that stores a cache and list.
type Orders struct {
	sync.Mutex
	cache *cache.Cache
	*items
}

// New returns a new Orders instance.
func New() *Orders {
	return &Orders{
		cache: cache.New(),
		items: newItems(),
	}
}

// Get returns a list associated with a key from Orders.list.
// If none are associated with specific key, return nil.
func (o *Orders) Get(key string) (*list, error) {
	var list *list

	var i, err = o.cache.Get(key)
	if err != nil {
		return nil, err
	}

	o.Lock()
	if list, err = o.items.get(i); err != nil {
		return list, err
	}
	o.Unlock()
	return list, nil
}

// Insert a order into a Orders's items linked list.
func (o *Orders) Insert(order instruments.Order) (err error) {
	var i uint32
	o.Lock()
	if i, err = o.cache.Put(order.Name); err != nil {
		o.Unlock()
		return err
	}
	o.items.insert(order, i)
	o.Unlock()
	return nil
}

// Remove a order into a Orders's items linked list.
// If nothing can be removed, return error.
func (o *Orders) Remove(key string) (err error) {
	var i uint32

	o.Lock()
	if i, err = o.cache.Remove(key); err != nil {
		o.Unlock()
		return err
	}
	o.items.remove(i)
	o.Unlock()
	return nil
}
