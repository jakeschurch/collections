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

	"github.com/jakeschurch/instruments"
)

func TestNewList(t *testing.T) {
	type args struct {
		summary instruments.Summary
	}
	tests := []struct {
		name string
		args args
		want *list
	}{
		{"base case", args{*mockSummary()},
			&list{Summary: mockSummary(),
				head: &node{Holding: nil, next: nil, prev: nil},
				tail: &node{Holding: nil, next: nil, prev: nil}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewList(tt.args.summary); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLinkedList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_list_Push(t *testing.T) {
	mockedNode := NewNode(*mockHolding(), nil, nil)
	mockedHead := &node{Holding: nil, next: nil, prev: nil}
	mockedTail := &node{Holding: mockHolding(), next: nil, prev: mockedHead}
	mockedHead.next = mockedTail

	type fields struct {
		Summary *instruments.Summary
		head    *node
		tail    *node
	}
	type args struct {
		node *node
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"base case",
			fields{
				Summary: mockSummary(),
				head:    &node{Holding: nil, next: nil, prev: nil},
				tail:    &node{Holding: nil, next: nil, prev: nil}},
			args{mockedNode}},
		{"tail not nil case",
			fields{
				Summary: mockSummary(),
				head:    mockedHead,
				tail:    mockedTail},
			args{mockedNode}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &list{
				Summary: tt.fields.Summary,
				head:    tt.fields.head,
				tail:    tt.fields.tail,
			}
			l.Push(tt.args.node)
			if tt.name == "base case" && ((l.head.next == nil) || (l.tail == nil)) {
				t.Error("node.Push(): l.head.next should not be nil")
			}
		})
	}
}

func Test_list_Pop(t *testing.T) {
	mockedHead := &node{Holding: nil, next: nil, prev: nil}
	mockedTail := &node{Holding: mockHolding(), next: nil, prev: mockedHead}
	mockedHead.next = mockedTail

	newTail := &node{Holding: mockHolding(), next: nil, prev: mockedTail}

	type fields struct {
		Summary *instruments.Summary
		head    *node
		tail    *node
	}
	tests := []struct {
		name    string
		fields  fields
		want    *node
		wantErr bool
	}{
		{"base case", fields{Summary: mockSummary(),
			head: &node{Holding: nil, next: nil, prev: nil},
			tail: &node{Holding: nil, next: nil, prev: nil}},
			nil, true},
		{"tail not nil", fields{Summary: mockSummary(),
			head: mockedHead,
			tail: mockedTail},
			mockedTail, false},
		{"3 elements", fields{Summary: mockSummary(),
			head: mockedHead,
			tail: newTail},
			newTail, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &list{
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

func Test_list_Peek(t *testing.T) {
	mockedHead := &node{Holding: nil, next: nil, prev: nil}
	mockedTail := &node{Holding: mockHolding(), next: nil, prev: mockedHead}
	mockedHead.next = mockedTail

	newTail := &node{Holding: mockHolding(), next: nil, prev: mockedTail}

	type fields struct {
		Summary *instruments.Summary
		head    *node
		tail    *node
	}
	tests := []struct {
		name    string
		fields  fields
		want    *node
		wantErr bool
	}{
		{"base case", fields{Summary: mockSummary(),
			head: &node{Holding: nil, next: nil, prev: nil},
			tail: &node{Holding: nil, next: nil, prev: nil}},
			nil, true},
		{"tail not nil", fields{Summary: mockSummary(),
			head: mockedHead,
			tail: mockedTail},
			mockedTail, false},
		{"3 elements", fields{Summary: mockSummary(),
			head: mockedHead,
			tail: newTail},
			newTail, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &list{
				Summary: tt.fields.Summary,
				head:    tt.fields.head,
				tail:    tt.fields.tail,
			}
			if tt.name == "3 elements" {
				l.tail.prev = mockedTail
			}

			got, err := l.Peek()
			if (err != nil) != tt.wantErr {
				t.Errorf("LinkedList.Peek() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LinkedList.Peek() = %v, want %v", got, tt.want)
			}
		})
	}
}
