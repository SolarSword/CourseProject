package dao

import (
	DB "course.project/management_system/internal/common/db"
)

func GetRoleType(roleId string) int {
	role := &Role{}
	DB.Db.GetDB().Where(&Role{RoleId: roleId}).Limit(1).Find(&role)
	return role.Type
}
