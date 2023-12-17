package student

import (
	"github.com/GDEIDevelopers/K8Sbackend/app/apputils"
	"github.com/GDEIDevelopers/K8Sbackend/pkg/model"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/authrequired

// 获取学生信息 godoc
// @Summary 获取学生相关信息
// @Schemes
// @Description 获取学生相关信息
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
// @Router /student/{action} [get]
func (t *Student) Get(c *gin.Context) {
	apputils.OK[model.GetUserResponse](c, model.GetUserResponse{})
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
// @Router /register [post]
func (t *Student) RegisterStudent(c *gin.Context) {

}

// 修改学生信息 godoc
// @Summary 修改学生相关信息
// @Schemes
// @Description 修改学生相关信息
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
// @Router /student/{action} [patch]
func (t *Student) Modify(c *gin.Context) {

}
