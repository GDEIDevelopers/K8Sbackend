package teacher

import (
	"encoding/json"

	"github.com/GDEIDevelopers/K8Sbackend/app/apputils"
	"github.com/GDEIDevelopers/K8Sbackend/pkg/errhandle"
	"github.com/GDEIDevelopers/K8Sbackend/pkg/model"
	"github.com/gin-gonic/gin"
)

// 教师加入班级 godoc
// @Summary 教师加入班级
// @Schemes
// @Description 教师加入班级
// @Tags classTeacher
// @Accept json
// @Produce json
// @Param   token     header    string  true   "登录返回的Token"
// @Param   classname    query    string  false "班级名称"
// @Success 200 {object} model.CommonResponse[any]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/teacher/class/join [post]
func (t *Teacher) Join(c *gin.Context) {
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
	var req model.CommonClassRequest
	json.Unmarshal(b, &req)

	if req.ClassName == "" {
		apputils.Throw(c, errhandle.ParamsError)
		return
	}

	err = t.Class.TeacherJoin(info.UserID, req.ClassName)
	if err != nil {
		apputils.ThrowError(c, err)
		return
	}

	apputils.OK[any](c, nil)
}

// 教师离开班级 godoc
// @Summary 教师离开班级
// @Schemes
// @Description 教师离开班级
// @Tags classTeacher
// @Accept json
// @Produce json
// @Param   token     header    string  true   "登录返回的Token"
// @Param   classname    query    string  false "班级名称"
// @Success 200 {object} model.CommonResponse[any]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/teacher/class/leave [delete]
func (t *Teacher) Leave(c *gin.Context) {
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
	var req model.CommonClassRequest
	json.Unmarshal(b, &req)

	if req.ClassName == "" {
		apputils.Throw(c, errhandle.ParamsError)
		return
	}

	err = t.Class.TeacherLeave(info.UserID, req.ClassName)
	if err != nil {
		apputils.ThrowError(c, err)
		return
	}

	apputils.OK[any](c, nil)
}

// 列出班级所有学生 godoc
// @Summary 列出班级所有学生
// @Schemes
// @Description 列出班级所有学生
// @Tags classTeacher
// @Accept json
// @Produce json
// @Param   classname    path     string  false  "班级名称(可选), 只显示网络工程B的学生: /authrequired/teacher/class/students/21网络工程B"
// @Param   token     header    string  true   "登录返回的Token"
// @Success 200 {object} model.CommonResponse[model.GetClassBelongsResponse]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/teacher/class/students/{classname} [get]
func (t *Teacher) ListStudents(c *gin.Context) {
	userinfo, ok := c.Get("info")
	if !ok {
		apputils.Throw(c, errhandle.InnerError)
		return
	}
	info := userinfo.(*model.UserInfo)

	apputils.OK(c,
		t.Class.BelongsTo(
			info.UserID,
			c.Param("classname"),
		),
	)
}

// 添加学生到指定班级 godoc
// @Summary 添加学生到指定班级
// @Schemes
// @Description 添加学生到指定班级
// @Tags classTeacher
// @Accept json
// @Produce json
// @Param   token     header    string  true   "登录返回的Token"
// @Param   studentid     query    []int64  true "学生用户ID"
// @Param   classname     query   string true "班级名称"
// @Success 200 {object} model.CommonResponse[any]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/teacher/class/students/join [post]
func (t *Teacher) AddStudents(c *gin.Context) {
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
	var req model.TeacherAddStudentRequest
	json.Unmarshal(b, &req)

	if req.ClassName == "" || len(req.StudentIDs) == 0 {
		apputils.Throw(c, errhandle.ParamsError)
		return
	}
	classes, err := t.Class.GetTeacherClass(info.UserID)
	if err != nil {
		apputils.Throw(c, errhandle.TeacherNotJoinClass)
		return
	}
	for _, studentid := range req.StudentIDs {
		if err := t.Class.StudentJoin(studentid, req.ClassName, classes...); err != nil {
			apputils.ThrowError(c, err)
			return
		}
	}

	apputils.OK[any](c, nil)
}

// 从指定班级移除学生 godoc
// @Summary 从指定班级移除学生
// @Schemes
// @Description 从指定班级移除学生
// @Tags classTeacher
// @Accept json
// @Produce json
// @Param   token     header    string  true   "登录返回的Token"
// @Param   studentid     query    []int64  true "学生用户ID"
// @Success 200 {object} model.CommonResponse[any]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/teacher/class/students/leave [delete]
func (t *Teacher) RemoveStudents(c *gin.Context) {
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
	var req model.TeacherRemoveStudentRequest
	json.Unmarshal(b, &req)

	if len(req.StudentIDs) == 0 {
		apputils.Throw(c, errhandle.ParamsError)
		return
	}

	classes, err := t.Class.GetTeacherClass(info.UserID)
	if err != nil {
		apputils.Throw(c, errhandle.TeacherNotJoinClass)
		return
	}

	for _, studentid := range req.StudentIDs {
		if err := t.Class.StudentLeave(studentid, classes...); err != nil {
			apputils.ThrowError(c, err)
			return
		}
	}

	apputils.OK[any](c, nil)
}

// 列出所有已加入的班级 godoc
// @Summary 列出所有已加入的班级
// @Schemes
// @Description 列出所有已加入的班级
// @Tags classTeacher
// @Accept json
// @Produce json
// @Param   token     header    string  true   "登录返回的Token"
// @Success 200 {object} model.CommonResponse[[]*model.Class]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/teacher/class [get]
func (t *Teacher) ListJoinedClass(c *gin.Context) {
	userinfo, ok := c.Get("info")
	if !ok {
		apputils.Throw(c, errhandle.InnerError)
		return
	}
	info := userinfo.(*model.UserInfo)
	var ret []*model.Class
	err := t.DB.Table("class").
		Where("teacherid = ?", info.UserID).
		Find(&ret).Error

	if err != nil {
		apputils.Throw(c, errhandle.TeacherNotJoinClass)
		return
	}

	apputils.OK(c, ret)
}
