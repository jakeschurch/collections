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

package cache

import (
	"sync"
	"sync/atomic"

	"github.com/pkg/errors"
)

var (
	ErrKeyExists   = errors.New("key already exists")
	ErrKeyNotFound = errors.New("key not found")
)

type keyIndex []uint32

// insertAndSort first checks to see if n exists in the keyIndex.
// If it does, an error will be returned. If not,  k will append n to it,
// and perform insertion sort on the slice k.
func (k keyIndex) insertAndSort(n uint32) (keyIndex, error) {
	for index := range k {
		if k[index] == n {
			return k, errors.Wrap(ErrKeyExists, "no insert of n necessary")
		}
	}
	k = append(k, n)

	for j := 1; j < len(k); j++ {
		key := k[j]
		i := j - 1

		for i >= 0 && k[i] > key { // shift
			k[i+1] = k[i]
			i--
		}
		k[i+1] = key
	}
	return k, nil
}

type keyMap map[string]uint32

type LookupCache struct {
	sync.RWMutex
	items     keyMap
	openSlots keyIndex
	n         uint32
}

// NewLookupCache creates a new LookupCache.
func NewLookupCache() *LookupCache {
	return &LookupCache{
		items:     make(keyMap),
		openSlots: make(keyIndex, 0),
		n:         0,
	}
}

// Put records the value of a key-index pair in a LookupCache's openSlots map.
func (l *LookupCache) Put(key string) error {
	var i uint32

	l.Lock()
	// Check to see if we can grab an index from an open slot.
	switch len(l.openSlots) {
	case 0:
		i = atomic.LoadUint32(&l.n)
		atomic.AddUint32(&l.n, 1)
	default:
		i, l.openSlots = l.openSlots[0], l.openSlots[1:]
	}
	// Check to see if key already exisits, if so throw error.
	if _, ok := l.items[key]; ok {
		l.Unlock()
		return errors.Wrap(ErrKeyExists, "in LookupCache's KeyMap")
	}
	l.items[key] = i

	return nil
}

// Get returns the value associated with key value.
// If key does not exist in map, 0 & error are returned.
func (l *LookupCache) Get(key string) (uint32, error) {
	var i uint32
	var ok bool

	l.Lock()
	if i, ok = l.items[key]; !ok {
		l.Unlock()
		return i, errors.Wrap(ErrKeyNotFound, "in LookupCache's KeyMap")
	}
	l.Unlock()
	return i, nil
}

func (l *LookupCache) Remove(key string) error {
	var i uint32
	var ok bool
	var err error

	l.Lock()
	if i, ok = l.items[key]; !ok {
		l.Unlock()
		return ErrKeyNotFound
	}
	delete(l.items, key)
	if l.openSlots, err = l.openSlots.insertAndSort(i); (err != nil) && (err != ErrKeyExists) {
		return err
	}
	l.Unlock()
	return nil
}
