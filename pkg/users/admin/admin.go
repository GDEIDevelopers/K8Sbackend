package admin

import (
	"io"

	"github.com/GDEIDevelopers/K8Sbackend/app/apputils"
	"github.com/GDEIDevelopers/K8Sbackend/app/routes"
	"github.com/gin-gonic/gin"
)

type Admin struct {
	*apputils.ServerUtils
}

var _ io.Closer = &Admin{}

func New(db *apputils.ServerUtils) routes.Routes {
	return &Admin{db}
}

func (a *Admin) Close() error {
	return nil
}

func (a *Admin) InitRoute(g *gin.RouterGroup) {
	g.GET("/classes", a.ListClasses)

	admin := g.Group("/admin")
	admin.POST("/teacher/new", a.RegisterTeacher)
	admin.POST("/student/new", a.RegisterStudent)
	admin.POST("/admin/new", a.RegisterAdmin)

	admin.GET("/teachers/:action", a.GetTeachers)
	admin.GET("/teacher/:action", a.GetTeacher)
	admin.GET("/students/:action", a.GetStudents)
	admin.GET("/student/:action", a.GetStudent)

	admin.PATCH("/teacher/password", a.ModifyTeacherPassword)
	admin.PATCH("/student/password", a.ModifyStudentPassword)
	admin.PATCH("/admin/password", a.ModifyAdminPassword)
	admin.PATCH("/teacher", a.ModifyTeacher)
	admin.PATCH("/student", a.ModifyStudent)
	admin.PATCH("/admin", a.ModifyAdmin)

	admin.DELETE("/teacher", a.DeleteTeacher)
	admin.DELETE("/student", a.DeleteStudent)
	admin.DELETE("/admin", a.DeleteAdmin)

	// class part
	//admin.GET("/classes", a.ListClasses)
	admin.POST("/class/new", a.AddClass)
	admin.POST("/class/teachers", a.ListClassTeacher)
	admin.POST("/class/students", a.ListClassStudent)
	admin.POST("/teacher/students", a.ListTeacherStudent)
	admin.POST("/teacher/class/new", a.AddTeacherToClass)
	admin.POST("/student/class/new", a.AddStudentToClass)

	admin.DELETE("/class", a.RemoveClass)
	admin.DELETE("/teacher/class", a.RemoveTeacherFromClass)
	admin.DELETE("/student/class", a.RemoveStudentFromClass)

	admin.PATCH("/student/class", a.ModifyStudentClass)
}

func (a *Admin) InitCMDRoute(cmd *gin.Engine) {
	cmd.POST("/admin/new", a.RegisterAdmin)
	cmd.POST("/teacher/new", a.RegisterTeacher)
	cmd.POST("/student/new", a.RegisterStudent)
}
