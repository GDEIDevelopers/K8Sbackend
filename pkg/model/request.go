package model

type QueryRequest struct {
	UserID     int    `json:"id"`
	Name       string `json:"name"`
	QueryEmail string `json:"queryemail"`
}

type RegisterUserRequest struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	RealName     string `json:"realName"`
	UserSchoollD string `json:"userSchoollD"`
	SchoolCode   string `json:"schoolCode"`
	Class        string `json:"class"`
	Sex          string `json:"sex"`
	Password     string `json:"password"`
}

type UserLoginRequest struct {
	QueryRequest
	Password string `json:"password"`
}

type ModifyUserRequest struct {
	Email        string `json:"email,omitempty"`
	RealName     string `json:"realName,omitempty"`
	UserSchoollD string `json:"userSchoollD,omitempty"`
	SchoolCode   string `json:"schoolCode,omitempty"`
	Class        string `json:"class,omitempty"`
	Sex          string `json:"sex,omitempty"`
}
type ModifyUserPasswordRequest struct {
	Password string `json:"password"`
}

type UserIDOnlyRequest struct {
	UserID int `json:"id"`
}

type AdminModifyPasswordRequest struct {
	UserIDOnlyRequest
	Password string `json:"password"`
}

type AdminModifyRequest struct {
	UserIDOnlyRequest
	Email        string `json:"email,omitempty"`
	RealName     string `json:"realName,omitempty"`
	UserSchoollD string `json:"userSchoollD,omitempty"`
	SchoolCode   string `json:"schoolCode,omitempty"`
	Class        string `json:"class,omitempty"`
	Sex          string `json:"sex,omitempty"`
}

type CommonClassRequest struct {
	ClassName string `json:"classname"`
}

type GetTeacherStudentRequest struct {
	TeacherID int64 `json:"teacherid"`
}

type ClassQueryRequest struct {
	TeacherID int64 `json:"teacherid"`
	ClassID   int64 `json:"classid"`
}

type TeacherClassRequest struct {
	TeacherID int64  `json:"teacherid"`
	ClassName string `json:"classname"`
}

type StudentClassRequest struct {
	StudentID int64  `json:"studentid"`
	ClassName string `json:"classname"`
}

type StudentLeaveClassRequest struct {
	StudentID int64 `json:"studentid"`
}

type TeacherAddStudentRequest struct {
	StudentIDs []int64 `json:"studentid"`
	ClassName  string  `json:"classname"`
}

type TeacherRemoveStudentRequest struct {
	StudentIDs []int64 `json:"studentid"`
}
