package admin

import (
	"encoding/json"
	"net/mail"

	"github.com/GDEIDevelopers/K8Sbackend/app/apputils"
	"github.com/GDEIDevelopers/K8Sbackend/pkg/errhandle"
	"github.com/GDEIDevelopers/K8Sbackend/pkg/model"
	"github.com/GDEIDevelopers/K8Sbackend/pkg/snowflake"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// 获取所有教师信息 godoc
// @Summary 获取所有教师信息
// @Schemes
// @Description 获取所有教师信息
// @Tags admin
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
// @Tags admin
// @Accept json
// @Produce json
// @Param   action    path     string  false  "查询过滤器，如果没有默认查询所以信息"
// @Param   token     header    string  true   "登录返回的Token"
// @Param   queryemail     query     string  false  "需要查询的教师邮箱" Format(email)
// @Param   id     query     string  false  "需要查询教师ID"
// @Param   name     query     string  false  "需要查询教师用户名"
// @Param   queryRealname     query     string  false  "需要查询教师真实姓名"
// @Param   queryUserSchoollD     query     string  false  "需要查询教师学号"
// @Success 200 {object} model.CommonResponse[model.GetUserResponse]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/admin/teacher/{action} [post]
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
// @Tags admin
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

		if classname, ok := t.Class.GetClassNameByID(tea.Class); ok {
			student.Class = classname
		}
		studentRes = append(studentRes, student)
	}

	apputils.OK[[]model.GetUserResponse](c, studentRes)
}

// 获得指定学生信息 godoc
// @Summary 获取指定学生信息
// @Schemes
// @Description 获取指定学生信息
// @Tags admin
// @Accept json
// @Produce json
// @Param   action    path     string  false  "查询过滤器，如果没有默认查询所以信息"
// @Param   token     header    string  true   "登录返回的Token"
// @Param   queryemail     query     string  false  "需要查询的学生邮箱" Format(email)
// @Param   id     query     string  false  "需要查询学生ID"
// @Param   name     query     string  false  "需要查询学生用户名"
// @Param   queryRealname     query     string  false  "需要查询学生真实姓名"
// @Param   queryUserSchoollD     query     string  false  "需要查询学生学号"
// @Success 200 {object} model.CommonResponse[model.GetUserResponse]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/admin/student/{action} [post]
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
	if classname, ok := t.Class.GetClassNameByID(user.Class); ok {
		res.Class = classname
	}
	apputils.OK[model.GetUserResponse](c, res)
}

// 修改指定学生信息 godoc
// @Summary 修改指定学生信息
// @Schemes
// @Description 修改指定学生信息
// @Tags admin
// @Accept json
// @Produce json
// @Param   token     header    string  true   "登录返回的Token"
// @Param   id     query     string  false  "需要查询学生ID"
// @Param   name     query     string  false  "需要修改学生用户名"
// @Param   email     query    string  false  "修改邮箱"  Format(email)
// @Param   realName     query    string  false  "修改真实姓名"
// @Param   userSchoollD     query    string  false  "修改学号"
// @Param   schoolCode     query    string  false  "修改学校代码"
// @Param   class     query    string  false  "修改班级"
// @Param   sex     query    string  false  "修改性别"
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

	var user model.User
	err = t.DB.Table("users").
		Where("id = ? AND role = ?", req.UserID, "student").
		First(&user).Error
	if err != nil {
		apputils.Throw(c, errhandle.UserNonExists)
		return
	}
	apputils.IgnoreStructCopy(&user, &req, "")

	if req.Class != "" {
		classid, ok := t.Class.GetClassIDByName(req.Class)
		if !ok {
			apputils.Throw(c, errhandle.ClassNotFound)
			return
		}
		user.Class = classid
	}

	t.DB.Table("users").
		Where("id = ? AND role = ?", req.UserID, "student").
		Save(&user)
	apputils.OK[any](c, nil)
}

// 修改指定教师信息 godoc
// @Summary 修改指定教师信息
// @Schemes
// @Description 修改指定教师信息
// @Tags admin
// @Accept json
// @Produce json
// @Param   token     header    string  true   "登录返回的Token"
// @Param   id     query     string  false  "需要查询ID"
// @Param   name     query     string  false  "需要修改用户名"
// @Param   email     query    string  false  "修改邮箱"  Format(email)
// @Param   realName     query    string  false  "修改真实姓名"
// @Param   userSchoollD     query    string  false  "修改学号"
// @Param   schoolCode     query    string  false  "修改学校代码"
// @Param   class     query    string  false  "修改班级"
// @Param   sex     query    string  false  "修改性别"
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

	var user model.User
	err = t.DB.Table("users").
		Where("id = ? AND role = ?", req.UserID, "teacher").
		First(&user).Error
	if err != nil {
		apputils.Throw(c, errhandle.UserNonExists)
		return
	}
	apputils.IgnoreStructCopy(&user, &req, "")

	t.DB.Table("users").
		Where("id = ? AND role = ?", req.UserID, "teacher").
		Save(&user)
	apputils.OK[any](c, nil)
}

