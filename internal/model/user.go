package model

type User struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	CreatedAt    string `json:"created_at"`
}

type Register struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

type LoginRequest struct {
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}
