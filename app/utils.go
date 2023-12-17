package app

import (
	"encoding/json"
	"net/http"

	"github.com/GDEIDevelopers/K8Sbackend/pkg/errhandle"
	"github.com/GDEIDevelopers/K8Sbackend/pkg/model"
	"github.com/gin-gonic/gin"
)

func ThrowError(c *gin.Context, err error) {
	errhandle.Log.Errorln(err)
	c.AbortWithStatusJSON(http.StatusBadRequest, &model.CommonResponse{
		Status: errhandle.InnerError,
		Reason: err.Error(),
	})
}

func Throw(c *gin.Context, errCode errhandle.ErrCode) {
	c.AbortWithStatusJSON(http.StatusBadRequest, &model.CommonResponse{
		Status: errCode,
		Reason: errCode.String(),
	})
}

func OK(c *gin.Context, data any) {
	b, _ := json.Marshal(data)
	c.JSON(http.StatusOK, &model.CommonResponse{
		Data: b,
	})
}
