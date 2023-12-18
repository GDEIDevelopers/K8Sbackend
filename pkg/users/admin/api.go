package admin

import (
	"encoding/json"

	"github.com/GDEIDevelopers/K8Sbackend/app/apputils"
	"github.com/GDEIDevelopers/K8Sbackend/pkg/errhandle"
	"github.com/GDEIDevelopers/K8Sbackend/pkg/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// 获取所有教师信息 godoc
// @Summary 获取所有教师信息
// @Schemes
// @Description 获取所有教师信息
// @Tags example
// @Accept json
// @Produce json
// @Param   action    path     string  false  "查询过滤器，如果没有默认查询所以信息"
// @Param   token     header    string  true   "登录返回的Token"
// @Success 200 {object} model.CommonResponse[[]model.GetUserResponse]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/admin/teachers/{action} [get]
func (t *Admin) GetTeachers(c *gin.Context) {
	var teachers []*model.User

	err := t.DB.Table("users").
		Where("role = ?", "teacher").
		Find(&teachers).Error

	if err != nil {
		apputils.ThrowError(c, err)
		return
	}

	var teacherRes []model.GetUserResponse

	for _, tea := range teachers {
		var teacher model.GetUserResponse
		apputils.IgnoreStructCopy(&teacher, &tea, c.Param("action"))
		teacherRes = append(teacherRes, teacher)
	}

	apputils.OK[[]model.GetUserResponse](c, teacherRes)
}

// 获得指定教师信息 godoc
// @Summary 获取指定教师信息
// @Schemes
// @Description 获取指定教师信息
// @Tags example
// @Accept json
// @Produce json
// @Param   action    path     string  false  "查询过滤器，如果没有默认查询所以信息"
// @Param   token     header    string  true   "登录返回的Token"
// @Param   queryemail     query     string  false  "需要查询的教师邮箱" Format(email)
// @Param   id     query     string  false  "需要查询教师ID"
// @Param   name     query     string  false  "需要查询教师用户名"
// @Success 200 {object} model.CommonResponse[model.GetUserResponse]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/admin/teacher/{action} [get]
func (t *Admin) GetTeacher(c *gin.Context) {
	b, err := c.GetRawData()
	if err != nil {
		apputils.ThrowError(c, err)
		return
	}
	var req model.QueryRequest
	json.Unmarshal(b, &req)

	tx := apputils.BuildQuerySQL(t.DB.Table("users"), &req, "teacher")
	if tx == nil {
		apputils.Throw(c, errhandle.ParamsError)
		return
	}
	var user model.User
	err = tx.First(&user).Error
	if err != nil {
		apputils.Throw(c, errhandle.UserNonExists)
		return
	}
	var res model.GetUserResponse
	apputils.IgnoreStructCopy(&res, &user, c.Param("action"))

	apputils.OK[model.GetUserResponse](c, res)
}

// 获取所有学生信息 godoc
// @Summary 获取所有学生信息
// @Schemes
// @Description 获取所有学生信息
// @Tags example
// @Accept json
// @Produce json
// @Param   action    path     string  false  "查询过滤器，如果没有默认查询所以信息"
// @Param   token     header    string  true   "登录返回的Token"
// @Success 200 {object} model.CommonResponse[model.GetUserResponse]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/admin/students/{action} [get]
func (t *Admin) GetStudents(c *gin.Context) {
	var students []*model.User

	err := t.DB.Table("users").
		Where("role = ?", "student").
		Find(&students).Error

	if err != nil {
		apputils.ThrowError(c, err)
		return
	}

	var studentRes []model.GetUserResponse

	for _, tea := range students {
		var student model.GetUserResponse
		apputils.IgnoreStructCopy(&student, &tea, c.Param("action"))
		studentRes = append(studentRes, student)
	}

	apputils.OK[[]model.GetUserResponse](c, studentRes)
}

// 获得指定学生信息 godoc
// @Summary 获取指定学生信息
// @Schemes
// @Description 获取指定学生信息
// @Tags example
// @Accept json
// @Produce json
// @Param   action    path     string  false  "查询过滤器，如果没有默认查询所以信息"
// @Param   token     header    string  true   "登录返回的Token"
// @Param   queryemail     query     string  false  "需要查询的学生邮箱" Format(email)
// @Param   id     query     string  false  "需要查询学生ID"
// @Param   name     query     string  false  "需要查询学生用户名"
// @Success 200 {object} model.CommonResponse[model.GetUserResponse]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/admin/student/{action} [get]
func (t *Admin) GetStudent(c *gin.Context) {
	b, err := c.GetRawData()
	if err != nil {
		apputils.ThrowError(c, err)
		return
	}
	var req model.QueryRequest
	json.Unmarshal(b, &req)

	tx := apputils.BuildQuerySQL(t.DB.Table("users"), &req, "student")
	if tx == nil {
		apputils.Throw(c, errhandle.ParamsError)
		return
	}
	var user model.User
	err = tx.First(&user).Error
	if err != nil {
		apputils.Throw(c, errhandle.UserNonExists)
		return
	}
	var res model.GetUserResponse
	apputils.IgnoreStructCopy(&res, &user, c.Param("action"))

	apputils.OK[model.GetUserResponse](c, res)
}

// 修改指定学生信息 godoc
// @Summary 修改指定学生信息
// @Schemes
// @Description 修改指定学生信息
// @Tags example
// @Accept json
// @Produce json
// @Param   action    path     string  false  "查询过滤器，如果没有默认查询所以信息"
// @Param   token     header    string  true   "登录返回的Token"
// @Param   queryemail     query     string  false  "需要查询的学生邮箱" Format(email)
// @Param   id     query     string  false  "需要查询学生ID"
// @Param   name     query     string  false  "需要查询学生用户名"
// @Param   email     query    string  false  "修改邮箱"  Format(email)
// @Param   realName     query    string  false  "修改真实姓名"
// @Param   userSchoollD     query    string  false  "修改学校ID"
// @Param   schoolCode     query    string  false  "修改学校代码"
// @Param   class     query    string  false  "修改班级"
// @Success 200 {object} model.CommonResponse[any]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/admin/student [patch]
func (t *Admin) ModifyStudent(c *gin.Context) {
	b, err := c.GetRawData()
	if err != nil {
		apputils.ThrowError(c, err)
		return
	}
	var req model.AdminModifyRequest
	json.Unmarshal(b, &req)

	tx := apputils.BuildQuerySQL(t.DB.Table("users"), &req.QueryRequest, "student")
	if tx == nil {
		apputils.Throw(c, errhandle.ParamsError)
		return
	}
	var user model.User
	err = tx.First(&user).Error
	if err != nil {
		apputils.Throw(c, errhandle.UserNonExists)
		return
	}
	apputils.IgnoreStructCopy(&user, &req, "")

	tx.Save(&user)
	apputils.OK[any](c, nil)
}

// 修改指定教师学生信息 godoc
// @Summary 修改指定教师学生信息
// @Schemes
// @Description 修改指定教师学生信息
// @Tags example
// @Accept json
// @Produce json
// @Param   action    path     string  false  "查询过滤器，如果没有默认查询所以信息"
// @Param   token     header    string  true   "登录返回的Token"
// @Param   queryemail     query     string  false  "需要查询的学生邮箱" Format(email)
// @Param   id     query     string  false  "需要查询学生ID"
// @Param   name     query     string  false  "需要查询学生用户名"
// @Param   email     query    string  false  "修改邮箱"  Format(email)
// @Param   realName     query    string  false  "修改真实姓名"
// @Param   userSchoollD     query    string  false  "修改学校ID"
// @Param   schoolCode     query    string  false  "修改学校代码"
// @Param   class     query    string  false  "修改班级"
// @Success 200 {object} model.CommonResponse[any]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/admin/teacher [patch]
func (t *Admin) ModifyTeacher(c *gin.Context) {
	b, err := c.GetRawData()
	if err != nil {
		apputils.ThrowError(c, err)
		return
	}
	var req model.AdminModifyRequest
	json.Unmarshal(b, &req)

	tx := apputils.BuildQuerySQL(t.DB.Table("users"), &req.QueryRequest, "teacher")
	if tx == nil {
		apputils.Throw(c, errhandle.ParamsError)
		return
	}
	var user model.User
	err = tx.First(&user).Error
	if err != nil {
		apputils.Throw(c, errhandle.UserNonExists)
		return
	}
	apputils.IgnoreStructCopy(&user, &req, "")

	tx.Save(&user)
	apputils.OK[any](c, nil)
}

// 修改指定管理员信息 godoc
// @Summary 修改指定管理员信息
// @Schemes
// @Description 修改指定管理员信息
// @Tags example
// @Accept json
// @Produce json
// @Param   action    path     string  false  "查询过滤器，如果没有默认查询所以信息"
// @Param   token     header    string  true   "登录返回的Token"
// @Param   queryemail     query     string  false  "需要查询的学生邮箱" Format(email)
// @Param   id     query     string  false  "需要查询学生ID"
// @Param   name     query     string  false  "需要查询学生用户名"
// @Param   email     query    string  false  "修改邮箱"  Format(email)
// @Param   realName     query    string  false  "修改真实姓名"
// @Param   userSchoollD     query    string  false  "修改学校ID"
// @Param   schoolCode     query    string  false  "修改学校代码"
// @Param   class     query    string  false  "修改班级"
// @Success 200 {object} model.CommonResponse[any]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/admin/admin [patch]
func (t *Admin) ModifyAdmin(c *gin.Context) {
	b, err := c.GetRawData()
	if err != nil {
		apputils.ThrowError(c, err)
		return
	}
	var req model.AdminModifyRequest
	json.Unmarshal(b, &req)

	tx := apputils.BuildQuerySQL(t.DB.Table("users"), &req.QueryRequest, "admin")
	if tx == nil {
		apputils.Throw(c, errhandle.ParamsError)
		return
	}
	var user model.User
	err = tx.First(&user).Error
	if err != nil {
		apputils.Throw(c, errhandle.UserNonExists)
		return
	}
	apputils.IgnoreStructCopy(&user, &req, "")

	tx.Save(&user)
	apputils.OK[any](c, nil)
}

// 修改指定教师密码 godoc
// @Summary 修改指定教师密码
// @Schemes
// @Description 修改指定教师密码
// @Tags example
// @Accept json
// @Produce json
// @Param   action    path     string  false  "查询过滤器，如果没有默认查询所以信息"
// @Param   token     header    string  true   "登录返回的Token"
// @Param   queryemail     query     string  false  "需要查询的学生邮箱" Format(email)
// @Param   id     query     string  false  "需要查询学生ID"
// @Param   name     query     string  false  "需要查询学生用户名"
// @Param   password     query     string  false  "新密码"
// @Success 200 {object} model.CommonResponse[any]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/admin/teacher/password [patch]
func (t *Admin) ModifyTeacherPassword(c *gin.Context) {
	b, err := c.GetRawData()
	if err != nil {
		apputils.ThrowError(c, err)
		return
	}
	var req model.AdminModifyPasswordRequest
	json.Unmarshal(b, &req)

	tx := apputils.BuildQuerySQL(t.DB.Table("users"), &req.QueryRequest, "teacher")
	if tx == nil {
		apputils.Throw(c, errhandle.ParamsError)
		return
	}
	if req.Password == "" {
		apputils.Throw(c, errhandle.ParamsError)
		return
	}
	hashed, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		apputils.ThrowError(c, err)
		return
	}
	err = tx.Update("password", string(hashed)).Error
	if err != nil {
		apputils.Throw(c, errhandle.UserNonExists)
		return
	}
	apputils.OK[any](c, nil)
}

