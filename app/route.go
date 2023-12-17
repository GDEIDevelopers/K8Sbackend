package app

import (
	"net/http"

	docs "github.com/GDEIDevelopers/K8Sbackend/docs"
	"github.com/GDEIDevelopers/K8Sbackend/pkg/errhandle"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (s *Server) dispatchRoute() {
	docs.SwaggerInfo.BasePath = "/api"
	e := gin.Default()

	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	a := e.Group("/api")

	a.POST("/login", s.UserLogin)
	a.POST("/refresh", s.UserLoginRefresh)
	a.POST("/register", s.student.RegisterStudent)

	requiredAuth := a.Group("/authrequired")
	requiredAuth.Use(s.UseTokenVerify())

	// teacher part
	teacher := requiredAuth.Group("/teacher")
	teacher.POST("/student/new", s.teacher.RegisterStudent)
	teacher.GET("/:action", s.teacher.Get)
	teacher.PATCH("/:action", s.teacher.Modify)

	// student part
	student := requiredAuth.Group("/student")
	student.GET("/:action", s.student.Get)
	student.PATCH("/:action", s.student.Modify)
	// admin part

	admin := requiredAuth.Group("/admin")
	admin.POST("/teacher/new", s.admin.RegistserTeacher)
	admin.POST("/student/new", s.admin.RegistserStudent)

	admin.GET("/teachers/:action", s.admin.GetTeachers)
	admin.GET("/teacher/:action", s.admin.GetTeacher)
	admin.GET("/students/:action", s.admin.GetStudents)
	admin.GET("/student/:action", s.admin.GetStudent)

	admin.DELETE("/teacher")
	admin.DELETE("/student")

	s.setupHTTPServer(e)
}

func (s *Server) setupHTTPServer(e *gin.Engine) {
	s.srv = &http.Server{
		Addr:    s.Config.HTTPServerListen,
		Handler: e,
	}

	go s.srv.ListenAndServe()

	errhandle.Log.Info("HTTP Server Starts At %s", s.srv.Addr)
}
