package student

import (
	"io"

	"github.com/GDEIDevelopers/K8Sbackend/app/apputils"
)

type Student struct {
	 *apputils.ServerUtils
}

var _ io.Closer = &Student{}

func New(db *apputils.ServerUtils) *Student {
	return &Student{db}
}

func (a *Student) Close() error {
	return nil
}
