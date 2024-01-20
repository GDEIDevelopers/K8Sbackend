package app

import (
	"context"
	"strings"
	"time"

	"github.com/GDEIDevelopers/K8Sbackend/app/apputils"
	"github.com/GDEIDevelopers/K8Sbackend/pkg/errhandle"
	"github.com/gin-gonic/gin"
)

const (
	RefreshTokenExpired = 24 * time.Hour * 3
	AccessTokenExpired  = 2 * time.Hour
)

func (s *Server) UseCORS() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Accept, Origin, Cache-Control, X-Requested-With")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}
		ctx.Next()
	}
}

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

		userinfo, ok := s.Token.Verify(context.Background(), token)
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

// @BasePath /api
