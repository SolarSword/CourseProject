package course

const (
	COURSE_NAME_MISSING      = 3001
	DUPLICATED_COURSE_NAME   = 3002
	COLLEGE_ID_MISSING       = 3003
	COLLEGE_ID_NON_EXISTING  = 3004
	INVALID_RECOMMENDED_YEAR = 3005
)

const (
	COURSE_NAME_MISSING_MSG      = "the course name is missing"
	DUPLICATED_COURSE_NAME_MSG   = "the course name is duplicated with an existing course"
	COLLEGE_ID_MISSING_MSG       = "the college id is missing"
	COLLEGE_ID_NON_EXISTING_MSG  = "the college id does not exist"
	INVALID_RECOMMENDED_YEAR_MSG = "the recommended year is invalid"
)

type CreateCourseRequest struct {
	CourseName      string `json:"course_name"`
	CollegeId       string `json:"college_id"`
	Credit          int    `json:"credit"`
	Brief           string `json:"brief"`
	CreatorRoleId   string `json:"creator_role_id"`
	RecommendedYear int    `json:"recommended_year"`
}

type CreateCourseResponse struct {
	CourseId     string `json:"course_id"`
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

type GetCourseRequest struct {
	RoleId string `json:"role_id"`
}

type GetCourseResponse struct {
	CourseList   []CourseElement `json:"course_list"`
	ErrorCode    int             `json:"error_code"`
	ErrorMessage string          `json:"error_message"`
}

type CourseElement struct {
	CourseId   string `json:"course_id"`
	CourseName string `json:"course_name"`
	CollegeId  string `json:"college_id"`
	Credit     int    `json:"credit"`
	Brief      string `json:"brief"`
}
