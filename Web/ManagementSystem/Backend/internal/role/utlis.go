package role

import (
	"course.project/management_system/internal/role/dao"
)

func IsAdmin(roleID string) bool {
	return dao.GetRoleType(roleID) == dao.ADMIN
}

func IsProfessor(roleID string) bool {
	return dao.GetRoleType(roleID) == dao.PROFESSOR
}

func IsFaculty(roleID string) bool {
	roleType := dao.GetRoleType(roleID)
	return roleType == dao.ADMIN || roleType == dao.PROFESSOR
}
