package dao

import (
	DB "course.project/authentication/internal/common/db"
)

func IsValidUser(userName string, passWord string) bool {
	user := &User{}
	DB.Db.GetDB().Where(&User{UserName: userName}).First(&user)
	return passWord == user.PassWord
}
