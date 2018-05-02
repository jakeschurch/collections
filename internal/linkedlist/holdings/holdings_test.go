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
	"time"

	"github.com/jakeschurch/collections/internal/cache"
	"github.com/jakeschurch/instruments"
)

func mockItems() *items {
	return &items{
		data: make([]*list, 0),
		len:  0,
	}
}
func Test_newitems(t *testing.T) {
	tests := []struct {
		name string
		want *items
	}{
		{"base case", mockItems()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newItems(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newitems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func mockSummary() *instruments.Summary {
	newPrice := instruments.NewPrice(10.00)
	metric := &instruments.SummaryMetric{Price: newPrice, Date: time.Time{}}
	return &instruments.Summary{
		Name: "Google", N: 0, Volume: instruments.NewVolume(10.00), AvgAsk: &newPrice, AvgBid: &newPrice,
		MaxBid: metric, MinBid: metric, MaxAsk: metric, MinAsk: metric,
	}
}

func Test_items_remove(t *testing.T) {
	mockedItems := mockItems()
	NewList := NewList(*mockSummary())
	mockedItems.data = append(mockedItems.data, NewList)

	type args struct {
		i uint32
	}
	tests := []struct {
		name string
		n    *items
		args args
	}{
		{"base case", mockedItems, args{0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.n.remove(tt.args.i)
		})
	}
}

func Test_items_insert(t *testing.T) {
	type args struct {
		holding instruments.Holding
		i       uint32
	}
	tests := []struct {
		name string
		n    *items
		args args
	}{
		{"base case", mockItems(), args{*mockHolding(), 0}},
		{"need to grow slice", mockItems(), args{*mockHolding(), 2}},
		{"just push", mockItems(), args{*mockHolding(), 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "just push" {
				tt.n.data = make([]*list, 1)
				tt.n.data[0] = NewList(*mockSummary())
			}
			tt.n.insert(tt.args.holding, tt.args.i)
		})
	}
}

func mockHoldings() *Holdings {
	return &Holdings{
		cache: cache.New(),
		items: newItems(),
	}
}
func TestHoldings_Get(t *testing.T) {
	mockedHoldings := mockHoldings()
	mockedHoldings.Insert(*mockHolding())

	type args struct {
		key string
	}
	tests := []struct {
		name    string
		h       *Holdings
		args    args
		want    *list
		wantErr bool
	}{
		{"base case", mockedHoldings, args{"Google"}, mockedHoldings.data[0], false},
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
	mockedHoldings := mockHoldings()
	mockedHoldings.data = append(mockedHoldings.data, NewList(*mockSummary()))
	mockedHoldings.data[0].Push(NewNode(*mockHolding(), nil, nil))
	mockedHoldings.cache.Put(mockedHoldings.data[0].Name)

	type args struct {
		holding instruments.Holding
	}
	tests := []struct {
		name string
		h    *Holdings
		args args
	}{
		{"base case", mockHoldings(), args{*mockHolding()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.Insert(tt.args.holding)
		})
	}
}

func TestHoldings_Remove(t *testing.T) {
	mockedHoldings := mockHoldings()
	mockedHoldings.data = append(mockedHoldings.data, NewList(*mockSummary()))
	mockedHoldings.data[0].Push(NewNode(*mockHolding(), nil, nil))
	mockedHoldings.cache.Put(mockedHoldings.data[0].Name)

	type args struct {
		key string
	}
	tests := []struct {
		name    string
		h       *Holdings
		args    args
		wantErr bool
	}{
		{"base case", mockedHoldings, args{"Google"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.h.Remove(tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("Holdings.Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_newItems(t *testing.T) {
	tests := []struct {
		name string
		want *items
	}{
		{"base case", &items{
			data: make([]*list, 0),
			len:  0,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newItems(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newItems() = %v, want %v", got, tt.want)
			}
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
		{"base case", mockItems(), args{10}},
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
		want *Holdings
	}{
		{"base case", &Holdings{
			cache: cache.New(),
			items: newItems(),
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_items_get(t *testing.T) {
	mockedItems := mockItems()
	NewList := NewList(*mockSummary())
	mockedItems.data = append(mockedItems.data, NewList)

	type fields struct {
		data []*list
		len  uint32
	}
	type args struct {
		i uint32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *list
		wantErr bool
	}{
		{"base case", fields{mockedItems.data, mockedItems.len}, args{0}, NewList, false},
		{"err case", fields{mockedItems.data, mockedItems.len}, args{1}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &items{
				data: tt.fields.data,
				len:  tt.fields.len,
			}
			got, err := n.get(tt.args.i)
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

func TestHoldings_Update(t *testing.T) {
	type args struct {
		quote instruments.Quote
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
			if err := tt.h.Update(tt.args.quote); (err != nil) != tt.wantErr {
				t.Errorf("Holdings.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
