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
	"sync"

	"github.com/jakeschurch/instruments"

	"github.com/jakeschurch/collections/internal/cache"
	"github.com/jakeschurch/collections/internal/linkedlist"
)

type nodelist struct {
	data []*linkedlist.List
	len  uint32
}

func newNodelist() *nodelist {
	return &nodelist{
		data: make([]*linkedlist.List, 0),
		len:  0,
	}
}

func (n *nodelist) get(i uint32) *linkedlist.List {
	return n.data[i]
}

func (n *nodelist) remove(i uint32) {
	n.data[i] = nil
}

func (n *nodelist) insert(holding instruments.Holding, i uint32) {
	var node = linkedlist.NewNode(holding, nil, nil)
	if i >= n.len {
		n.grow()
	}
	n.data[i].Push(node)
}

func (n *nodelist) grow() {
	n.len = (1 + n.len) * 2
	n.data = append(make([]*linkedlist.List, n.len), n.data...)
}

// ------------------------------------------------------------------

// Holdings is a collection that stores a cache and list.
type Holdings struct {
	sync.Mutex
	cache *cache.Cache
	*nodelist
}

// New returns a new Holdings instance.
func New() *Holdings {
	return &Holdings{
		cache:    cache.New(),
		nodelist: newNodelist(),
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
	list = h.nodelist.get(i)
	h.Unlock()
	return list, nil
}

// TODO(Insert)
// Insert a new
func (h *Holdings) Insert(holding instruments.Holding) (err error) {
	var i uint32
	h.Lock()
	if i, err = h.cache.Put(holding.Name); err != nil {
		h.Unlock()
		return err
	}
	h.nodelist.insert(holding, i)
	return nil
}

func (h *Holdings) Remove(key string) (err error) {
	var i uint32

	h.Lock()
	if i, err = h.cache.Remove(key); err != nil {
		h.Unlock()
		return err
	}
	h.nodelist.remove(i)
	h.Unlock()
	return nil
}
