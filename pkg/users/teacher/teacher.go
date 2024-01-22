package teacher

import (
	"io"

	"github.com/GDEIDevelopers/K8Sbackend/app/apputils"
	"github.com/GDEIDevelopers/K8Sbackend/app/routes"
	"github.com/gin-gonic/gin"
)

type Teacher struct {
	*apputils.ServerUtils
}

var _ io.Closer = &Teacher{}

func New(db *apputils.ServerUtils) routes.Routes {
	return &Teacher{db}
}

func (a *Teacher) Close() error {
	return nil
}

func (a *Teacher) InitRoute(g *gin.RouterGroup) {
	teacher := g.Group("/teacher")
	teacher.POST("/student/new", a.RegisterStudent)
	teacher.PATCH("/password", a.ModifyPassword)
	teacher.PATCH("/", a.Modify)

	// class part
	teacher.GET("/class", a.ListJoinedClass)
	teacher.GET("/class/students", a.ListStudents)
	teacher.GET("/class/students/:classname", a.ListStudents)
	teacher.POST("/class/join", a.Join)
	teacher.POST("/class/students/join", a.AddStudents)

	teacher.DELETE("/class/leave", a.Leave)
	teacher.DELETE("/class/students/leave", a.RemoveStudents)
}
