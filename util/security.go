package util

import (
	"github.com/WeisswurstSystems/WWM-BB/meeting"
	"github.com/WeisswurstSystems/WWM-BB/user"
)

func IsMeetingCreator(u user.User, m meeting.Meeting) bool {
	if u.Mail == m.Creator {
	return true
	}
	return adminCheck(u)
}

func IsMeetingCreatorOrBuyer(u user.User, m meeting.Meeting) bool {
	if u.Mail == m.Creator {
		return true
	}
	if u.Mail == m.Buyer {
		return true
	}
	return adminCheck(u)
}

func IsOrderCustomer(u user.User, o meeting.Order) bool {
	if u.Mail == o.Customer {
		return true
	}
	return adminCheck(u)
}

func adminCheck(u user.User) bool {
	return Contains(u.Roles, "admin")
}