// 修改指定管理员信息 godoc
// @Summary 修改指定管理员信息
// @Schemes
// @Description 修改指定管理员信息
// @Tags admin
// @Accept json
// @Produce json
// @Param   token     header    string  true   "登录返回的Token"
// @Param   id     query     string  false  "需要查询学生ID"
// @Param   name     query     string  false  "需要修改的学生用户名"
// @Param   email     query    string  false  "修改邮箱"  Format(email)
// @Param   realName     query    string  false  "修改真实姓名"
// @Param   userSchoollD     query    string  false  "修改学号"
// @Param   schoolCode     query    string  false  "修改学校代码"
// @Param   class     query    string  false  "修改班级"
// @Param   sex     query    string  false  "修改性别"
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

	var user model.User
	err = t.DB.Table("users").
		Where("id = ? AND role = ?", req.UserID, "admin").
		First(&user).Error
	if err != nil {
		apputils.Throw(c, errhandle.UserNonExists)
		return
	}
	apputils.IgnoreStructCopy(&user, &req, "")

	t.DB.Table("users").
		Where("id = ? AND role = ?", req.UserID, "admin").
		Save(&user)
	apputils.OK[any](c, nil)
}

// 修改指定教师密码 godoc
// @Summary 修改指定教师密码
// @Schemes
// @Description 修改指定教师密码
// @Tags admin
// @Accept json
// @Produce json
// @Param   token     header    string  true   "登录返回的Token"
// @Param   id     query     string  false  "需要查询学生ID"
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
	err = t.DB.Table("users").
		Where("id = ? AND role = ?", req.UserID, "teacher").
		Update("password", string(hashed)).Error
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
// @Tags admin
// @Accept json
// @Produce json
// @Param   token     header    string  true   "登录返回的Token"
// @Param   id     query     string  false  "需要查询学生ID"
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
	err = t.DB.Table("users").
		Where("id = ? AND role = ?", req.UserID, "student").
		Update("password", string(hashed)).Error
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
// @Tags admin
// @Accept json
// @Produce json
// @Param   token     header    string  true   "登录返回的Token"
// @Param   id     query     string  false  "需要查询管理员ID"
// @Param   password     query     string  false  "新密码"
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
	err = t.DB.Table("users").
		Where("id = ? AND role = ?", req.UserID, "student").
		Update("password", string(hashed)).Error
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
// @Tags admin
// @Accept json
// @Produce json
// @Param   token     header    string  true   "登录返回的Token"
// @Param   id     query     string  false  "需要删除ID"
// @Success 200 {object} model.CommonResponse[any]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/admin/teacher [delete]
func (t *Admin) DeleteTeacher(c *gin.Context) {
	b, err := c.GetRawData()
	if err != nil {
		apputils.ThrowError(c, err)
		return
	}
	var req model.UserIDOnlyRequest
	json.Unmarshal(b, &req)

	var user model.User
	err = t.DB.Table("users").
		Where("id = ? AND role = ?", req.UserID, "teacher").
		Delete(&user).Error
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
// @Tags admin
// @Accept json
// @Produce json
// @Param   token     header    string  true   "登录返回的Token"
// @Param   id     query     string  false  "需要查询ID"
// @Success 200 {object} model.CommonResponse[any]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/admin/student [delete]
func (t *Admin) DeleteStudent(c *gin.Context) {
	b, err := c.GetRawData()
	if err != nil {
		apputils.ThrowError(c, err)
		return
	}
	var req model.UserIDOnlyRequest
	json.Unmarshal(b, &req)

	var user model.User
	err = t.DB.Table("users").
		Where("id = ? AND role = ?", req.UserID, "student").
		Delete(&user).Error
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
// @Tags admin
// @Accept json
// @Produce json
// @Param   token     header    string  true   "登录返回的Token"
// @Param   id     query     string  false  "需要查询ID"
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

	var user model.User
	err = t.DB.Table("users").
		Where("id = ? AND role = ?", req.UserID, "admin").
		Delete(&user).Error
	if err != nil {
		apputils.Throw(c, errhandle.UserNonExists)
		return
	}
	apputils.OK[any](c, nil)
}

// 注册学生 godoc
// @Summary 注册学生
// @Schemes
// @Description 注册学生
// @Tags admin
// @Accept json
// @Produce json
// @Param   token     header    string  true   "登录返回的Token"
// @Param   name     query     string  true  "新用户用户名"
// @Param   email     query    string  true  "新用户邮箱"  Format(email)
// @Param   realName     query    string  true  "新用户真实姓名"
// @Param   userSchoollD     query    string  true  "新用户学号"
// @Param   schoolCode     query    string  true  "新用户学校代码"
// @Param   class     query    string  true  "新用户班级"
// @Param   sex     query    string  true  "新用户性别"
// @Success 200 {object} model.CommonResponse[any]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/admin/student/new [post]
func (t *Admin) RegisterStudent(c *gin.Context) {
	b, err := c.GetRawData()
	if err != nil {
		apputils.ThrowError(c, err)
		return
	}
	var req model.RegisterUserRequest
	json.Unmarshal(b, &req)

	if req.Class == "" {
		apputils.Throw(c, errhandle.ClassError)
		return
	}
	if _, err := mail.ParseAddress(req.Email); err != nil {
		apputils.Throw(c, errhandle.EmailFormatError)
		return
	}
	if req.Sex != "男" && req.Sex != "女" {
		apputils.Throw(c, errhandle.SexError)
		return
	}
	if req.SchoolCode == "" || req.UserSchoollD == "" {
		apputils.Throw(c, errhandle.SchoolError)
		return
	}
	if !apputils.IsValidPassword(req.Password) {
		apputils.Throw(c, errhandle.PasswordTooShort)
		return
	}

	classid, ok := t.Class.GetClassIDByName(req.Class)
	if !ok {
		apputils.Throw(c, errhandle.ClassNotFound)
		return
	}

	hashed, _ := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
	var found model.User
	col := t.DB.Table("users").FirstOrCreate(&found, model.User{
		ID:           snowflake.ID(),
		Role:         "student",
		SchoolCode:   req.SchoolCode,
		UserSchoollD: req.UserSchoollD,
		Name:         req.Name,
		RealName:     req.RealName,
		Sex:          req.Sex,
		Class:        classid,
		Password:     string(hashed),
		Email:        req.Email,
	})
	// this shouldn't happen
	if col.RowsAffected == 0 {
		apputils.Throw(c, errhandle.InnerError)
		return
	}

	apputils.OK[any](c, nil)
}

// 注册教师 godoc
// @Summary 注册教师
// @Schemes
// @Description 注册教师
// @Tags admin
// @Accept json
// @Produce json
// @Param   token     header    string  true   "登录返回的Token"
// @Param   name     query     string  true  "新用户用户名"
// @Param   email     query    string  true  "新用户邮箱"  Format(email)
// @Param   realName     query    string  true  "新用户真实姓名"
// @Param   sex     query    string  true  "新用户性别"
// @Success 200 {object} model.CommonResponse[any]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/admin/teacher/new [post]
func (t *Admin) RegisterTeacher(c *gin.Context) {
	b, err := c.GetRawData()
	if err != nil {
		apputils.ThrowError(c, err)
		return
	}
	var req model.RegisterUserRequest
	json.Unmarshal(b, &req)

	if _, err := mail.ParseAddress(req.Email); err != nil {
		apputils.Throw(c, errhandle.EmailFormatError)
		return
	}
	if req.Sex != "男" && req.Sex != "女" {
		apputils.Throw(c, errhandle.SexError)
		return
	}
	if !apputils.IsValidPassword(req.Password) {
		apputils.Throw(c, errhandle.PasswordTooShort)
		return
	}

	hashed, _ := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
	var found model.User
	col := t.DB.Table("users").FirstOrCreate(&found, model.User{
		ID:           snowflake.ID(),
		Role:         "teacher",
		SchoolCode:   req.SchoolCode,
		UserSchoollD: req.UserSchoollD,
		Name:         req.Name,
		RealName:     req.RealName,
		Sex:          req.Sex,
		Class:        0,
		Password:     string(hashed),
		Email:        req.Email,
	})
	// this shouldn't happen
	if col.RowsAffected == 0 {
		apputils.Throw(c, errhandle.InnerError)
		return
	}

	apputils.OK[any](c, nil)
}

// 注册管理员 godoc
// @Summary 注册管理员
// @Schemes
// @Description 注册管理员
// @Tags admin
// @Accept json
// @Produce json
// @Param   token     header    string  true   "登录返回的Token"
// @Param   name     query     string  true  "新用户用户名"
// @Param   email     query    string  true  "新用户邮箱"  Format(email)
// @Param   realName     query    string  true  "新用户真实姓名"
// @Param   sex     query    string  true  "新用户性别"
// @Success 200 {object} model.CommonResponse[any]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/admin/admin/new [post]
func (t *Admin) RegisterAdmin(c *gin.Context) {
	b, err := c.GetRawData()
	if err != nil {
		apputils.ThrowError(c, err)
		return
	}
	var req model.RegisterUserRequest
	json.Unmarshal(b, &req)

	if _, err := mail.ParseAddress(req.Email); err != nil {
		apputils.Throw(c, errhandle.EmailFormatError)
		return
	}
	if req.Sex != "男" && req.Sex != "女" {
		apputils.Throw(c, errhandle.SexError)
		return
	}
	if !apputils.IsValidPassword(req.Password) {
		apputils.Throw(c, errhandle.PasswordTooShort)
		return
	}

	hashed, _ := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
	var found model.User
	col := t.DB.Table("users").FirstOrCreate(&found, model.User{
		ID:           snowflake.ID(),
		Role:         "admin",
		SchoolCode:   req.SchoolCode,
		UserSchoollD: req.UserSchoollD,
		Name:         req.Name,
		RealName:     req.RealName,
		Sex:          req.Sex,
		Class:        0,
		Password:     string(hashed),
		Email:        req.Email,
	})
	// this shouldn't happen
	if col.RowsAffected == 0 {
		apputils.Throw(c, errhandle.InnerError)
		return
	}

	apputils.OK[any](c, nil)
}
