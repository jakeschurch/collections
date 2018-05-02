package collections

import (
	"reflect"
	"testing"
	"time"

	"github.com/jakeschurch/instruments"
)

func mockOrder() *instruments.Order {
	return instruments.NewOrder("GOOGL", true, instruments.Market, instruments.NewPrice(10.00), instruments.NewVolume(10.00), time.Now())
}

func TestNewOrderBook(t *testing.T) {
	tests := []struct {
		name string
		want *OrderBook
	}{
		{"base case", NewOrderBook()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOrderBook(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOrderBook() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderBook_Insert(t *testing.T) {
	o := mockOrder()
	type args struct {
		o *instruments.Order
	}
	tests := []struct {
		name string
		ob   *OrderBook
		args args
	}{
		{"base case", NewOrderBook(), args{o}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ob.Insert(tt.args.o)
		})
	}
}

func TestOrderBook_GetBuys(t *testing.T) {
	// SETUP
	orderBook := NewOrderBook()
	order := mockOrder()
	// END SETUP

	type args struct {
		key string
	}
	tests := []struct {
		name    string
		ob      *OrderBook
		args    args
		want    []*instruments.Order
		wantErr bool
	}{
		{"sells nil", orderBook, args{order.Name}, nil, true},
		{"base case", orderBook, args{order.Name}, []*instruments.Order{order}, false},
	}
	for _, tt := range tests {

		if tt.name == "base case" {
			orderBook.Insert(order)
		}
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ob.GetBuys(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderBook.GetBuys() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderBook.GetBuys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderBook_GetSells(t *testing.T) {
	// SETUP
	orderBook := NewOrderBook()
	order := mockOrder()
	order.Buy = false
	// END SETUP

	type args struct {
		key string
	}
	tests := []struct {
		name    string
		ob      *OrderBook
		args    args
		want    []*instruments.Order
		wantErr bool
	}{
		{"sells nil", orderBook, args{order.Name}, nil, true},
		{"base case", orderBook, args{order.Name}, []*instruments.Order{order}, false},
	}
	for _, tt := range tests {

		if tt.name == "base case" {
			orderBook.Insert(order)
		}

		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ob.GetSells(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderBook.GetSells() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderBook.GetSells() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderBook_Remove(t *testing.T) {
	// SETUP
	ob := NewOrderBook()
	o := mockOrder()
	ob.Insert(o)
	// END SETUP

	type fields struct {
		orderBook *OrderBook
	}
	type args struct {
		o *instruments.Order
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantOk bool
	}{
		{"base case", fields{ob}, args{o}, true},
		{"err case", fields{ob}, args{o}, false},
	}
	for _, tt := range tests {
		if tt.name == "err case" {
			tt.args.o.Name = "A"
		}

		t.Run(tt.name, func(t *testing.T) {
			ob := tt.fields.orderBook
			if gotOk := ob.Remove(tt.args.o); gotOk != tt.wantOk {
				t.Errorf("OrderBook.Remove() = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
