package apputils

import (
	"net/http"
	"reflect"
	"strings"
	"unicode"

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

func BuildQuerySQL(tx *gorm.DB, query *model.QueryRequest, role ...string) *gorm.DB {
	var where []string
	var params []any
	if query.UserID != 0 {
		where = append(where, "id = ?")
		params = append(params, query.UserID)
	}
	if query.QueryEmail != "" {
		where = append(where, "email = ?")
		params = append(params, query.QueryEmail)
	}
	if query.Name != "" {
		where = append(where, "name = ?")
		params = append(params, query.Name)
	}
	if query.QueryUserSchoollD != "" {
		where = append(where, "userSchoollD LIKE ?")
		params = append(params, query.QueryUserSchoollD+"%")
	}
	if query.QueryRealName != "" {
		where = append(where, "realName LIKE ?")
		params = append(params, query.QueryRealName+"%")
	}

	if len(where) == 0 {
		return nil
	}

	whereStatement := strings.Join(where, " AND ")

	if len(role) > 0 {
		whereStatement += " AND role = ?"
		params = append(params, role[0])
	}
	return tx.Where(whereStatement, params...)
}

func IgnoreStructCopy(to, from any, ignore string) {
	copier.CopyWithOption(to, from, copier.Option{
		IgnoreEmpty: true,
	})

	if ignore == "" {
		return
	}
	elem := reflect.Indirect(reflect.ValueOf(to))
	ignoreLower := strings.ToUpper(ignore[0:1]) + ignore[1:]
	for i := 0; i < elem.NumField(); i++ {
		current := elem.Field(i)
		if elem.Type().Field(i).Name == ignoreLower {
			current.Set(reflect.Zero(current.Type()))
			break
		}
	}
}

func IsValidPassword(password string) bool {
	if len(password) < 8 {
		return false
	}
	var upper, lower, number int

	for _, char := range password {
		switch {
		case unicode.IsLower(char):
			lower++
		case unicode.IsUpper(char):
			upper++
		case unicode.IsNumber(char):
			number++
		}
	}
	return upper > 0 && lower > 0 && number > 0
}
