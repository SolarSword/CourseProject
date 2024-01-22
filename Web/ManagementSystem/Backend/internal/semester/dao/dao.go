package dao

import (
	"log"

	DB "course.project/management_system/internal/common/db"
	"course.project/management_system/internal/common/logger"
)

func GetSemester(semester string) *Semester {
	se := &Semester{}
	DB.Db.GetDB().Where(&Semester{Semester: semester}).Limit(1).Find(&se)
	return se
}

func CreateSemester(semester Semester) error {
	result := DB.Db.GetDB().Create(&semester)
	if result.Error != nil {
		log.Printf("[%s]|datebase insertion error: %v|value: %v|\n", logger.ERROR, result.Error, semester)
		return result.Error
	}
	return nil
}
