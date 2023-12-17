package model

import "encoding/json"

type CommonRequest struct {
	Token string          `json:"token"`
	Data  json.RawMessage `json:"data"`
}
