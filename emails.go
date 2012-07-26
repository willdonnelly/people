package people

// Produce a random email of the form 'username@example.com' (where 'username' is produced by Username())
func Email() string {
	return Username() + "@example.com"
}

