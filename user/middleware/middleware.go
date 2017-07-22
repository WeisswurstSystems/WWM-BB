package middleware

import (
	"github.com/WeisswurstSystems/WWM-BB/user"
)

type MiddlewareHandler struct {
	UserStore user.Store
}
