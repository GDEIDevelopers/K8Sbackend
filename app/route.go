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
	e.Use(s.UseCORS())

	cmd := gin.Default()
	cmd.POST("/admin/new", s.admin.RegisterAdmin)
	cmd.POST("/teacher/new", s.admin.RegisterTeacher)
	cmd.POST("/student/new", s.admin.RegisterStudent)

	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	a := e.Group("/api")
	a.GET("/isvalid", s.IsValidSession)
	a.POST("/login", s.UserLogin)
	a.POST("/refresh", s.UserLoginRefresh)
	a.POST("/register", s.student.RegisterStudent)

	requiredAuth := a.Group("/authrequired")
	requiredAuth.Use(s.UseTokenVerify())

	// teacher part
	teacher := requiredAuth.Group("/teacher")
	teacher.POST("/student/new", s.teacher.RegisterStudent)
	teacher.GET("/:action", s.teacher.Get)
	teacher.PATCH("/password", s.teacher.Modify)
	teacher.PATCH("/", s.teacher.Modify)

	// student part
	student := requiredAuth.Group("/student")
	student.GET("/:action", s.student.Get)
	student.PATCH("/:action", s.student.Modify)
	// admin part

	admin := requiredAuth.Group("/admin")
	admin.POST("/teacher/new", s.admin.RegisterTeacher)
	admin.POST("/student/new", s.admin.RegisterStudent)
	admin.POST("/admin/new", s.admin.RegisterAdmin)

	admin.GET("/teachers/:action", s.admin.GetTeachers)
	admin.GET("/teacher/:action", s.admin.GetTeacher)
	admin.GET("/students/:action", s.admin.GetStudents)
	admin.GET("/student/:action", s.admin.GetStudent)

	admin.PATCH("/teacher/password", s.admin.ModifyTeacherPassword)
	admin.PATCH("/student/password", s.admin.ModifyStudentPassword)
	admin.PATCH("/admin/password", s.admin.ModifyAdminPassword)
	admin.PATCH("/teacher", s.admin.ModifyTeacher)
	admin.PATCH("/student", s.admin.ModifyStudent)
	admin.PATCH("/admin", s.admin.ModifyAdmin)

	admin.DELETE("/teacher", s.admin.DeleteTeacher)
	admin.DELETE("/student", s.admin.DeleteStudent)
	admin.DELETE("/admin", s.admin.DeleteAdmin)
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
