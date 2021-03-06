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
	"reflect"
	"testing"
)

func Test_keyIndex_insertAndSort(t *testing.T) {
	type args struct {
		n uint32
	}
	tests := []struct {
		name    string
		k       keyIndex
		args    args
		want    keyIndex
		wantErr bool
	}{
		{"empty slice", make(keyIndex, 0), args{1}, keyIndex{1}, false},
		{"err: key exists in slice", keyIndex{1}, args{1}, keyIndex{1}, true},
		{"check sort", keyIndex{4, 3, 2, 1}, args{5}, keyIndex{1, 2, 3, 4, 5}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.k.insertAndSort(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("keyIndex.insertAndSort() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("keyIndex.insertAndSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCache_Put(t *testing.T) {
	type fields struct {
		items     map[string]uint32
		openSlots keyIndex
		n         uint32
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"base case",
			fields{openSlots: make(keyIndex, 0), items: make(map[string]uint32), n: 0},
			args{"GOOGL"}, false},
		{"key already exists in map[string]uint32",
			fields{openSlots: make(keyIndex, 0), items: map[string]uint32{"GOOGL": 0}, n: 1}, args{"GOOGL"}, true},
		{"take index from openSlots",
			fields{openSlots: append(make(keyIndex, 1), 1), items: map[string]uint32{"GOOGL": 0}, n: 2}, args{"GOOG"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Cache{
				items:     tt.fields.items,
				openSlots: tt.fields.openSlots,
				n:         tt.fields.n,
			}
			if _, err := l.Put(tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("Cache.Put() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCache_Get(t *testing.T) {
	type fields struct {
		items     map[string]uint32
		openSlots keyIndex
		n         uint32
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    uint32
		wantErr bool
	}{
		{"base case",
			fields{openSlots: make(keyIndex, 0), items: map[string]uint32{"GOOGL": 0}, n: 0},
			args{"GOOGL"}, 0, false},
		{"key not found",
			fields{openSlots: make(keyIndex, 0), items: make(map[string]uint32), n: 0},
			args{"GOOGL"}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Cache{
				items:     tt.fields.items,
				openSlots: tt.fields.openSlots,
				n:         tt.fields.n,
			}
			got, err := l.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Cache.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Cache.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCache_Remove(t *testing.T) {
	type fields struct {
		items     map[string]uint32
		openSlots keyIndex
		n         uint32
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"base case",
			fields{openSlots: make(keyIndex, 0), items: map[string]uint32{"GOOGL": 0}, n: 0},
			args{"GOOGL"}, false},
		{"key not found",
			fields{openSlots: make(keyIndex, 0), items: make(map[string]uint32), n: 0},
			args{"GOOGL"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Cache{
				items:     tt.fields.items,
				openSlots: tt.fields.openSlots,
				n:         tt.fields.n,
			}
			if _, err := l.Remove(tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("Cache.Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func BenchmarkCache_Put(b *testing.B) {
	var c = New()
	for i := 0; i < b.N; i++ {
		c.Put("GOOGL")
	}
}

func BenchmarkCache_Get(b *testing.B) {
	var c = New()
	c.Put("GOOGL")
	for i := 0; i < b.N; i++ {
		c.Get("GOOGL")
	}
}

func BenchmarkCache_Remove(b *testing.B) {
	var c = New()
	for i := 0; i < b.N; i++ {
		c.Remove("GOOGL")
	}
}
