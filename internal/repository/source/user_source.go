package source

const (
	USER_INSERT         = "INSERT INTO users(name, email, password, role) VALUES(?,?,?,?)"
	USER_FIND_BY_EMAIL  = "SELECT id, name, email, password, role, created_at FROM users WHERE email = ?"
	USER_COUNT_BY_EMAIL = "SELECT COUNT(*) FROM users WHERE email = ?"
)
