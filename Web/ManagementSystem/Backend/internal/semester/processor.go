package semester

import (
	"net/http"
	"strconv"
	"strings"
	"sync"
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
		return
	}
	UpdateCurrentSemester(req.Semester, req.StartTime, req.EndTime, req.Type)
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
	return len(semesterFetched.Semester) == 0
}

type CurrentSemester struct {
	semester     string
	startTime    int64
	endTime      int64
	semesterType int
}

func (s CurrentSemester) GetSemester() string {
	return s.semester
}

func (s CurrentSemester) GetStartTime() int64 {
	return s.startTime
}

func (s CurrentSemester) GetEndTime() int64 {
	return s.endTime
}

func (s CurrentSemester) GetSemesterType() int {
	return s.semesterType
}

var semesterSingleton *CurrentSemester
var lock = &sync.Mutex{}

func GetSemesterSingleton() *CurrentSemester {
	lock.Lock()
	defer lock.Unlock()
	// by default, the current semester would be vacation
	if semesterSingleton == nil {
		semesterSingleton = &CurrentSemester{
			semesterType: dao.VACATION,
		}
	} else if semesterSingleton.endTime > 0 &&
		semesterSingleton.endTime <= time.Now().Unix() {
		semesterSingleton.semester = ""
		semesterSingleton.startTime = 0
		semesterSingleton.endTime = 0
		semesterSingleton.semesterType = 0
	}
	return semesterSingleton
}

func UpdateCurrentSemester(semester string, startTime, endTime int64, semesterType int) {
	lock.Lock()
	defer lock.Unlock()
	if semesterSingleton == nil {
		semesterSingleton = &CurrentSemester{
			semester:     semester,
			startTime:    startTime,
			endTime:      endTime,
			semesterType: semesterType,
		}
	} else {
		semesterSingleton.semester = semester
		semesterSingleton.startTime = startTime
		semesterSingleton.endTime = endTime
		semesterSingleton.semesterType = semesterType
	}
}

// /api/v1/get_current_semester
func GetCurrentSemester(c *gin.Context) {
	se := GetSemesterSingleton()
	c.JSON(http.StatusOK, GetCurrentSemesterResponse{
		Semester:  se.semester,
		StartTime: se.startTime,
		EndTime:   se.endTime,
		Type:      se.semesterType,
	})
}
