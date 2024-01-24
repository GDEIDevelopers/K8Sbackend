package teacher

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

// 修改教师信息 godoc
// @Summary 修改教师相关信息
// @Schemes
// @Description 修改教师相关信息
// @Tags teacher
// @Accept json
// @Produce json
// @Param   action    path     string  false  "查询过滤器，如果没有默认查询所以信息"
// @Param   token     header   string  true   "登录返回的Token"
// @Param   email     query    string  false  "修改邮箱"  Format(email)
// @Param   realName     query    string  false  "修改真实姓名"
// @Param   userSchoollD     query    string  false  "修改学号"
// @Param   schoolCode     query    string  false  "修改学校代码"
// @Param   class     query    string  false  "修改班级"
// @Param   sex     query    string  false  "修改性别"
// @Success 200 {object} model.CommonResponse[any]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/teacher [patch]
func (t *Teacher) Modify(c *gin.Context) {
	userinfo, ok := c.Get("info")
	if !ok {
		apputils.Throw(c, errhandle.InnerError)
		return
	}
	info := userinfo.(*model.UserInfo)

	b, err := c.GetRawData()
	if err != nil {
		apputils.ThrowError(c, err)
		return
	}
	var user model.User
	var modifyReq model.ModifyUserRequest
	json.Unmarshal(b, &modifyReq)

	err = t.DB.Table("users").
		Where("id = ?", info.UserID).
		First(&user).Error
	if err != nil {
		apputils.ThrowError(c, err)
		return
	}

	apputils.IgnoreStructCopy(&user, &modifyReq, "")

	err = t.DB.Table("users").
		Where("id = ?", info.UserID).
		Save(&user).Error
	if err != nil {
		apputils.ThrowError(c, err)
		return
	}
	apputils.OK[any](c, nil)
}

// 修改教师密码 godoc
// @Summary 修改教师密码
// @Schemes
// @Description 修改教师密码
// @Tags teacher
// @Accept json
// @Produce json
// @Param   action    path     string  false  "查询过滤器，如果没有默认查询所以信息"  Format(email)
// @Param   token     header    string  true   "登录返回的Token"
// @Param   password     query    string  false "需要修改的密码"
// @Success 200 {object} model.CommonResponse[any]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/teacher/password [patch]
func (t *Teacher) ModifyPassword(c *gin.Context) {
	userinfo, ok := c.Get("info")
	if !ok {
		apputils.Throw(c, errhandle.InnerError)
		return
	}
	info := userinfo.(*model.UserInfo)

	b, err := c.GetRawData()
	if err != nil {
		apputils.ThrowError(c, err)
		return
	}
	var modifyReq model.ModifyUserPasswordRequest
	json.Unmarshal(b, &modifyReq)

	if modifyReq.Password == "" {
		apputils.Throw(c, errhandle.ParamsError)
		return
	}
	hashed, err := bcrypt.GenerateFromPassword(
		[]byte(modifyReq.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		apputils.ThrowError(c, err)
		return
	}
	err = t.DB.Table("users").
		Where("id = ?", info.UserID).
		Update("password", string(hashed)).Error
	if err != nil {
		apputils.ThrowError(c, err)
		return
	}
	apputils.OK[any](c, nil)
}

// 添加/注册一位学生 godoc
// @Summary 添加/注册一个学生
// @Schemes
// @Description 添加/注册一个学生
// @Tags teacher
// @Accept json
// @Produce json
// @Param   token     query    string  true   "登录返回的Token"
// @Param   name     query     string  true  "新用户用户名"
// @Param   email     query    string  true  "新用户邮箱"  Format(email)
// @Param   realName     query    string  true  "新用户真实姓名"
// @Param   userSchoollD     query    string  true  "新用户学号"
// @Param   schoolCode     query    string  true  "新用户学校代码"
// @Param   class     query    string  true  "新用户班级"
// @Param   sex     query    string  true  "新用户性别"
// @Success 200 {object} model.CommonResponse[any]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/teacher/student/new [post]
func (t *Teacher) RegisterStudent(c *gin.Context) {
	userinfo, ok := c.Get("info")
	if !ok {
		apputils.Throw(c, errhandle.InnerError)
		return
	}
	info := userinfo.(*model.UserInfo)

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
	myclasses, err := t.Class.GetTeacherClass(info.UserID)
	if err != nil {
		apputils.Throw(c, errhandle.TeacherNotJoinClass)
		return
	}
	foundClass := false
	for _, myclass := range myclasses {
		if myclass == classid {
			foundClass = true
			break
		}
	}
	if !foundClass {
		apputils.Throw(c, errhandle.PermissionDenied)
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

// 获得指定学生信息 godoc
// @Summary 获取指定学生信息
// @Schemes
// @Description 获取指定学生信息
// @Tags teacher
// @Accept json
// @Produce json
// @Param   token     header    string  true   "登录返回的Token"
// @Param   queryemail     query     string  false  "需要查询的学生邮箱" Format(email)
// @Param   id     query     string  false  "需要查询学生ID"
// @Param   name     query     string  false  "需要查询学生用户名"
// @Param   queryRealname     query     string  false  "需要查询学生真实姓名"
// @Param   queryUserSchoollD     query     string  false  "需要查询学生学号"
// @Success 200 {object} model.CommonResponse[[]model.GetUserResponse]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/teacher/student [post]
func (t *Teacher) GetStudent(c *gin.Context) {
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
	var user []*model.User
	err = tx.Find(&user).Error
	if err != nil {
		apputils.Throw(c, errhandle.UserNonExists)
		return
	}
	var studentRes []*model.GetUserResponse

	for _, tea := range user {
		var student model.GetUserResponse
		apputils.IgnoreStructCopy(&student, &tea, "")

		if classname, ok := t.Class.GetClassNameByID(tea.Class); ok {
			student.Class = classname
		}
		studentRes = append(studentRes, &student)
	}

	apputils.OK(c, studentRes)
}
