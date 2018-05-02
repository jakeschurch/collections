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
	"reflect"
	"testing"

	"github.com/jakeschurch/instruments"
)

func Test_newItems(t *testing.T) {
	tests := []struct {
		name string
		want *items
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newItems(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_items_get(t *testing.T) {
	type args struct {
		i uint32
	}
	tests := []struct {
		name    string
		n       *items
		args    args
		want    *list
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.n.get(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("items.get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("items.get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_items_remove(t *testing.T) {
	type args struct {
		i uint32
	}
	tests := []struct {
		name string
		n    *items
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

func Test_items_insert(t *testing.T) {
	type args struct {
		order *instruments.Order
		i     uint32
	}
	tests := []struct {
		name string
		n    *items
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.n.insert(tt.args.order, tt.args.i)
		})
	}
}

func Test_items_grow(t *testing.T) {
	type args struct {
		i uint32
	}
	tests := []struct {
		name string
		n    *items
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.n.grow(tt.args.i)
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *Orders
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

func TestOrders_Get(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		o       *Orders
		args    args
		want    uint32
		want1   *list
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.o.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Orders.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Orders.Get() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Orders.Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestOrders_Search(t *testing.T) {
	type args struct {
		order *instruments.Order
	}
	tests := []struct {
		name string
		o    *Orders
		args args
		want *node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.Search(tt.args.order); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Orders.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrders_Remove(t *testing.T) {
	type args struct {
		data *node
	}
	tests := []struct {
		name   string
		o      *Orders
		args   args
		wantOk bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOk := tt.o.Remove(tt.args.data); gotOk != tt.wantOk {
				t.Errorf("Orders.Remove() = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestOrders_Insert(t *testing.T) {
	type args struct {
		order *instruments.Order
	}
	tests := []struct {
		name string
		o    *Orders
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.o.Insert(tt.args.order)
		})
	}
}

func TestOrders_GetSlice(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		o       *Orders
		args    args
		want    []*instruments.Order
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.o.GetSlice(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Orders.GetSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Orders.GetSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
