package teacher

import (
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
// @Param   token     query    string  true   "登录返回的Token"
// @Success 200 {object} model.CommonResponse
// @Failure 400  {object} model.CommonResponse
// @Router /teacher/{action} [get]
func (t *Teacher) Get(c *gin.Context) {

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
// @Success 200 {object} model.CommonResponse
// @Failure 400  {object} model.CommonResponse
// @Router /teacher/student/new [post]
func (t *Teacher) RegisterStudent(c *gin.Context) {

}
