package admin

import (
	"encoding/json"

	"github.com/GDEIDevelopers/K8Sbackend/app/apputils"
	"github.com/GDEIDevelopers/K8Sbackend/pkg/errhandle"
	"github.com/GDEIDevelopers/K8Sbackend/pkg/model"
	"github.com/gin-gonic/gin"
)

var _ = &model.CommonResponse[any]{}

// 管理员添加班级 godoc
// @Summary 管理员添加班级
// @Schemes
// @Description 管理员添加班级
// @Tags classAdmin
// @Accept json
// @Produce json
// @Param   token     header    string  true   "登录返回的Token"
// @Param   classname     query     string  true  "新班级名称"
// @Success 200 {object} model.CommonResponse[model.AddClassResponse]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/admin/class/new [post]
func (t *Admin) AddClass(c *gin.Context) {
	var req model.CommonClassRequest

	b, err := c.GetRawData()
	if err != nil {
		apputils.ThrowError(c, err)
		return
	}
	json.Unmarshal(b, &req)

	if req.ClassName == "" {
		apputils.Throw(c, errhandle.ParamsError)
		return
	}

	id, err := t.Class.AddClass(req.ClassName)
	if err != nil {
		apputils.ThrowError(c, err)
		return
	}

	apputils.OK(c, &model.AddClassResponse{
		ClassID: id,
	})
}

// 管理员删除班级 godoc
// @Summary 管理员删除班级
// @Schemes
// @Description 管理员删除班级
// @Tags classAdmin
// @Accept json
// @Produce json
// @Param   token     header    string  true   "登录返回的Token"
// @Param   classname     query     string  true  "班级名称"
// @Success 200 {object} model.CommonResponse[any]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/admin/class [delete]
func (t *Admin) RemoveClass(c *gin.Context) {
	var req model.CommonClassRequest

	b, err := c.GetRawData()
	if err != nil {
		apputils.ThrowError(c, err)
		return
	}
	json.Unmarshal(b, &req)

	if req.ClassName == "" {
		apputils.Throw(c, errhandle.ParamsError)
		return
	}

	classid, ok := t.Class.GetClassIDByName(req.ClassName)
	if !ok {
		apputils.Throw(c, errhandle.ClassError)
		return
	}
	err = t.Class.RemoveClassByID(classid)
	if err != nil {
		apputils.ThrowError(c, err)
		return
	}

	apputils.OK[any](c, nil)
}

// 获取所有班级 godoc
// @Summary 获取所有班级
// @Schemes
// @Description 获取所有班级
// @Tags auth
// @Accept json
// @Produce json
// @Param   token     header    string  true   "登录返回的Token"
// @Success 200 {object} model.CommonResponse[[]model.GetClassResponse]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/classes [get]
func (t *Admin) ListClasses(c *gin.Context) {
	var res []*model.GetClassResponse

	err := t.DB.Table("classMap").
		Joins("LEFT JOIN class ON classMap.classid = class.classid").
		Scan(&res).Error
	if err != nil {
		apputils.ThrowError(c, err)
		return
	}

	apputils.OK(c, res)
}

// 管理员列出指定班级老师 godoc
// @Summary 管理员列出指定班级老师
// @Schemes
// @Description 管理员列出指定班级老师
// @Tags classAdmin
// @Accept json
// @Produce json
// @Param   token     header    string  true   "登录返回的Token"
// @Param   teacherid     query    int  false   "教师用户ID(可选)"
// @Param   classid     query    int  false   "班级ID(可选)"
// @Success 200 {object} model.CommonResponse[[]model.Class]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/admin/class/teachers [post]
func (t *Admin) ListClassTeacher(c *gin.Context) {
	b, err := c.GetRawData()
	if err != nil {
		apputils.ThrowError(c, err)
		return
	}
	var req model.ClassQueryRequest
	var res []*model.Class
	json.Unmarshal(b, &req)

	tx := apputils.BuildClassIndex(t.DB.Table("class"), &req)

	err = tx.Find(&res).Error
	if err != nil {
		apputils.ThrowError(c, err)
		return
	}

	apputils.OK(c, res)
}

// 管理员列出指定班级所有学生 godoc
// @Summary 管理员列出指定班级所有学生
// @Schemes
// @Description 管理员列出指定班级所有学生
// @Tags classAdmin
// @Accept json
// @Produce json
// @Param   token     header    string  true   "登录返回的Token"
// @Param   classname     query     string  true  "班级名称"
// @Success 200 {object} model.CommonResponse[[]model.GetUserResponse]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/admin/class/students [post]
func (t *Admin) ListClassStudent(c *gin.Context) {
	b, err := c.GetRawData()
	if err != nil {
		apputils.ThrowError(c, err)
		return
	}
	var req model.CommonClassRequest
	json.Unmarshal(b, &req)
	if req.ClassName == "" {
		apputils.Throw(c, errhandle.ParamsError)
		return
	}
	classid, ok := t.Class.GetClassIDByName(req.ClassName)
	if !ok {
		apputils.Throw(c, errhandle.ClassError)
		return
	}
	res := t.Class.GetStudents(classid)
	if res == nil {
		apputils.Throw(c, errhandle.ClassError)
		return
	}

	apputils.OK(c, res)
}

