package dao

import (
	"log"

	DB "course.project/management_system/internal/common/db"
	"course.project/management_system/internal/common/logger"
)

func GetCourseIdByName(courseName string) string {
	course := &Course{}
	if result := DB.Db.GetDB().Where(&Course{CourseName: courseName}).Limit(1).Find(&course); result.Error != nil {
		return ""
	}
	return course.CourseId
}

func CreateCourse(course Course) error {
	result := DB.Db.GetDB().Create(&course)
	if result.Error != nil {
		log.Printf("[%s]|datebase insertion error: %v|value: %v|\n", logger.ERROR, result.Error, course)
		return result.Error
	}
	return nil
}

func GetCollegeCourseCount(collegeId string) int64 {
	var count int64
	DB.Db.GetDB().Model(&Course{}).Where("college_id = ?", collegeId).Count(&count)
	return count
}

func BatchGetCourse(limit, offset int, collegeId, courseId, courseName string) []Course {
	var courses []Course
	DB.Db.GetDB().Where(&Course{CollegeId: collegeId, CourseId: courseId, CourseName: courseName}).Limit(limit).Offset(offset).Find(&courses)
	return courses
}
