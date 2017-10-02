package order

import (
	"testing"

	"github.com/WeisswurstSystems/WWM-BB/meeting"

	"github.com/WeisswurstSystems/WWM-BB/meeting/driver"
	"github.com/WeisswurstSystems/WWM-BB/user"
	"github.com/WeisswurstSystems/WWM-BB/user/usecase/authenticate"
)

func TestInteractor_Order(t *testing.T) {
	type args struct {
		req Request
	}
	tests := []struct {
		name            string
		i               *Interactor
		args            args
		wantErr         bool
		wantBananaCount int
	}{
		{
			name: "order a single item",
			i: &Interactor{
				AuthenticateUseCase: authenticate.NewAlwaysAuthenticator(),
				Store: func() meeting.Store {
					store := driver.NewMemoryStore()
					store.Save(meeting.Meeting{
						ID: meeting.MeetingID("123"),
					})
					return store
				}(),
			},
			args: args{req: Request{
				Login:     user.Login{Mail: "asdf"},
				Item:      meeting.OrderItem{Amount: 1, ItemName: "banana"},
				MeetingID: meeting.MeetingID("123"),
			}},
			wantErr:         false,
			wantBananaCount: 1,
		},
		{
			name: "meeting does not exist",
			i: &Interactor{
				AuthenticateUseCase: authenticate.NewAlwaysAuthenticator(),
				Store:               driver.NewMemoryStore(),
			},
			args: args{req: Request{
				Login:     user.Login{Mail: "asdf"},
				Item:      meeting.OrderItem{Amount: 1, ItemName: "banana"},
				MeetingID: meeting.MeetingID("notfound"),
			}},
			wantErr: true,
		},
		{
			name: "not authenticated",
			i: &Interactor{
				AuthenticateUseCase: authenticate.NewDefectAuthenticator(authenticate.ErrNotAuthenticated),
				Store: func() meeting.Store {
					store := driver.NewMemoryStore()
					store.Save(meeting.Meeting{
						ID: meeting.MeetingID("123"),
					})
					return store
				}(),
			},
			args: args{req: Request{
				Login:     user.Login{Mail: "asdf"},
				Item:      meeting.OrderItem{Amount: 1, ItemName: "banana"},
				MeetingID: meeting.MeetingID("123"),
			}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.i.Order(tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("Interactor.Order() error = %v, wantErr %v", err, tt.wantErr)
			}
			m, _ := tt.i.FindOne(meeting.MeetingID("123"))
			_, o, _ := m.FindOrderByCustomer("asdf")
			_, item, _ := o.FindItemByProductName("banana")
			if item.Amount != tt.wantBananaCount {
				t.Errorf("wantBananaCount %v, got %v", tt.wantBananaCount, item.Amount)
			}
		})
	}
}
