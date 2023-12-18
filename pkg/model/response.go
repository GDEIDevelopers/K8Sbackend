package model

import (
	"github.com/GDEIDevelopers/K8Sbackend/pkg/errhandle"
)

type CommonResponse[T any] struct {
	Data   T                 `json:"data,omitempty"`
	Status errhandle.ErrCode `json:"status"`
	Reason string            `json:"reason"`
}

type TokenResponse struct {
	AccessToken  string `json:"token"`
	RefreshToken string `json:"refreshtoken"`
	Scope        string `json:"scope"`
	ExpiredAt    int64  `json:"expiredAt"`
}

type UserInfo struct {
	UserID int64  `json:"userid"`
	Name   string `json:"name"`
	Role   string `json:"role"`
}

type GetUserResponse struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email,omitempty"`
	RealName     string `json:"realName,omitempty"`
	UserSchoollD string `json:"userSchoollD,omitempty"`
	SchoolCode   string `json:"schoolCode,omitempty"`
	Class        string `json:"class,omitempty"`
	Sex          string `json:"sex,omitempty"`
}
