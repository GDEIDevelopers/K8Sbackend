package apputils

import (
	"errors"
	"strings"

	"github.com/GDEIDevelopers/K8Sbackend/pkg/model"
	"github.com/GDEIDevelopers/K8Sbackend/pkg/snowflake"
	"gorm.io/gorm"
)

var (
	ErrClassNotExists = errors.New("class doesn't exist")
)

type Class struct {
	*ServerUtils
}

func NewClass(db *ServerUtils) *Class {
	return &Class{db}
}

func (c *Class) GetClassIDByName(name string) (int64, bool) {
	var class model.ClassMap
	err := c.DB.Table("classMap").
		Where("classname = ?", name).
		First(&class).Error
	if err != nil {
		return 0, false
	}
	return class.ClassID, true
}

func (c *Class) GetClassNameByID(id int64) (string, bool) {
	var class model.ClassMap
	err := c.DB.Table("classMap").
		Where("classid = ?", id).
		First(&class).Error
	if err != nil {
		return "", false
	}
	return class.Name, true
}

func (c *Class) GetStudents(id int64) (ret []*model.GetUserResponse) {
	var users []*model.User
	err := c.DB.Table("users").
		Where("class = ?", id).
		Find(&users).Error
	if err != nil {
		return
	}
	for _, user := range users {
		var res model.GetUserResponse
		IgnoreStructCopy(&res, user, "")
		ret = append(ret, &res)
	}
	return
}

func (c *Class) BelongsTo(teacherid int64, classname ...string) *model.GetClassBelongsResponse {
	var class []*model.Class

	filterName := func(name string) bool {
		if len(classname) > 0 && classname[0] != "" {
			return name == classname[0]
		}
		return true
	}

	err := c.DB.Table("class").
		Where("teacherid = ?", teacherid).
		Find(&class).Error
	if err != nil {
		return nil
	}
	var ret model.GetClassBelongsResponse
	for _, classes := range class {
		if name, ok := c.GetClassNameByID(classes.ClassID); ok && filterName(name) {
			current := &model.ClassWithStudent{
				ClassName: name,
				Studetns:  c.GetStudents(classes.ClassID),
			}
			ret.Classes = append(ret.Classes, current)
		}
	}
	return &ret
}

func (c *Class) AddClass(name string) (id int64, err error) {
	id = snowflake.ID()
	err = c.DB.Table("classMap").Create(&model.ClassMap{
		ClassID: id,
		Name:    name,
	}).Error
	return
}

func (c *Class) RemoveClassByID(id int64) (err error) {
	err = c.DB.Table("classMap").
		Where("classid = ?", id).
		Delete(&model.ClassMap{}).Error
	if err != nil {
		return
	}
	err = c.DB.Table("class").
		Where("classid = ?", id).
		Delete(&model.Class{}).Error
	if err != nil {
		return
	}
	err = c.DB.Table("users").
		Where("class = ?", id).
		Update("class", 0).Error

	return
}

func (c *Class) GetStudentClassID(studentid int64) (int64, error) {
	var usr model.User

	err := c.DB.Table("users").
		Where("id = ?", studentid).
		First(&usr).Error

	if err != nil {
		return 0, err
	}

	return usr.Class, nil
}

func (c *Class) StudentJoin(studentid int64, name string, filter ...int64) (err error) {
	classid, ok := c.GetClassIDByName(name)
	if !ok {
		return ErrClassNotExists
	}
	sql := c.DB.Table("users")

	if len(filter) > 0 {
		sql = sql.Where("id = ? AND class IN ?", studentid, filter)
	} else {
		sql = sql.Where("id = ?", studentid)
	}

	err = sql.Update("class", classid).Error
	return
}

func (c *Class) StudentLeave(studentid int64, filter ...int64) (err error) {
	sql := c.DB.Table("users")

	if len(filter) > 0 {
		sql = sql.Where("id = ? AND class IN ?", studentid, filter)
	} else {
		sql = sql.Where("id = ?", studentid)
	}

	err = sql.Update("class", 0).Error
	return
}

func (c *Class) TeacherJoin(Teacherid int64, name string) (err error) {
	classid, ok := c.GetClassIDByName(name)
	if !ok {
		return ErrClassNotExists
	}
	err = c.DB.Table("class").
		Where("teacherid = ?", Teacherid).
		Create(&model.Class{
			TeacherID: Teacherid,
			ClassID:   classid,
		}).Error
	return
}

func (c *Class) TeacherLeave(Teacherid int64, classname string) (err error) {
	if classname == "" {
		err = c.DB.Table("class").
			Where("teacherid = ?", Teacherid).
			Delete(&model.Class{}).Error
		return
	}
	classid, ok := c.GetClassIDByName(classname)
	if !ok {
		err = ErrClassNotExists
		return
	}
	err = c.DB.Table("class").
		Where("teacherid = ? AND classid = ?", Teacherid, classid).
		Delete(&model.Class{}).Error
	return
}

func (c *Class) GetTeacherClass(teacherid int64) (ret []int64, err error) {
	err = c.DB.Table("class").
		Select("classid").
		Where("teacherid = ?", teacherid).
		Find(&ret).Error

	return
}

func BuildClassIndex(tx *gorm.DB, req *model.ClassQueryRequest) *gorm.DB {
	var where []string
	var value []any
	if req.ClassID != 0 {
		where = append(where, "classid = ?")
		value = append(value, req.ClassID)
	}

	if req.TeacherID != 0 {
		where = append(where, "teacherid = ?")
		value = append(value, req.TeacherID)
	}

	if len(where) == 0 {
		return tx
	}

	return tx.Where(strings.Join(where, " OR "), value...)
}
