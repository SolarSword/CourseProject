package dao

import (
	DB "course.project/management_system/internal/common/db"
)

func GetCollegeById(collegeId string) College {
	college := &College{}
	DB.Db.GetDB().Where(&College{CollegeId: collegeId}).First(&college)
	return *college
}
