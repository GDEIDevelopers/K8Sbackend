package app

import (
	"context"
	"net/http"
	"time"

	"github.com/GDEIDevelopers/K8Sbackend/app/apputils"
	"github.com/GDEIDevelopers/K8Sbackend/app/routes"
	"github.com/GDEIDevelopers/K8Sbackend/config"
	"github.com/GDEIDevelopers/K8Sbackend/pkg/errhandle"
	"github.com/GDEIDevelopers/K8Sbackend/pkg/users/admin"
	"github.com/GDEIDevelopers/K8Sbackend/pkg/users/auth"
	"github.com/GDEIDevelopers/K8Sbackend/pkg/users/student"
	"github.com/GDEIDevelopers/K8Sbackend/pkg/users/teacher"
)

type Server struct {
	*apputils.ServerUtils
	srv    *http.Server
	routes []routes.Routes
}

func initModules(s *apputils.ServerUtils) []routes.Routes {
	return []routes.Routes{
		auth.New(s),
		admin.New(s),
		teacher.New(s),
		student.New(s),
	}
}

func NewServer(cfg *config.Config) *Server {
	utils := apputils.NewServerUtils(cfg)
	s := &Server{
		ServerUtils: utils,
		routes:      initModules(utils),
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
	s.shutdownHTTPServer()
}
