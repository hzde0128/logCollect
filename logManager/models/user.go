package models

type User struct {
	Id       int
	Username string `orm:"unique"` // 用户名
	Password string // 用户密码
}
