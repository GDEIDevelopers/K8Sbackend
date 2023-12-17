package model

import (
	"encoding/json"

	"github.com/GDEIDevelopers/K8Sbackend/pkg/errhandle"
)

type CommonResponse struct {
	Data   json.RawMessage   `json:"data"`
	Status errhandle.ErrCode `json:"status"`
	Reason string            `json:"reason"`
}

type TokenResponse struct {
	AccessToken  string `json:"token"`
	RefreshToken string `json:"refreshtoken"`
	Scope        string `json:"scope"`
	ExpiredAt    int64
}

type UserInfo struct {
	UserID int64  `json:"userid"`
	Name   string `json:"name"`
	Role   string
}