// 修改指定学生密码 godoc
// @Summary 修改指定学生密码
// @Schemes
// @Description 修改指定学生密码
// @Tags example
// @Accept json
// @Produce json
// @Param   action    path     string  false  "查询过滤器，如果没有默认查询所以信息"
// @Param   token     header    string  true   "登录返回的Token"
// @Param   queryemail     query     string  false  "需要查询的学生邮箱" Format(email)
// @Param   id     query     string  false  "需要查询学生ID"
// @Param   name     query     string  false  "需要查询学生用户名"
// @Param   password     query     string  false  "新密码"
// @Success 200 {object} model.CommonResponse[any]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/admin/student/password [patch]
func (t *Admin) ModifyStudentPassword(c *gin.Context) {
	b, err := c.GetRawData()
	if err != nil {
		apputils.ThrowError(c, err)
		return
	}
	var req model.AdminModifyPasswordRequest
	json.Unmarshal(b, &req)

	tx := apputils.BuildQuerySQL(t.DB.Table("users"), &req.QueryRequest, "student")
	if tx == nil {
		apputils.Throw(c, errhandle.ParamsError)
		return
	}
	if req.Password == "" {
		apputils.Throw(c, errhandle.ParamsError)
		return
	}
	hashed, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		apputils.ThrowError(c, err)
		return
	}
	err = tx.Update("password", string(hashed)).Error
	if err != nil {
		apputils.Throw(c, errhandle.UserNonExists)
		return
	}
	apputils.OK[any](c, nil)
}

