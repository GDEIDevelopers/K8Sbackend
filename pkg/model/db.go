package model

type UserSecurity struct {
	EmailVerify bool `json:"emailVerify"`
}

type SecurityQuestion struct {
	Title  string `json:"title"`
	Answer string `json:"answer"`
}

type User struct {
	ID               int64              `gorm:"column:id"`
	Name             string             `gorm:"column:name"`
	Email            string             `gorm:"column:email"`
	RealName         string             `gorm:"column:realName"`
	UserSchoollD     string             `gorm:"column:userSchoollD"`
	SchoolCode       string             `gorm:"column:schoolCode"`
	Class            string             `gorm:"column:class"`
	Role             string             `gorm:"column:role"`
	Sex              string             `gorm:"column:sex"`
	Password         string             `gorm:"column:password"`
	Security         UserSecurity       `gorm:"serializer:json;column:security"`
	SecurityQuestion []SecurityQuestion `gorm:"serializer:json;column:securityQuestion"`
}
