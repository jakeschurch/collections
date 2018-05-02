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

	"github.com/jakeschurch/instruments"
)

func mockHolding() *instruments.Holding {
	return &instruments.Holding{
		Name: "Google", Volume: instruments.NewVolume(20.00),
		Buy: instruments.TxMetric{
			Price: instruments.NewPrice(15.00), Date: time.Time{}},
	}
}
func TestHoldingNode_Next(t *testing.T) {

	firstNode := NewNode(*mockHolding(), nil, nil)
	secondNode := NewNode(*mockHolding(), firstNode, nil)

	type fields struct {
		Holding *instruments.Holding
		next    *node
		prev    *node
	}
	tests := []struct {
		name   string
		fields fields
		want   *node
	}{
		{"next is nil", fields{mockHolding(), nil, nil}, nil},
		{"next is not nil", fields{firstNode.Holding, secondNode, nil}, secondNode},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := &node{
				Holding: tt.fields.Holding,
				next:    tt.fields.next,
				prev:    tt.fields.prev,
			}
			if got := node.Next(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("node.Next() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHoldingNode_Prev(t *testing.T) {
	firstNode := NewNode(*mockHolding(), nil, nil)
	secondNode := NewNode(*mockHolding(), firstNode, nil)

	type fields struct {
		Holding *instruments.Holding
		next    *node
		prev    *node
	}
	tests := []struct {
		name   string
		fields fields
		want   *node
	}{
		{"prev is nil", fields{mockHolding(), nil, nil}, nil},
		{"prev is not nil", fields{secondNode.Holding, nil, firstNode}, firstNode},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := &node{
				Holding: tt.fields.Holding,
				next:    tt.fields.next,
				prev:    tt.fields.prev,
			}
			if got := node.Prev(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("node.Prev() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewNode(t *testing.T) {
	type args struct {
		h    instruments.Holding
		prev *node
		next *node
	}
	tests := []struct {
		name string
		args args
		want *node
	}{
		{"base case", args{*mockHolding(), nil, nil}, &node{mockHolding(), nil, nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNode(tt.args.h, tt.args.prev, tt.args.next); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_node_Next(t *testing.T) {
	tests := []struct {
		name string
		node *node
		want *node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.node.Next(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("node.Next() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_node_Prev(t *testing.T) {
	tests := []struct {
		name string
		node *node
		want *node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.node.Prev(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("node.Prev() = %v, want %v", got, tt.want)
			}
		})
	}
}
