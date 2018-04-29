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

package linkedlist

import (
	"reflect"
	"testing"

	"github.com/jakeschurch/instruments"
)

func mockSummary() instruments.Summary {
	return instruments.Summary{
		N:      0,
		Volume: instruments.NewVolume(0),
	}
}
func TestNewLinkedList(t *testing.T) {
	mockedSummary := mockSummary()
	type args struct {
		summary instruments.Summary
	}
	tests := []struct {
		name string
		args args
		want *LinkedList
	}{
		{"base case", args{mockSummary()},
			&LinkedList{Summary: &mockedSummary,
				head: &HoldingNode{next: nil, prev: nil},
				tail: &HoldingNode{next: nil, prev: nil}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLinkedList(tt.args.summary); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLinkedList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_Push(t *testing.T) {
	mockedSummary := mockSummary()
	mockedNode := newNode(*mockHolding(), nil, nil)

	type fields struct {
		Summary *instruments.Summary
		head    *HoldingNode
		tail    *HoldingNode
	}
	type args struct {
		node *HoldingNode
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"base case",
			fields{
				Summary: &mockedSummary,
				head:    &HoldingNode{next: nil, prev: nil},
				tail:    &HoldingNode{next: nil, prev: nil}},
			args{mockedNode}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				Summary: tt.fields.Summary,
				head:    tt.fields.head,
				tail:    tt.fields.tail,
			}
			l.Push(tt.args.node)
			if tt.name == "base case" && ((l.head.next == nil) || (l.tail == nil)) {
				t.Error("HoldingNode.Push(): l.head.next should not be nil")
			}
		})
	}
}

func TestLinkedList_Pop(t *testing.T) {
	mockedSummary := mockSummary()
	mockedHead := &HoldingNode{next: nil, prev: nil}
	mockedTail := &HoldingNode{next: nil, prev: mockedHead}
	mockedHead.next = mockedTail

	newTail := &HoldingNode{next: nil, prev: mockedTail}

	type fields struct {
		Summary *instruments.Summary
		head    *HoldingNode
		tail    *HoldingNode
	}
	tests := []struct {
		name    string
		fields  fields
		want    *HoldingNode
		wantErr bool
	}{
		{"base case", fields{Summary: &mockedSummary,
			head: &HoldingNode{next: nil, prev: nil},
			tail: &HoldingNode{next: nil, prev: nil}},
			nil, true},
		{"tail not nil", fields{Summary: &mockedSummary,
			head: mockedHead,
			tail: mockedTail},
			mockedTail, false},
		{"3 elements", fields{Summary: &mockedSummary,
			head: mockedHead,
			tail: newTail},
			newTail, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				Summary: tt.fields.Summary,
				head:    tt.fields.head,
				tail:    tt.fields.tail,
			}
			if tt.name == "3 elements" {
				l.tail.prev = mockedTail
			}

			got, err := l.Pop()
			if (err != nil) != tt.wantErr {
				t.Errorf("LinkedList.Pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LinkedList.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}
