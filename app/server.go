package app

import (
	"context"
	"net/http"
	"time"

	"github.com/GDEIDevelopers/K8Sbackend/app/apputils"
	"github.com/GDEIDevelopers/K8Sbackend/config"
	"github.com/GDEIDevelopers/K8Sbackend/pkg/errhandle"
	"github.com/GDEIDevelopers/K8Sbackend/pkg/token"
	"github.com/GDEIDevelopers/K8Sbackend/pkg/users/admin"
	"github.com/GDEIDevelopers/K8Sbackend/pkg/users/student"
	"github.com/GDEIDevelopers/K8Sbackend/pkg/users/teacher"
	"github.com/golang-jwt/jwt/v5"
)

type Server struct {
	*apputils.ServerUtils

	srv   *http.Server
	token token.TokenGenerate
	// user management parts
	admin   *admin.Admin
	student *student.Student
	teacher *teacher.Teacher
}

func NewServer(cfg *config.Config) *Server {
	utils := apputils.NewServerUtils(cfg)
	s := &Server{
		ServerUtils: utils,
		admin:       admin.New(utils),
		teacher:     teacher.New(utils),
		student:     student.New(utils),
		token:       token.NewJWTAccessGenerate(utils.RedisDB, jwt.SigningMethodHS256),
	}
	s.dispatchRoute()
	return s
}

func (s *Server) shutdownHTTPServer() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := s.srv.Shutdown(ctx); err != nil {
		errhandle.Log.Fatal(err)
	}
}

func (s *Server) Close() {
	s.admin.Close()
	s.teacher.Close()
	s.student.Close()
	s.shutdownHTTPServer()
}
