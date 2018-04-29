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

package holdinglist

import (
	"sync"

	"github.com/jakeschurch/collections/internal/cache"
	"github.com/jakeschurch/collections/internal/linkedlist"
)

// Holdings is a collection that stores a cache and list.
type Holdings struct {
	sync.Mutex
	cache *cache.Cache
	list  []*linkedlist.List
}

func New() *Holdings {
	return &Holdings{
		cache: cache.New(),
		list:  make([]*linkedlist.List, 0),
	}
}

func (h *Holdings) Get(key string) (*linkedlist.List, error) {
	var list *linkedlist.List

	var i, err = h.cache.Get(key)
	if err != nil {
		return nil, err
	}

	h.Lock()
	list = h.list[i]
	h.Unlock()
	return list, nil
}

func (h *Holdings) Insert(key string) error {
	return nil
}
func (h *Holdings) Remove(key string) error {
	return nil
}
