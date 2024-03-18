package model

// Login represents the required information for a user login attempt.
type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