// 修改指定管理员密码 godoc
// @Summary 修改指定管理员密码
// @Schemes
// @Description 修改指定管理员密码
// @Tags example
// @Accept json
// @Produce json
// @Param   action    path     string  false  "查询过滤器，如果没有默认查询所以信息"
// @Param   token     header    string  true   "登录返回的Token"
// @Param   queryemail     query     string  false  "需要查询的学生邮箱" Format(email)
// @Param   id     query     string  false  "需要查询学生ID"
// @Param   name     query     string  false  "需要查询学生用户名"
// @Success 200 {object} model.CommonResponse[any]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/admin/admin/password [patch]
func (t *Admin) ModifyAdminPassword(c *gin.Context) {
	b, err := c.GetRawData()
	if err != nil {
		apputils.ThrowError(c, err)
		return
	}
	var req model.AdminModifyPasswordRequest
	json.Unmarshal(b, &req)

	tx := apputils.BuildQuerySQL(t.DB.Table("users"), &req.QueryRequest, "admin")
	if tx == nil {
		apputils.Throw(c, errhandle.ParamsError)
		return
	}
	if req.Password == "" {
		apputils.Throw(c, errhandle.ParamsError)
		return
	}
	hashed, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		apputils.ThrowError(c, err)
		return
	}
	err = tx.Update("password", string(hashed)).Error
	if err != nil {
		apputils.Throw(c, errhandle.UserNonExists)
		return
	}
	apputils.OK[any](c, nil)
}

