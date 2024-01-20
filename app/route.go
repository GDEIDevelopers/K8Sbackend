package app

import (
	"net/http"

	"github.com/GDEIDevelopers/K8Sbackend/app/routes"
	docs "github.com/GDEIDevelopers/K8Sbackend/docs"
	"github.com/GDEIDevelopers/K8Sbackend/pkg/errhandle"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (s *Server) initCMDRoutes(g *gin.Engine) {
	for _, r := range s.routes {
		if route, ok := r.(routes.CMDRoutes); ok {
			route.InitCMDRoute(g)
		}
	}
}

func (s *Server) initAuthRequiedRoutes(g *gin.RouterGroup) {
	for _, r := range s.routes {
		r.InitRoute(g)
	}
}

func (s *Server) initGlobalRoutes(g *gin.RouterGroup) {
	for _, r := range s.routes {
		if route, ok := r.(routes.GlobalRoutes); ok {
			route.InitGlobalRoute(g)
		}
	}
}

func (s *Server) dispatchRoute() {
	docs.SwaggerInfo.BasePath = "/api"
	e := gin.Default()
	e.Use(s.UseCORS())
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	cmd := gin.Default()
	s.initCMDRoutes(cmd)

	api := e.Group("/api")
	s.initGlobalRoutes(api)

	requiredAuth := api.Group("/authrequired")
	requiredAuth.Use(s.UseTokenVerify())
	s.initAuthRequiedRoutes(requiredAuth)

	s.setupHTTPServer(e, cmd)
}

func (s *Server) setupHTTPServer(e, cmd *gin.Engine) {
	s.srv = &http.Server{
		Addr:    s.Config.HTTPServerListen,
		Handler: e,
	}

	cmdServer := &http.Server{
		Addr:    s.Config.CMDHTTPServerListen,
		Handler: cmd,
	}

	go s.srv.ListenAndServe()
	go cmdServer.ListenAndServe()

	errhandle.Infof("HTTP Server Starts At %s", s.srv.Addr)
	errhandle.Infof("CMD HTTP Server Starts At %s", cmdServer.Addr)
}
