package middleware

func RoleAdmin(role string) bool {
	if role != "admin" {
		return false
	}

	return true
}
