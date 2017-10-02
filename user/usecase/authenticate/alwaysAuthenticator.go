package authenticate

import "github.com/WeisswurstSystems/WWM-BB/user"

type alwaysAuthenticator struct {
}

func NewAlwaysAuthenticator() AuthenticateUseCase {
	return &alwaysAuthenticator{}
}

func (aa alwaysAuthenticator) Authenticate(login user.Login) (user.User, error) {
	return user.User{
		Login: login,
		Roles: []string{"admin"},
	}, nil
}