// 删除指定教师 godoc
// @Summary 删除指定教师
// @Schemes
// @Description 删除指定教师
// @Tags example
// @Accept json
// @Produce json
// @Param   action    path     string  false  "查询过滤器，如果没有默认查询所以信息"
// @Param   token     header    string  true   "登录返回的Token"
// @Param   queryemail     query     string  false  "需要查询的学生邮箱" Format(email)
// @Param   id     query     string  false  "需要查询学生ID"
// @Param   name     query     string  false  "需要查询学生用户名"
// @Success 200 {object} model.CommonResponse[any]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/admin/teacher [delete]
func (t *Admin) DeleteTeacher(c *gin.Context) {
	b, err := c.GetRawData()
	if err != nil {
		apputils.ThrowError(c, err)
		return
	}
	var req model.QueryRequest
	json.Unmarshal(b, &req)

	tx := apputils.BuildQuerySQL(t.DB.Table("users"), &req, "teacher")
	if tx == nil {
		apputils.Throw(c, errhandle.ParamsError)
		return
	}
	var user model.User
	err = tx.Delete(&user).Error
	if err != nil {
		apputils.Throw(c, errhandle.UserNonExists)
		return
	}
	apputils.OK[any](c, nil)
}

// 删除指定学生 godoc
// @Summary 删除指定学生
// @Schemes
// @Description 删除指定学生
// @Tags example
// @Accept json
// @Produce json
// @Param   action    path     string  false  "查询过滤器，如果没有默认查询所以信息"
// @Param   token     header    string  true   "登录返回的Token"
// @Param   queryemail     query     string  false  "需要查询的学生邮箱" Format(email)
// @Param   id     query     string  false  "需要查询学生ID"
// @Param   name     query     string  false  "需要查询学生用户名"
// @Success 200 {object} model.CommonResponse[any]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/admin/student [delete]
func (t *Admin) DeleteStudent(c *gin.Context) {
	b, err := c.GetRawData()
	if err != nil {
		apputils.ThrowError(c, err)
		return
	}
	var req model.QueryRequest
	json.Unmarshal(b, &req)

	tx := apputils.BuildQuerySQL(t.DB.Table("users"), &req, "student")
	if tx == nil {
		apputils.Throw(c, errhandle.ParamsError)
		return
	}
	var user model.User
	err = tx.Delete(&user).Error
	if err != nil {
		apputils.Throw(c, errhandle.UserNonExists)
		return
	}
	apputils.OK[any](c, nil)
}

// 删除指定管理员 godoc
// @Summary 删除指定管理员
// @Schemes
// @Description 删除指定管理员
// @Tags example
// @Accept json
// @Produce json
// @Param   action    path     string  false  "查询过滤器，如果没有默认查询所以信息"
// @Param   token     header    string  true   "登录返回的Token"
// @Param   queryemail     query     string  false  "需要查询的学生邮箱" Format(email)
// @Param   id     query     string  false  "需要查询学生ID"
// @Param   name     query     string  false  "需要查询学生用户名"
// @Success 200 {object} model.CommonResponse[any]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/admin/admin [delete]
func (t *Admin) DeleteAdmin(c *gin.Context) {
	b, err := c.GetRawData()
	if err != nil {
		apputils.ThrowError(c, err)
		return
	}
	var req model.QueryRequest
	json.Unmarshal(b, &req)

	tx := apputils.BuildQuerySQL(t.DB.Table("users"), &req, "admin")
	if tx == nil {
		apputils.Throw(c, errhandle.ParamsError)
		return
	}
	var user model.User
	err = tx.Delete(&user).Error
	if err != nil {
		apputils.Throw(c, errhandle.UserNonExists)
		return
	}
	apputils.OK[any](c, nil)
}

func (t *Admin) RegistserStudent(c *gin.Context) {

}

func (t *Admin) RegistserTeacher(c *gin.Context) {

}
