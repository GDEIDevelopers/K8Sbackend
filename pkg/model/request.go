package model

type QueryRequest struct {
	UserID     int    `json:"id"`
	Name       string `json:"name"`
	QueryEmail string `json:"queryemail"`
}

type RegisterUserRequest struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	RealName     string `json:"realName"`
	UserSchoollD string `json:"userSchoollD"`
	SchoolCode   string `json:"schoolCode"`
	Class        string `json:"class"`
	Sex          string `json:"sex"`
	Password     string `json:"password"`
}

type UserLoginRequest struct {
	QueryRequest
	Password string `json:"password"`
}

type ModifyUserRequest struct {
	Email        string `json:"email,omitempty"`
	RealName     string `json:"realName,omitempty"`
	UserSchoollD string `json:"userSchoollD,omitempty"`
	SchoolCode   string `json:"schoolCode,omitempty"`
	Class        string `json:"class,omitempty"`
	Sex          string `json:"sex,omitempty"`
}
type ModifyUserPasswordRequest struct {
	Password string `json:"password"`
}

type AdminModifyPasswordRequest struct {
	QueryRequest
	Password string `json:"password"`
}

type AdminModifyRequest struct {
	QueryRequest
	Email        string `json:"email,omitempty"`
	RealName     string `json:"realName,omitempty"`
	UserSchoollD string `json:"userSchoollD,omitempty"`
	SchoolCode   string `json:"schoolCode,omitempty"`
	Class        string `json:"class,omitempty"`
	Sex          string `json:"sex,omitempty"`
}