// 管理员列出指定教师所有学生 godoc
// @Summary 管理员列出指定教师所有学生
// @Schemes
// @Description 管理员列出指定教师所有学生
// @Tags classAdmin
// @Accept json
// @Produce json
// @Param   token     header    string  true   "登录返回的Token"
// @Param   teacherid     query    int  true   "教师用户ID"
// @Success 200 {object} model.CommonResponse[model.GetClassBelongsResponse]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/admin/teacher/students [post]
func (t *Admin) ListTeacherStudent(c *gin.Context) {
	b, err := c.GetRawData()
	if err != nil {
		apputils.ThrowError(c, err)
		return
	}
	var req model.GetTeacherStudentRequest
	json.Unmarshal(b, &req)

	res := t.Class.BelongsTo(req.TeacherID)
	if res == nil {
		apputils.Throw(c, errhandle.TeacherNotFound)
		return
	}

	apputils.OK(c, res)
}

// 管理员添加老师到班级 godoc
// @Summary 管理员添加老师到班级
// @Schemes
// @Description 管理员添加老师到班级
// @Tags classAdmin
// @Accept json
// @Produce json
// @Param   token     header    string  true   "登录返回的Token"
// @Param   teacherid     query     int  true  "教师ID"
// @Param   classname     query     string  true  "班级名称"
// @Success 200 {object} model.CommonResponse[any]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/admin/teacher/class/new [post]
func (t *Admin) AddTeacherToClass(c *gin.Context) {
	b, err := c.GetRawData()
	if err != nil {
		apputils.ThrowError(c, err)
		return
	}
	var req model.TeacherClassRequest
	json.Unmarshal(b, &req)

	err = t.Class.TeacherJoin(req.TeacherID, req.ClassName)
	if err != nil {
		apputils.Throw(c, errhandle.TeacherNotFound)
		return
	}
	apputils.OK[any](c, nil)
}

// 管理员将老师移除班级 godoc
// @Summary 管理员将老师移除班级
// @Schemes
// @Description 管理员将老师移除班级
// @Tags classAdmin
// @Accept json
// @Produce json
// @Param   token     header    string  true   "登录返回的Token"
// @Param   teacherid     query     int  true  "教师ID"
// @Param   classname     query     string  false  "班级名称(可选)"
// @Success 200 {object} model.CommonResponse[any]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/admin/teacher/class [delete]
func (t *Admin) RemoveTeacherFromClass(c *gin.Context) {
	b, err := c.GetRawData()
	if err != nil {
		apputils.ThrowError(c, err)
		return
	}
	var req model.TeacherClassRequest
	json.Unmarshal(b, &req)

	err = t.Class.TeacherLeave(req.TeacherID, req.ClassName)
	if err != nil {
		apputils.Throw(c, errhandle.TeacherNotFound)
		return
	}
	apputils.OK[any](c, nil)
}

// 管理员添加学生到指定班级 godoc
// @Summary 管理员添加学生到指定班级
// @Schemes
// @Description 管理员添加学生到指定班级
// @Tags classAdmin
// @Accept json
// @Produce json
// @Param   token     header    string  true   "登录返回的Token"
// @Param   studentid     query     int  true  "学生用户ID"
// @Param   classname     query     string  true  "新班级名称"
// @Success 200 {object} model.CommonResponse[any]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/admin/student/class/new [post]
func (t *Admin) AddStudentToClass(c *gin.Context) {
	b, err := c.GetRawData()
	if err != nil {
		apputils.ThrowError(c, err)
		return
	}
	var req model.StudentClassRequest
	json.Unmarshal(b, &req)

	err = t.Class.StudentJoin(req.StudentID, req.ClassName)
	if err != nil {
		apputils.Throw(c, errhandle.TeacherNotFound)
		return
	}
	apputils.OK[any](c, nil)
}

// 管理员移除学生班级 godoc
// @Summary 管理员移除学生班级
// @Schemes
// @Description 管理员移除学生班级
// @Tags classAdmin
// @Accept json
// @Produce json
// @Param   token     header    string  true   "登录返回的Token"
// @Param   studentid     query     int  true  "学生用户ID"
// @Success 200 {object} model.CommonResponse[any]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/admin/student/class [delete]
func (t *Admin) RemoveStudentFromClass(c *gin.Context) {
	b, err := c.GetRawData()
	if err != nil {
		apputils.ThrowError(c, err)
		return
	}
	var req model.StudentLeaveClassRequest
	json.Unmarshal(b, &req)

	err = t.Class.StudentLeave(req.StudentID)
	if err != nil {
		apputils.Throw(c, errhandle.TeacherNotFound)
		return
	}
	apputils.OK[any](c, nil)
}

// 管理员修改学生班级 godoc
// @Summary 管理员修改学生班级
// @Schemes
// @Description 管理员修改学生班级
// @Tags classAdmin
// @Accept json
// @Produce json
// @Param   token     header    string  true   "登录返回的Token"
// @Param   studentid     query     int  true  "学生用户ID"
// @Param   classname     query     string  true  "新班级名称"
// @Success 200 {object} model.CommonResponse[any]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/admin/student/class [patch]
func (t *Admin) ModifyStudentClass(c *gin.Context) {
	b, err := c.GetRawData()
	if err != nil {
		apputils.ThrowError(c, err)
		return
	}
	var req model.StudentClassRequest
	json.Unmarshal(b, &req)
	classid, ok := t.Class.GetClassIDByName(req.ClassName)
	if !ok {
		apputils.Throw(c, errhandle.ClassNotFound)
		return
	}
	err = t.DB.Table("users").
		Where("id = ?", req.StudentID).
		Update("class = ?", classid).Error
	if err != nil {
		apputils.ThrowError(c, err)
		return
	}
	apputils.OK[any](c, nil)
}
