package course

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"course.project/management_system/internal/college"
	m "course.project/management_system/internal/common/model"
	"course.project/management_system/internal/course/dao"
	"course.project/management_system/internal/role"
)

const (
	RECOMMEDEND_YEAR_MAX = 10
)

// /api/v1/create_course
func CreateCourse(c *gin.Context) {
	var req CreateCourseRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, CreateCourseResponse{
			ErrorCode:    m.REQUEST_BODY_ERROR,
			ErrorMessage: m.REQUEST_BODY_ERROR_MSG,
		})
		return
	}
	if !createCourseFieldValidation(req, c) {
		return
	}
	courseId := generateCourseID(req.CollegeId, req.RecommendedYear)
	course := dao.Course{
		CourseId:   courseId,
		CourseName: req.CourseName,
		CollegeId:  req.CollegeId,
		Credit:     req.Credit,
		Brief:      req.Brief,
	}
	err = dao.CreateCourse(course)
	if err != nil {
		c.JSON(http.StatusOK, CreateCourseResponse{
			ErrorCode:    m.DB_ERROR,
			ErrorMessage: m.DB_ERROR_MSG,
		})
		return
	}
	c.JSON(http.StatusOK, CreateCourseResponse{
		CourseId: courseId,
	})
}

func createCourseFieldValidation(req CreateCourseRequest, c *gin.Context) bool {
	if req.CreatorRoleId == "" {
		c.JSON(http.StatusOK, CreateCourseResponse{
			ErrorCode:    m.COMPULSORY_FIELD_MISSING,
			ErrorMessage: m.COMPULSORY_FIELD_MISSING_MSG,
		})
		return false
	}
	if !role.IsFaculty(req.CreatorRoleId) {
		c.JSON(http.StatusOK, CreateCourseResponse{
			ErrorCode:    m.PERMISSON_ERROR,
			ErrorMessage: m.PERMISSON_ERROR_MSG,
		})
		return false
	}

	if req.CollegeId == "" {
		c.JSON(http.StatusOK, CreateCourseResponse{
			ErrorCode:    COLLEGE_ID_MISSING,
			ErrorMessage: COLLEGE_ID_MISSING_MSG,
		})
		return false
	}
	if !college.IsValidCollege(req.CollegeId) {
		c.JSON(http.StatusOK, CreateCourseResponse{
			ErrorCode:    COLLEGE_ID_NON_EXISTING,
			ErrorMessage: COLLEGE_ID_NON_EXISTING_MSG,
		})
		return false
	}
	if req.CourseName == "" {
		c.JSON(http.StatusOK, CreateCourseResponse{
			ErrorCode:    COURSE_NAME_MISSING,
			ErrorMessage: COURSE_NAME_MISSING_MSG,
		})
		return false
	}
	if dao.GetCourseIdByName(req.CourseName) != "" {
		c.JSON(http.StatusOK, CreateCourseResponse{
			ErrorCode:    DUPLICATED_COURSE_NAME,
			ErrorMessage: DUPLICATED_COURSE_NAME_MSG,
		})
		return false
	}
	if req.RecommendedYear == 0 || req.RecommendedYear >= RECOMMEDEND_YEAR_MAX {
		c.JSON(http.StatusOK, CreateCourseResponse{
			ErrorCode:    INVALID_RECOMMENDED_YEAR,
			ErrorMessage: INVALID_RECOMMENDED_YEAR_MSG,
		})
		return false
	}
	return true
}

func generateCourseID(collegeId string, recommendedYear int) string {
	abbr := college.GetCollegeNameAbbrByID(collegeId)
	count := dao.GetCollegeCourseCount(collegeId) + 1
	countStr := strconv.FormatInt(count, 10)
	for len(countStr) < 4 {
		countStr = "0" + countStr
	}
	return fmt.Sprintf("%s%d%s", abbr, recommendedYear, countStr)
}

// /api/v1/get_course
func GetCourse(c *gin.Context) {
	var req GetCourseRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, GetCourseResponse{
			ErrorCode:    m.REQUEST_BODY_ERROR,
			ErrorMessage: m.REQUEST_BODY_ERROR_MSG,
		})
		return
	}
	fmt.Printf("req: %v\n", req)
	if req.RoleId == "" {
		c.JSON(http.StatusOK, GetCourseResponse{
			ErrorCode:    m.COMPULSORY_FIELD_MISSING,
			ErrorMessage: m.COMPULSORY_FIELD_MISSING_MSG,
		})
		return
	}
	if !role.IsFaculty(req.RoleId) {
		c.JSON(http.StatusOK, GetCourseResponse{
			ErrorCode:    m.PERMISSON_ERROR,
			ErrorMessage: m.PERMISSON_ERROR_MSG,
		})
		return
	}
	ok, limit, offset, collegeId, courseId, courseName := getCourseQueryParameterHandling(c)
	if !ok {
		return
	}
	courses := dao.BatchGetCourse(limit, offset, collegeId, courseId, courseName)
	courseList := []CourseElement{}
	for _, c := range courses {
		courseList = append(courseList, CourseElement{
			CourseId:   c.CourseId,
			CourseName: c.CourseName,
			CollegeId:  c.CollegeId,
			Credit:     c.Credit,
			Brief:      c.Brief,
		})
	}
	c.JSON(http.StatusOK, GetCourseResponse{
		CourseList: courseList,
	})
}

func getCourseQueryParameterHandling(c *gin.Context) (bool, int, int, string, string, string) {
	limitStr := c.DefaultQuery("limit", m.DEFAULT_LIMIT)
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusOK, CreateCourseResponse{
			ErrorCode:    m.INVALID_LIMIT,
			ErrorMessage: m.INVALID_LIMIT_MSG,
		})
		return false, 0, 0, "", "", ""
	}
	offsetStr := c.DefaultQuery("offset", m.DEFAULT_OFFSET)
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(http.StatusOK, CreateCourseResponse{
			ErrorCode:    m.INVALID_OFFSET,
			ErrorMessage: m.INVALID_OFFSET_MSG,
		})
		return false, 0, 0, "", "", ""
	}
	collegeId := c.Query("college_id")
	courseId := c.Query("course_id")
	courseName := c.Query("course_name")
	return true, limit, offset, collegeId, courseId, courseName
}
