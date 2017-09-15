package user

import "testing"

func TestUser_IsAdmin(t *testing.T) {
	tests := []struct {
		name string
		u    User
		want bool
	}{
		{
			"User is admin",
			User{Roles: []string{RoleAdmin}},
			true,
		},
		{
			"User is not an admin",
			User{},
			false,
		},
		{
			"User has another role",
			User{Roles: []string{"asdf"}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.u.IsAdmin(); got != tt.want {
				t.Errorf("User.IsAdmin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_HasRole(t *testing.T) {
	type args struct {
		roles []string
	}
	tests := []struct {
		name string
		u    User
		args args
		want bool
	}{
		{
			"Want role A",
			User{Roles: []string{"A", "B"}},
			args{[]string{"A"}},
			true,
		},
		{
			"Want role B",
			User{Roles: []string{"A", "B"}},
			args{[]string{"B"}},
			true,
		},
		{
			"Should not have role A",
			User{Roles: []string{"B"}},
			args{[]string{"A"}},
			false,
		},
		{
			"Should not have role A but is admin",
			User{Roles: []string{"admin", "B"}},
			args{[]string{"A"}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.u.HasRole(tt.args.roles...); got != tt.want {
				t.Errorf("User.HasRole() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_HasMail(t *testing.T) {
	type args struct {
		mails []string
	}
	tests := []struct {
		name string
		u    User
		args args
		want bool
	}{
		{
			"Has mail peter7995@gmail.com",
			User{Login: Login{Mail: "peter7995@gmail.com"}},
			args{[]string{"peter7995@gmail.com"}},
			true,
		},
		{
			"Has mail asdf@gmail.com",
			User{Login: Login{Mail: "asdf@gmail.com"}},
			args{[]string{"peter7995@gmail.com", "asdf@gmail.com"}},
			true,
		},
		{
			"Has not mail asdf@gmail.com",
			User{Login: Login{Mail: "asdf@gmail.com"}},
			args{[]string{"peter7995@gmail.com"}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.u.HasMail(tt.args.mails...); got != tt.want {
				t.Errorf("User.HasMail() = %v, want %v", got, tt.want)
			}
		})
	}
}
