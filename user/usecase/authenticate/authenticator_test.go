package authenticate

import (
	"errors"
	"reflect"
	"testing"

	"github.com/WeisswurstSystems/WWM-BB/user"
	"github.com/WeisswurstSystems/WWM-BB/user/driver"
)

func TestInteractor_Authenticate(t *testing.T) {
	storeWithHans := driver.NewMemoryStore()
	storeWithHans.Save(user.Hans)

	storeWithUnregisteredHans := driver.NewMemoryStore()
	unregistered := user.Hans
	unregistered.RegistrationID = "notyetcleared"
	storeWithUnregisteredHans.Save(unregistered)

	type args struct {
		l user.Login
	}
	tests := []struct {
		name    string
		i       Interactor
		args    args
		want    user.User
		wantErr bool
	}{
		{
			name:    "user not found",
			i:       Interactor{ReadStore: driver.NewMemoryStore()},
			args:    args{l: user.Hans.Login},
			want:    user.User{},
			wantErr: true,
		},
		{
			name:    "store error",
			i:       Interactor{ReadStore: driver.NewDefectStore(errors.New("asdf"))},
			args:    args{l: user.Hans.Login},
			want:    user.User{},
			wantErr: true,
		},
		{
			name:    "wrong login error",
			i:       Interactor{ReadStore: storeWithHans},
			args:    args{l: user.Login{Mail: "wrong", Password: "wrong"}},
			want:    user.User{},
			wantErr: true,
		},
		{
			name:    "not registered error",
			i:       Interactor{ReadStore: storeWithUnregisteredHans},
			args:    args{l: user.Hans.Login},
			want:    user.User{},
			wantErr: true,
		},
		{
			name:    "succesful login",
			i:       Interactor{ReadStore: storeWithHans},
			args:    args{l: user.Hans.Login},
			want:    user.Hans,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.i.Authenticate(tt.args.l)
			if (err != nil) != tt.wantErr {
				t.Errorf("Interactor.Authenticate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Interactor.Authenticate() = %v, want %v", got, tt.want)
			}
		})
	}
}
