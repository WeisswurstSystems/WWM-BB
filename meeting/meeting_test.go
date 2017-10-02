package meeting

import (
	"reflect"
	"testing"
)

func TestMeeting_FindOrderByCustomer(t *testing.T) {
	example := Meeting{
		Orders: []Order{
			Order{
				Customer: "asdf",
				Items:    []OrderItem{OrderItem{Amount: 1, ItemName: "banana"}},
			},
		}}
	type args struct {
		customer CustomerMail
	}
	tests := []struct {
		name      string
		m         *Meeting
		args      args
		wantIndex int
		wantOrder Order
		wantFound bool
	}{
		{
			name: "Find order that exists",
			m:    &example,
			args: args{
				customer: "asdf",
			},
			wantIndex: 0,
			wantOrder: example.Orders[0],
			wantFound: true,
		},
		{
			name: "Find order that does not exist",
			m:    &example,
			args: args{
				customer: "b",
			},
			wantIndex: -1,
			wantOrder: Order{Customer: "b"},
			wantFound: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIndex, gotOrder, gotFound := tt.m.FindOrderByCustomer(tt.args.customer)
			if gotIndex != tt.wantIndex {
				t.Errorf("Meeting.FindOrderByCustomer() gotIndex = %v, want %v", gotIndex, tt.wantIndex)
			}
			if !reflect.DeepEqual(gotOrder, tt.wantOrder) {
				t.Errorf("Meeting.FindOrderByCustomer() gotOrder = %v, want %v", gotOrder, tt.wantOrder)
			}
			if gotFound != tt.wantFound {
				t.Errorf("Meeting.FindOrderByCustomer() gotFound = %v, want %v", gotFound, tt.wantFound)
			}
		})
	}
}

func TestMeeting_AddOrderItemForCustomer(t *testing.T) {
	type args struct {
		item     OrderItem
		customer CustomerMail
	}
	tests := []struct {
		name             string
		m                *Meeting
		args             args
		wantBananaAmount int
	}{
		{
			name: "add already existing product",
			m: &Meeting{
				Orders: []Order{
					Order{
						Customer: "asdf",
						Items:    []OrderItem{OrderItem{Amount: 1, ItemName: "banana"}},
					},
				}},
			args: args{
				item:     OrderItem{Amount: 2, ItemName: "banana"},
				customer: "asdf",
			},
			wantBananaAmount: 3,
		},
		{
			name: "add new product",
			m:    &Meeting{},
			args: args{
				item:     OrderItem{Amount: 2, ItemName: "banana"},
				customer: "asdf",
			},
			wantBananaAmount: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.AddOrderItemForCustomer(tt.args.item, tt.args.customer)
			_, order, found := tt.m.FindOrderByCustomer("asdf")
			if !found {
				t.Errorf("Order was not found for customer asdf")
			}
			_, bananaItem, found := order.FindItemByProductName("banana")
			if !found {
				t.Errorf("Banana item was not found in order of customer asdf")
			}
			if bananaItem.Amount != tt.wantBananaAmount {
				t.Errorf("got bananaAmount %v, want %v", bananaItem.Amount, tt.wantBananaAmount)
			}
		})
	}
}
