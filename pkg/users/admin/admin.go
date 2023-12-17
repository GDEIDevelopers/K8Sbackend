package admin

import (
	"io"

	"github.com/GDEIDevelopers/K8Sbackend/app/apputils"
)

type Admin struct {
	db *apputils.ServerUtils
}

var _ io.Closer = &Admin{}

func New(db *apputils.ServerUtils) *Admin {
	return &Admin{db}
}

func (a *Admin) Close() error {
	return nil
}
