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

package collections

import (
	"reflect"
	"testing"
	"time"

	"github.com/jakeschurch/instruments"
)

func mockHolding() instruments.Holding {
	return instruments.Holding{
		Name: "GOOGL", Volume: instruments.NewVolume(15.00),
		Buy: instruments.TxMetric{Price: instruments.NewPrice(10.00), Date: time.Time{}},
	}
}
func mockQuote() instruments.Quote {
	quoted := instruments.NewQuotedMetric(10, 11)
	return instruments.Quote{
		Name: "GOOGL", Bid: quoted, Ask: quoted, Timestamp: time.Time{},
	}
}

func TestNewPortfolio(t *testing.T) {
	tests := []struct {
		name string
		want *Portfolio
	}{
		{"base case", NewPortfolio()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPortfolio(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPortfolio() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPortfolio_Remove(t *testing.T) {
	// Setup Code
	port := NewPortfolio()
	port.Insert(mockHolding())
	// END Setup Code

	type args struct {
		key string
	}
	tests := []struct {
		name    string
		p       *Portfolio
		args    args
		wantErr bool
	}{
		{"base case", port, args{"GOOGL"}, false},
		{"err case", port, args{"A"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.Remove(tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("Portfolio.Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPortfolio_Insert(t *testing.T) {
	type args struct {
		holding instruments.Holding
	}
	tests := []struct {
		name string
		p    *Portfolio
		args args
	}{
		{"base case", NewPortfolio(), args{mockHolding()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.Insert(tt.args.holding)
		})
	}
}

func TestPortfolio_Update(t *testing.T) {
	port := NewPortfolio()
	port.Insert(mockHolding())

	type args struct {
		quote instruments.Quote
	}
	tests := []struct {
		name    string
		p       *Portfolio
		args    args
		wantErr bool
	}{
		{"base case", port, args{mockQuote()}, false},
		{"err not found", port,
			args{
				instruments.Quote{Name: "A", Bid: instruments.NewQuotedMetric(10, 11), Ask: instruments.NewQuotedMetric(12, 10), Timestamp: time.Time{}}},
			true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.Update(tt.args.quote); (err != nil) != tt.wantErr {
				t.Errorf("Portfolio.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
