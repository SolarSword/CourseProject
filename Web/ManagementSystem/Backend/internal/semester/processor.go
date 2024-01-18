package semester

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	m "course.project/management_system/internal/common/model"
	"course.project/management_system/internal/role"
	dao "course.project/management_system/internal/semester/dao"
)

// /api/v1/create_semester
func CreateSemester(c *gin.Context) {
	var req CreateSemesterRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, CreateSemesterResponse{
			ErrorCode:    m.REQUEST_BODY_ERROR,
			ErrorMessage: m.REQUEST_BODY_ERROR_MSG,
		})
		return
	}
	if req.RoleID == "" {
		c.JSON(http.StatusOK, CreateSemesterResponse{
			ErrorCode:    m.COMPULSORY_FIELD_MISSING,
			ErrorMessage: m.COMPULSORY_FIELD_MISSING_MSG,
		})
		return
	}
	if !role.IsAdmin(req.RoleID) {
		c.JSON(http.StatusOK, CreateSemesterResponse{
			ErrorCode:    m.PERMISSON_ERROR,
			ErrorMessage: m.PERMISSON_ERROR_MSG,
		})
		return
	}
	// duration check
	if req.EndTime <= req.StartTime || !semesterDurationCheck(req.StartTime, req.EndTime, req.Type) {
		c.JSON(http.StatusOK, CreateSemesterResponse{
			ErrorCode:    INVALID_TIME_DURATION,
			ErrorMessage: INVALID_TIME_DURATION_MSG,
		})
		return
	}
	// semester format check
	if !semesterFormatCheck(req.Semester, req.Type) {
		c.JSON(http.StatusOK, CreateSemesterResponse{
			ErrorCode:    INVALID_SEMESTER,
			ErrorMessage: INVALID_SEMESTER_MSG,
		})
		return
	}
	// duplicate semester check
	if !duplicateSemesterCheck(req.Semester) {
		c.JSON(http.StatusOK, CreateSemesterResponse{
			ErrorCode:    DUPLICATED_SEMESTER,
			ErrorMessage: DUPLICATED_SEMESTER_MSG,
		})
		return
	}
	se := dao.Semester{
		Type:      req.Type,
		Semester:  req.Semester,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	}
	err = dao.CreateSemester(se)
	if err != nil {
		c.JSON(http.StatusOK, CreateSemesterResponse{
			ErrorCode:    m.DB_ERROR,
			ErrorMessage: m.DB_ERROR_MSG,
		})
	}
	c.JSON(http.StatusOK, CreateSemesterResponse{})
}

func semesterDurationCheck(startTime, endTime int64, semesterType int) bool {
	if startTime >= endTime {
		return false
	}
	start := time.Unix(startTime, 0)
	end := time.Unix(endTime, 0)
	duration := end.Sub(start)
	if semesterType == dao.SEMESTER && (duration < SEMESTER_MIN_HOURS || duration > SEMESTER_MAX_HOURS) {
		return false
	}
	if semesterType == dao.SHORT_SEMESTER && (duration < SHORT_SEMESTER_MIN_HOURS || duration > SHORT_SEMESTER_MAX_HOURS) {
		return false
	}
	return true
}

func semesterFormatCheck(semester string, semesterType int) bool {
	split := strings.Split(semester, "/")
	startYearStr := split[0]
	endYearStr := split[1][:len(split[1])-2]
	symbol := split[1][len(split[1])-2:]
	startYear, err := strconv.Atoi(startYearStr)
	if err != nil {
		return false
	}
	endYear, err := strconv.Atoi(endYearStr)
	if err != nil {
		return false
	}

	if endYear-startYear != 1 {
		return false
	}
	if semesterType == dao.SEMESTER && symbol[0] != 'A' {
		return false
	}
	if semesterType == dao.SHORT_SEMESTER && symbol[0] != 'S' {
		return false
	}
	if symbol[1] != '1' && symbol[1] != '2' {
		return false
	}
	return true
}

func duplicateSemesterCheck(semester string) bool {
	semesterFetched := dao.GetSemester(semester)
	if len(semesterFetched.Semester) != 0 {
		return false
	}
	return true
}
