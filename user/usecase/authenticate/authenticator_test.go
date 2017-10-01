package authenticate

import (
	"reflect"
	"testing"

	"github.com/WeisswurstSystems/WWM-BB/user"
)

func TestInteractor_Authenticate(t *testing.T) {
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
	// TODO: Add test cases.
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

func Test_authenticated(t *testing.T) {
	type args struct {
		user user.User
		l    user.Login
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "user has not yet registered",
			args: args{
				user: user.User{Login: user.Hans.Login, RegistrationID: "asdf"},
				l:    user.Hans.Login,
			},
			want: false,
		},
		{
			name: "user has wrong password",
			args: args{
				user: user.Hans,
				l:    user.Login{Mail: "wrong", Password: user.Hans.Login.Password},
			},
			want: false,
		},
		{
			name: "user has wrong mail",
			args: args{
				user: user.Hans,
				l:    user.Login{Password: "wrong", Mail: user.Hans.Login.Mail},
			},
			want: false,
		},
		{
			name: "user has right login",
			args: args{
				user: user.Hans,
				l:    user.Hans.Login,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := authenticated(tt.args.user, tt.args.l); got != tt.want {
				t.Errorf("authenticated() = %v, want %v", got, tt.want)
			}
		})
	}
}
