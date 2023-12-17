package app

import (
	"context"
	"strings"

	"github.com/GDEIDevelopers/K8Sbackend/app/apputils"
	"github.com/GDEIDevelopers/K8Sbackend/pkg/errhandle"
	"github.com/gin-gonic/gin"
)

func (s *Server) UseTokenVerify() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.Request.Header.Get("Authorization")
		prefix := "Bearer "
		token := ""

		if auth != "" && strings.HasPrefix(auth, prefix) {
			token = auth[len(prefix):]
		}

		if token == "" {
			apputils.Throw(ctx, errhandle.TokenError)
			return
		}

		userinfo, ok := s.token.Verify(context.Background(), token)
		if !ok {
			apputils.Throw(ctx, errhandle.PermissionDenied)
			return
		}

		path := strings.Split(ctx.Request.URL.Path, "/")
		if len(path) >= 4 {
			switch path[3] {
			case "teacher":
				if userinfo.Role == "student" {
					apputils.Throw(ctx, errhandle.PermissionDenied)
					return
				}
			case "admin":
				if userinfo.Role != "admin" {
					apputils.Throw(ctx, errhandle.PermissionDenied)
					return
				}
			}
		}
		ctx.Set("info", userinfo)
		ctx.Next()
	}
}

func (s *Server) UserLogin(c *gin.Context) {

}

func (s *Server) UserRegister(c *gin.Context) {

}
