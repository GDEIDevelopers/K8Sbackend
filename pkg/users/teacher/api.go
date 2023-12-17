package teacher

import (
	"encoding/json"

	"github.com/GDEIDevelopers/K8Sbackend/app/apputils"
	"github.com/GDEIDevelopers/K8Sbackend/pkg/errhandle"
	"github.com/GDEIDevelopers/K8Sbackend/pkg/model"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/authrequired

// 获取教师信息 godoc
// @Summary 获取教师相关信息
// @Schemes
// @Description 获取教师相关信息
// @Tags example
// @Accept json
// @Produce json
// @Param   action    path     string  false  "查询过滤器，如果没有默认查询所以信息"  Format(email)
// @Param   token     header    string  true   "登录返回的Token"
// @Param   userid    query    int     false  "用户ID"
// @Param   name      query    string  false  "用户名"
// @Param   email     query    string  false  "用户邮箱"
// @Success 200 {object} model.CommonResponse[model.GetUserResponse]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /teacher/{action} [get]
func (t *Teacher) Get(c *gin.Context) {
	b, err := c.GetRawData()
	if err != nil {
		apputils.ThrowError(c, err)
		return
	}
	var query model.QueryRequest
	json.Unmarshal(b, &query)
	tx := apputils.BuildQuerySQL(t.DB.Table("users"), &query)
	if tx == nil {
		apputils.Throw(c, errhandle.ParamsError)
		return
	}
	var teacher model.User
	err = tx.First(&teacher).Error
	if err != nil {
		apputils.ThrowError(c, err)
		return
	}
	var getResponse model.GetUserResponse

	apputils.IgnoreStructCopy(&getResponse, &teacher, c.Param("action"))

	apputils.OK[model.GetUserResponse](c, getResponse)
}

// 添加/注册一位学生 godoc
// @Summary 添加/注册一个学生
// @Schemes
// @Description 添加/注册一个学生
// @Tags example
// @Accept json
// @Produce json
// @Param   token     query    string  true   "登录返回的Token"
// @Param   name      query    string  true   "登录返回的Token"
// @Param   token     query    string  true   "登录返回的Token"
// @Success 200 {object} model.CommonResponse[any]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /teacher/student/new [post]
func (t *Teacher) RegisterStudent(c *gin.Context) {

}
