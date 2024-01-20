package student

import (
	"io"

	"github.com/GDEIDevelopers/K8Sbackend/app/apputils"
	"github.com/GDEIDevelopers/K8Sbackend/app/routes"
	"github.com/gin-gonic/gin"
)

type Student struct {
	*apputils.ServerUtils
}

var _ io.Closer = &Student{}

func New(db *apputils.ServerUtils) routes.Routes {
	return &Student{db}
}

func (a *Student) Close() error {
	return nil
}
func (a *Student) InitRoute(g *gin.RouterGroup) {
	student := g.Group("/student")
	student.PATCH("/:action", a.Modify)

	student.POST("/student/class/join", a.Join)
	student.DELETE("/student/class/leave", a.Leave)
	student.PATCH("/student/class", a.ClassModify)
}

func (a *Student) InitGlobalRoute(g *gin.RouterGroup) {
	g.POST("/register", a.RegisterStudent)
}
