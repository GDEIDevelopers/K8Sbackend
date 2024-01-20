package student

import (
	"encoding/json"

	"github.com/GDEIDevelopers/K8Sbackend/app/apputils"
	"github.com/GDEIDevelopers/K8Sbackend/pkg/errhandle"
	"github.com/GDEIDevelopers/K8Sbackend/pkg/model"
	"github.com/gin-gonic/gin"
)

// 学生加入班级 godoc
// @Summary 学生加入班级
// @Schemes
// @Description 学生加入班级
// @Tags classStudent
// @Accept json
// @Produce json
// @Param   token     header    string  true   "登录返回的Token"
// @Param   classname    query    string  false "班级名称"
// @Success 200 {object} model.CommonResponse[any]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/student/class/join [post]
func (t *Student) Join(c *gin.Context) {
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

	err = t.Class.StudentJoin(info.UserID, req.ClassName)
	if err != nil {
		apputils.ThrowError(c, err)
		return
	}

	apputils.OK[any](c, nil)
}

// 学生离开班级 godoc
// @Summary 学生离开班级
// @Schemes
// @Description 学生离开班级
// @Tags classStudent
// @Accept json
// @Produce json
// @Param   token     header    string  true   "登录返回的Token"
// @Success 200 {object} model.CommonResponse[any]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/student/class/leave [delete]
func (t *Student) Leave(c *gin.Context) {
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

	err = t.Class.StudentLeave(info.UserID)
	if err != nil {
		apputils.ThrowError(c, err)
		return
	}

	apputils.OK[any](c, nil)
}

// 学生修改班级 godoc
// @Summary 学生修改班级
// @Schemes
// @Description 学生修改班级
// @Tags classStudent
// @Accept json
// @Produce json
// @Param   token     header    string  true   "登录返回的Token"
// @Param   classname    query    string  false "班级名称"
// @Success 200 {object} model.CommonResponse[any]
// @Failure 400  {object} model.CommonResponse[any]
// @Router /authrequired/teacher/class [patch]
func (t *Student) ClassModify(c *gin.Context) {
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
	classid, ok := t.Class.GetClassIDByName(req.ClassName)
	if !ok {
		apputils.Throw(c, errhandle.ClassNotFound)
		return
	}
	err = t.DB.Table("users").
		Where("id = ?", info.UserID).
		Update("class = ?", classid).Error
	if err != nil {
		apputils.ThrowError(c, err)
		return
	}
	apputils.OK[any](c, nil)
}
