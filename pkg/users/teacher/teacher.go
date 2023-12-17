package teacher

import (
	"io"

	"github.com/GDEIDevelopers/K8Sbackend/app/apputils"
)

type Teacher struct {
	db *apputils.ServerUtils
}

var _ io.Closer = &Teacher{}

func New(db *apputils.ServerUtils) *Teacher {
	return &Teacher{db}
}

func (a *Teacher) Close() error {
	return nil
}
