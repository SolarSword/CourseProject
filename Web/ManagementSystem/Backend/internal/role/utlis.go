package role

import (
	"course.project/management_system/internal/role/dao"
)

func IsAdmin(roleID string) bool {
	return dao.GetRoleType(roleID) == dao.ADMIN
}
