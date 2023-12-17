package apputils

import (
	"net/http"
	"reflect"
	"strings"

	"github.com/GDEIDevelopers/K8Sbackend/pkg/errhandle"
	"github.com/GDEIDevelopers/K8Sbackend/pkg/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

func ThrowError(c *gin.Context, err error) {
	errhandle.Log.Errorln(err)
	c.AbortWithStatusJSON(http.StatusBadRequest, &model.CommonResponse[any]{
		Status: errhandle.InnerError,
		Reason: err.Error(),
	})
}

func Throw(c *gin.Context, errCode errhandle.ErrCode) {
	c.AbortWithStatusJSON(http.StatusBadRequest, &model.CommonResponse[any]{
		Status: errCode,
		Reason: errCode.String(),
	})
}

func OK[T any](c *gin.Context, data T) {
	c.JSON(http.StatusOK, &model.CommonResponse[T]{
		Data: data,
	})
}

func BuildQuerySQL(tx *gorm.DB, query *model.QueryRequest) *gorm.DB {
	var where []string
	var params []any
	if query.UserID != 0 {
		where = append(where, "id = ?")
		params = append(params, query.UserID)
	}
	if query.Email != "" {
		where = append(where, "email = ?")
		params = append(params, query.Email)
	}
	if query.Name != "" {
		where = append(where, "name = ?")
		params = append(params, query.Name)
	}

	if len(where) == 0 {
		return nil
	}

	return tx.Where(strings.Join(where, " OR "), params...)
}

func reset(p any) any {
	var isPtr bool
	v := reflect.ValueOf(p)
	for v.Kind() == reflect.Ptr {
		isPtr = true
		v = v.Elem()
	}
	if isPtr {
		return reflect.New(v.Type()).Interface()

	}

	return reflect.Zero(v.Type()).Interface()

}

func IgnoreStructCopy(to, from any, ignore string) {
	copier.Copy(to, from)

	if ignore == "" {
		return
	}
	elem := reflect.Indirect(reflect.ValueOf(to))
	ignoreLower := strings.ToUpper(ignore[0:1]) + ignore[1:]
	for i := 0; i < elem.NumField(); i++ {
		current := elem.Field(i)
		if elem.Type().Field(i).Name == ignoreLower {
			current.Set(reflect.Zero(current.Type()))
		}
	}
}
