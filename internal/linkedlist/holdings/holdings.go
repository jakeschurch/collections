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

package holdings

import (
	"errors"
	"sync"

	"github.com/jakeschurch/instruments"

	"github.com/jakeschurch/collections/internal/cache"
	"github.com/jakeschurch/collections/internal/linkedlist"
)

var ErrIndexRange = errors.New("index specified out of bounds")

// items is a container for multiple linked list instances.
type items struct {
	data []*linkedlist.List
	len  uint32
}

// newItems is a constructor for the items struct.
func newItems() *items {
	return &items{
		data: make([]*linkedlist.List, 0),
		len:  0,
	}
}

// get returns a linked list from data located at index i.
func (n *items) get(i uint32) (*linkedlist.List, error) {
	if i > n.len {
		return nil, ErrIndexRange
	}
	return n.data[i], nil
}

func (n *items) remove(i uint32) {
	n.data[i] = nil
}

func (n *items) insert(holding instruments.Holding, i uint32) {
	var node = linkedlist.NewNode(holding, nil, nil)

	if i >= n.len {
		n.grow(i)
	}
	if n.data[i] == nil {
		n.data[i] = linkedlist.NewLinkedList(*instruments.NewSummary(holding))
	}
	n.data[i].Push(node)
}

func (n *items) grow(i uint32) {
	for ; n.len <= i; n.len = (1 + n.len) * 2 {
	}
	n.data = append(n.data, make([]*linkedlist.List, n.len)...)
}

// ------------------------------------------------------------------

// Holdings is a collection that stores a cache and list.
type Holdings struct {
	sync.Mutex
	cache *cache.Cache
	*items
}

// New returns a new Holdings instance.
func New() *Holdings {
	return &Holdings{
		cache: cache.New(),
		items: newItems(),
	}
}

// Get returns a linkedlist.List associated with a key from Holdings.list.
// If none are associated with specific key, return nil.
func (h *Holdings) Get(key string) (*linkedlist.List, error) {
	var list *linkedlist.List

	var i, err = h.cache.Get(key)
	if err != nil {
		return nil, err
	}

	h.Lock()
	if list, err = h.items.get(i); err != nil {
		return list, err
	}
	h.Unlock()
	return list, nil
}

// Insert a holding into a Holdings's items linked list.
func (h *Holdings) Insert(holding instruments.Holding) (err error) {
	var i uint32
	h.Lock()
	if i, err = h.cache.Put(holding.Name); err != nil {
		h.Unlock()
		return err
	}
	h.items.insert(holding, i)
	h.Unlock()
	return nil
}

// Remove a holding into a Holdings's items linked list.
// If nothing can be removed, return error.
func (h *Holdings) Remove(key string) (err error) {
	var i uint32

	h.Lock()
	if i, err = h.cache.Remove(key); err != nil {
		h.Unlock()
		return err
	}
	h.items.remove(i)
	h.Unlock()
	return nil
}
