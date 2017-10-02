package authenticate

import "github.com/WeisswurstSystems/WWM-BB/user"

type defectAuthenticator struct {
	defect error
}

func NewDefectAuthenticator(defect error) AuthenticateUseCase {
	return &defectAuthenticator{
		defect: defect,
	}
}

func (aa defectAuthenticator) Authenticate(login user.Login) (user.User, error) {
	return user.User{}, aa.defect
}
