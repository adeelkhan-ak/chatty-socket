package models

// Message Define our message object
type Message struct {
	Username  string `json:"username"`
	Message   string `json:"message"`
	Email     string `json:"email"`
	UpdatedAt string `json:"updated_at"`
}
