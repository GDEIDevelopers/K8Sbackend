package app

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/GDEIDevelopers/K8Sbackend/pkg/errhandle"
	"github.com/GDEIDevelopers/K8Sbackend/pkg/model"
	"github.com/gin-gonic/gin"
)

func (s *Server) UseTokenVerify() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		b, err := ctx.GetRawData()
		if err != nil {
			ThrowError(ctx, err)
			return
		}
		var req model.CommonRequest
		json.Unmarshal(b, &req)
		if req.Token == "" {
			Throw(ctx, errhandle.TokenError)
			return
		}

		userinfo, ok := s.token.Verify(context.Background(), req.Token)
		if !ok {
			Throw(ctx, errhandle.PermissionDenied)
			return
		}

		path := strings.Split(ctx.Request.URL.Path, "/")
		if len(path) >= 4 {
			switch path[3] {
			case "teacher":
				if userinfo.Role == "student" {
					Throw(ctx, errhandle.PermissionDenied)
					return
				}
			case "admin":
				if userinfo.Role != "admin" {
					Throw(ctx, errhandle.PermissionDenied)
					return
				}
			}
		}
		ctx.Set("data", req.Data)
		ctx.Set("info", userinfo)
		ctx.Next()
	}
}

func (s *Server) UserLogin(c *gin.Context) {

}

func (s *Server) UserRegister(c *gin.Context) {

}
