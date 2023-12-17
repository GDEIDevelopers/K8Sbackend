package model

type QueryRequest struct {
	UserID int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

type RegisterUserRequest struct {
}
