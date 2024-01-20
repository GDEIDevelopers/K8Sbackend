package auth

import (
	"github.com/GDEIDevelopers/K8Sbackend/app/apputils"
	"github.com/GDEIDevelopers/K8Sbackend/app/routes"
	"github.com/gin-gonic/gin"
)

type Auth struct {
	*apputils.ServerUtils
}

func New(db *apputils.ServerUtils) routes.Routes {
	return &Auth{db}
}

func (a *Auth) InitRoute(g *gin.RouterGroup) {
}

func (a *Auth) InitGlobalRoute(g *gin.RouterGroup) {
	g.GET("/isvalid", a.IsValidSession)
	g.POST("/login", a.UserLogin)
	g.POST("/refresh", a.UserLoginRefresh)
}
