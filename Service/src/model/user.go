package model

type User struct {
	Id 			string
	// 用户名
	Username 	string
	// 两次md5后的密码
	Password 	string
	// 权限
	Privilege 	int
	// 状态
	State		int
	// 空间使用情况
	EmptySpace	int
	UseSpace	int
	// 分享码
	ShareCode	string
}