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
// FITNESS FOR a PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package orders

import (
	"reflect"
	"testing"
	"time"

	"github.com/jakeschurch/instruments"
)

func mockOrder() *instruments.Order {
	return instruments.NewOrder("GOOGL", true, instruments.Market, instruments.NewPrice(10.00), instruments.NewVolume(10.00), time.Time{})
}

func Test_node_Next(t *testing.T) {

	firstNode := NewNode(mockOrder(), nil, nil)
	secondNode := NewNode(mockOrder(), firstNode, nil)

	type fields struct {
		Order *instruments.Order
		next  *node
		prev  *node
	}
	tests := []struct {
		name   string
		fields fields
		want   *node
	}{
		{"next is nil", fields{mockOrder(), nil, nil}, nil},
		{"next is not nil", fields{firstNode.Order, secondNode, nil}, secondNode},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := &node{
				Order: tt.fields.Order,
				next:  tt.fields.next,
				prev:  tt.fields.prev,
			}
			if got := node.Next(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("node.Next() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_node_Prev(t *testing.T) {
	firstNode := NewNode(mockOrder(), nil, nil)
	secondNode := NewNode(mockOrder(), firstNode, nil)

	type fields struct {
		Order *instruments.Order
		next  *node
		prev  *node
	}
	tests := []struct {
		name   string
		fields fields
		want   *node
	}{
		{"prev is nil", fields{mockOrder(), nil, nil}, nil},
		{"prev is not nil", fields{secondNode.Order, nil, firstNode}, firstNode},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := &node{
				Order: tt.fields.Order,
				next:  tt.fields.next,
				prev:  tt.fields.prev,
			}
			if got := node.Prev(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("node.Prev() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewNode(t *testing.T) {
	type args struct {
		o    *instruments.Order
		prev *node
		next *node
	}
	tests := []struct {
		name string
		args args
		want *node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNode(tt.args.o, tt.args.prev, tt.args.next); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
