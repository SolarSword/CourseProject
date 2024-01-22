package college

import (
	"strings"

	"course.project/management_system/internal/college/dao"
)

func IsValidCollege(collegeId string) bool {
	college := dao.GetCollegeById(collegeId)
	return college.CollegeName != ""
}

func GetCollegeNameAbbrByID(collegeId string) string {
	abbr := ""
	college := dao.GetCollegeById(collegeId)
	if college.CollegeName == "" {
		return abbr
	}
	split := strings.Split(college.CollegeName, " ")
	for _, s := range split {
		abbr += string(s[0])
	}
	return abbr
}
