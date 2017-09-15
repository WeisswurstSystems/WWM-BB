package user

var (
	// RoleAdmin the super role for a user
	RoleAdmin = "admin"
)

// IsAdmin checks if the user is an admin
func (u User) IsAdmin() bool {
	return u.HasRole(RoleAdmin)
}

// HasRole finds out if the user has at a least one of the given roles.
// If the user is an admin he always wins.
func (u User) HasRole(roles ...string) bool {
	roles = append(roles, RoleAdmin)
	for _, want := range roles {
		for _, got := range u.Roles {
			if got == want {
				return true
			}
		}
	}
	return false
}

// HasMail finds out if the User has at least one of the given mails.
// If the user is an admin he always wins.
func (u User) HasMail(mails ...string) bool {
	if u.IsAdmin() {
		return true
	}
	for _, want := range mails {
		if u.Mail == want {
			return true
		}
	}
	return false
}
