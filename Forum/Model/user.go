package Model

import "time"

type UserInfo struct {
	//id、用户名、密码、SessionId
	Id   int64 `db:"user_id"`
	Name string `db:"user_name"`
	Password  string `db:"user_password"`
	CreateTime time.Time `db:"createtime"`
}

