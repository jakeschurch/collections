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
	"reflect"
	"testing"

	"github.com/jakeschurch/collections/internal/linkedlist"
	"github.com/jakeschurch/instruments"
)

func Test_newNodelist(t *testing.T) {
	tests := []struct {
		name string
		want *nodelist
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newNodelist(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newNodelist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nodelist_get(t *testing.T) {
	type args struct {
		i uint32
	}
	tests := []struct {
		name string
		n    *nodelist
		args args
		want *linkedlist.List
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.get(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nodelist.get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nodelist_remove(t *testing.T) {
	type args struct {
		i uint32
	}
	tests := []struct {
		name string
		n    *nodelist
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.n.remove(tt.args.i)
		})
	}
}

func Test_nodelist_insert(t *testing.T) {
	type args struct {
		holding instruments.Holding
		i       uint32
	}
	tests := []struct {
		name string
		n    *nodelist
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.n.insert(tt.args.holding, tt.args.i)
		})
	}
}

func Test_nodelist_grow(t *testing.T) {
	tests := []struct {
		name string
		n    *nodelist
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.n.grow()
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *Holdings
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHoldings_Get(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		h       *Holdings
		args    args
		want    *linkedlist.List
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.h.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Holdings.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Holdings.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHoldings_Insert(t *testing.T) {
	type args struct {
		holding instruments.Holding
	}
	tests := []struct {
		name    string
		h       *Holdings
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.Insert(tt.args.holding); (err != nil) != tt.wantErr {
				t.Errorf("Holdings.Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHoldings_Remove(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		h       *Holdings
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.Remove(tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("Holdings.Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
